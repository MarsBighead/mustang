# xm 开发设计文档

## 项目概述

`xm` 是一个用 Go 编写的命令行工具，用于检测和清理重复文件，特别针对微信存储目录优化。

## 技术栈

| 组件 | 技术 | 版本 |
|------|------|------|
| 编程语言 | Go | 1.25.0 |
| 命令行框架 | cobra | 1.10.2 |
| 数据库 | SQLite (modernc.org/sqlite) | 1.56.0 |
| 数据库模式 | WAL | - |

## 架构设计

### 目录结构

```
mustang/
├── cmd/
│   └── xm/
│       └── main.go          # 命令行入口
├── internal/
│   ├── db.go                # 数据库操作
│   ├── model.go             # 数据模型
│   └── scanner.go           # 文件扫描和 MD5 计算
├── docs/                    # 文档
├── bin/                     # 编译输出（gitignore）
├── Makefile                 # 构建脚本
└── go.mod                   # Go 模块定义
```

### 核心模块

#### 1. 数据模型 (`internal/model.go`)

```go
// 文件状态常量
const (
    StatusAnalyze = 0  // 已扫描，MD5 待计算
    StatusDedup   = 1  // 已去重处理
)

// 文件信息
type FileInfo struct {
    Path    string    // 相对路径
    Size    int64
    MD5     string
    ModTime time.Time
    Inode   uint64    // 文件系统 inode
    RunID   int64
    Status  int
}

// 重复组
type DuplicateGroup struct {
    MD5   string
    Files []FileInfo
}
```

#### 2. 数据库层 (`internal/db.go`)

**表结构**：

```sql
-- 文件表
CREATE TABLE files (
    path    TEXT NOT NULL,
    size    INTEGER NOT NULL,
    md5     TEXT NOT NULL DEFAULT '',
    modtime INTEGER NOT NULL,
    inode   INTEGER NOT NULL DEFAULT 0,
    run_id  INTEGER NOT NULL DEFAULT 0,
    status  INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (path, run_id)
);

-- 重复文件表
CREATE TABLE duplicates (
    run_id  INTEGER NOT NULL,
    md5     TEXT NOT NULL,
    path    TEXT NOT NULL,
    size    INTEGER NOT NULL,
    modtime INTEGER NOT NULL,
    inode   INTEGER NOT NULL,
    keep    INTEGER NOT NULL DEFAULT 0,  -- 1=保留，0=删除
    PRIMARY KEY (run_id, md5, path)
);

-- 元数据表
CREATE TABLE meta (
    key   TEXT PRIMARY KEY,
    value TEXT NOT NULL
);
```

**关键方法**：

| 方法 | 说明 |
|------|------|
| `OpenDB(path)` | 打开数据库（WAL 模式） |
| `NewRun(command)` | 创建新运行记录 |
| `Scan.Scan(ctx)` | 扫描目录，检测变化 |
| `Scan.ComputeMD5s(ctx, paths)` | 并发计算 MD5 |
| `DB.ComputeGrouped(runID, paths)` | 按大小分组（分批处理） |
| `DB.FindDuplicates(runID)` | 查找重复文件 |
| `DB.InsertDuplicates(runID, groups, strategy)` | 写入 duplicates 表 |
| `DB.GetDuplicates(runID)` | 读取 duplicates 表 |

#### 3. 扫描器 (`internal/scanner.go`)

**扫描流程**：

1. **增量检测**：
   - 比较文件大小和修改时间（Unix 秒精度）
   - 跳过未改变的文件
   - 记录新增和修改的文件

2. **相对路径**：
   - 数据库存储相对路径
   - 支持多个根目录
   - `toRel()` 和 `toAbs()` 转换路径

3. **并发 MD5**：
   - Worker 池模式（默认 4 个 worker）
   - Mutex 保护数据库写入
   - 支持上下文取消

4. **Inode 检测**：
   - 使用 `syscall.Stat_t.Ino` 获取 inode
   - 相同 inode 的文件视为硬链接
   - 硬链接不算重复

#### 4. 命令行 (`cmd/xm/main.go`)

**命令结构**：

```
xm
├── analyze [dirs...]    # 扫描并建立索引
├── stats                # 统计并识别重复
└── dedup                # 隔离重复文件
```

**全局选项**：
- `--db`: 数据库路径（默认 `~/.xm/xm.db`）

**analyze 选项**：
- 无额外选项

**stats 选项**：
- `--keep`: 保留策略（oldest/newest，默认 oldest）
- `--top`: 显示前 N 个重复组（默认 5）

**dedup 选项**：
- `--quarantine`: 隔离目录（默认 `~/.xm/quarantine`）
- `--dry-run`: 预览模式

## 核心算法

### 1. 重复检测流程

```
analyze:
  1. 扫描目录，记录文件元数据
  2. 找出大小重复的文件
  3. 只对大小重复的文件计算 MD5
  4. 大小唯一的文件标记为 'unique'

stats:
  1. 查询 MD5 重复的文件
  2. 按 keep 策略标记保留/删除
  3. 写入 duplicates 表
  4. 显示统计信息

dedup:
  1. 从 duplicates 表读取
  2. 移动标记为删除的文件到 quarantine
  3. 更新数据库
```

### 2. 性能优化

#### 大小预过滤
- 只对大小重复的文件计算 MD5
- 大幅减少 MD5 计算量
- 基于假设：大小不同的文件内容不同

#### 增量扫描
- 缓存文件大小和修改时间
- 只处理变化的文件
- 支持中断后继续

#### 分批处理
- SQL 查询分批（每批 500 个）
- 避免 SQLite 变量限制（999）
- 减少内存占用

#### 并发 MD5
- 4 个 worker 并发计算
- Mutex 保护数据库写入
- WAL 模式支持并发读写

### 3. 硬链接检测

**问题**：微信大量使用硬链接，相同内容可能被多次链接

**解决**：
```go
// 获取 inode
func getInode(info os.FileInfo) uint64 {
    if st, ok := info.Sys().(*syscall.Stat_t); ok {
        return st.Ino
    }
    return 0
}

// SQL 查询排除硬链接
SELECT md5 FROM files
WHERE run_id=? AND md5 != '' AND md5 != 'unique'
GROUP BY md5 HAVING COUNT(DISTINCT inode) > 1
```

## 数据流

```
用户输入目录
    ↓
filepath.WalkDir 扫描
    ↓
检测文件变化（size + mtime）
    ↓
插入/更新 files 表
    ↓
ComputeGrouped 按大小分组
    ↓
并发计算 MD5（仅重复大小）
    ↓
更新 files 表的 md5 字段
    ↓
stats: FindDuplicates 查找重复
    ↓
InsertDuplicates 写入 duplicates 表
    ↓
dedup: GetDuplicates 读取
    ↓
移动文件到 quarantine
    ↓
更新数据库
```

## 配置和常量

### 默认路径

```go
defaultDBPath = "~/.xm/xm.db"
defaultQuarantine = "~/.xm/quarantine"

// 微信默认目录
defaultWeChatDirs() = []string{
    "~/Library/Containers/com.tencent.xinWeChat/Data/Library/Application Support/com.tencent.xinWeChat",
    "~/Library/Containers/com.tencent.xinWeChat/Data/Documents",
}
```

### 数据库配置

```go
// WAL 模式 + 5秒超时
"sqlite", path+"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)"
```

### 并发配置

```go
Scanner.Workers = 4  // MD5 计算 worker 数
batchSize = 500      // SQL 批处理大小
```

## 错误处理

### 数据库错误
- 使用 `busy_timeout(5000)` 处理并发写入
- Mutex 保护关键操作
- 事务支持批量操作

### 文件操作错误
- 跳过无法访问的文件
- 隔离失败时回滚计数
- 支持跨文件系统移动（copy + delete）

### 上下文取消
- 支持 Ctrl+C 中断
- 检查 `ctx.Err()` 提前退出
- 已处理的数据会保存

## 扩展性设计

### 当前限制
- 不支持排除目录
- 不支持文件类型过滤
- 单线程目录遍历

### 未来可扩展
- 添加排除模式（glob 匹配）
- 添加文件类型过滤
- 并发目录遍历
- 支持多种哈希算法（SHA256 等）
- 支持自定义保留策略
- Web UI 查看和管理

## 测试建议

### 单元测试
- 测试 `toRel()` 和 `toAbs()` 路径转换
- 测试 `ComputeGrouped` 分批逻辑
- 测试 `pickFilesForDup` 保留策略

### 集成测试
- 创建测试目录
- 运行完整流程
- 验证数据库状态
- 验证文件隔离

### 性能测试
- 大规模文件（100 万+）
- 测量扫描时间
- 测量 MD5 计算时间
- 测量内存占用

## 安全考虑

1. **文件隔离而非删除**: 可恢复
2. **硬链接检测**: 避免误删
3. **预览模式**: --dry-run 先查看
4. **权限检查**: 使用标准文件操作
5. **路径验证**: 避免路径遍历攻击

## 版本历史

### v1.0.0 (2026-08-18)
- 初始版本
- 支持 analyze、stats、dedup 命令
- 增量扫描
- 并发 MD5 计算
- 硬链接检测
- duplicates 表设计
- WAL 模式支持
