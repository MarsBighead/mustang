# xm 操作指南

## 概述

`xm` 是一个用于检测和清理重复文件的命令行工具，主要针对微信存储目录优化。

## 安装

```bash
cd mustang
make build
# 或
go build -o bin/xm ./cmd/xm
```

## 快速开始

### 1. 分析微信目录

```bash
# 使用默认微信目录
./bin/xm analyze

# 或指定自定义目录
./bin/xm analyze /path/to/dir1 /path/to/dir2
```

**说明**：
- 扫描目录并计算文件元数据
- 只对大小重复的文件计算 MD5（性能优化）
- 数据存储在 `~/.xm/xm.db`

### 2. 查看统计信息

```bash
# 查看重复统计（默认保留最旧的文件）
./bin/xm stats

# 保留最新的文件
./bin/xm stats --keep newest

# 显示前 10 个重复组
./bin/xm stats --top 10
```

**输出包括**：
- 运行信息（命令、开始/结束时间、根目录）
- 文件统计（总数、总大小、已哈希文件数）
- 重复组统计
- 浪费的空间
- Top N 重复组详情

**重要**：stats 命令会将重复文件信息写入 `duplicates` 表，为 dedup 做准备。

### 3. 预览去重操作

```bash
# 预览哪些文件会被移动（不实际执行）
./bin/xm dedup --dry-run
```

**输出示例**：
```
[5b7b5b16] 15 B x 2 copies (keep: dir1/file1.txt)
  + dir1/file1.txt
  - dir2/file1_dup.txt

Would move 2 duplicate files, reclaimed 30 B
```

**符号说明**：
- `+` 保留的文件
- `-` 将被隔离的文件
- `~` 硬链接（跳过）
- `->` 隔离目标路径

### 4. 执行去重

```bash
# 实际执行隔离操作
./bin/xm dedup

# 自定义隔离目录
./bin/xm dedup --quarantine ~/my-quarantine
```

**隔离后的文件**：
- 原位置的文件被删除
- 文件移动到 `~/.xm/quarantine/` 目录
- 保持原始目录结构
- 可随时手动恢复

### 5. 恢复隔离的文件

```bash
# 查看隔离目录
ls ~/.xm/quarantine/

# 手动恢复文件
mv ~/.xm/quarantine/path/to/file /original/location/
```

## 完整工作流程

```bash
# 步骤 1: 清理旧数据（可选）
rm ~/.xm/xm.db*

# 步骤 2: 分析
./bin/xm analyze

# 步骤 3: 查看统计
./bin/xm stats --keep oldest --top 20

# 步骤 4: 预览去重
./bin/xm dedup --dry-run

# 步骤 5: 执行去重
./bin/xm dedup

# 步骤 6: 验证结果
./bin/xm stats
```

## 高级用法

### 指定数据库位置

```bash
./bin/xm --db /custom/path/xm.db analyze
./bin/xm --db /custom/path/xm.db stats
```

### 增量扫描

再次运行 `analyze` 时：
- 未改变的文件（大小和修改时间相同）会被跳过
- 只处理新增和修改的文件
- 大幅减少扫描时间

```bash
# 第一次完整扫描
./bin/xm analyze

# 几天后再次扫描（只处理变化）
./bin/xm analyze
```

### 查看数据库

```bash
# 使用 sqlite3 直接查询
sqlite3 ~/.xm/xm.db

# 查看文件表结构
.schema files

# 查看重复组
SELECT md5, COUNT(*) as cnt FROM files 
WHERE run_id=1 AND md5 != '' AND md5 != 'unique'
GROUP BY md5 HAVING cnt > 1
ORDER BY cnt DESC;

# 查看隔离状态
SELECT * FROM duplicates WHERE run_id=1;
```

## 数据库文件说明

| 文件 | 说明 |
|------|------|
| `xm.db` | 主数据库文件 |
| `xm.db-wal` | WAL 日志（Write-Ahead Log） |
| `xm.db-shm` | 共享内存文件 |

**清理 WAL 文件**：
```bash
sqlite3 ~/.xm/xm.db "PRAGMA wal_checkpoint(TRUNCATE);"
```

## 常见问题

### Q: 为什么有些文件没有被标记为重复？

A: 可能的原因：
1. **硬链接**: 相同 inode 的文件被视为同一文件
2. **大小不同**: 即使内容相似，大小不同也不会被识别为重复
3. **MD5 未计算**: 大小唯一的文件不会计算 MD5

### Q: 隔离后文件还能恢复吗？

A: 可以。隔离只是移动文件到 quarantine 目录，可以随时手动恢复。

### Q: 如何排除某些目录？

A: 目前不支持排除。可以手动指定要扫描的目录：
```bash
./bin/xm analyze /path/to/include1 /path/to/include2
```

### Q: 扫描太慢怎么办？

A: 
1. 使用增量扫描（再次运行 analyze）
2. 缩小扫描范围（指定具体目录）
3. 检查磁盘 I/O 性能

### Q: 如何查看浪费了多少空间？

A: 运行 stats 命令：
```bash
./bin/xm stats
```
查看 "Wasted space" 字段。

## 安全建议

1. **先用 --dry-run 预览**: 总是先预览再去重
2. **检查 quarantine 目录**: 确认文件被正确隔离
3. **保留备份**: 重要数据建议先备份
4. **逐步清理**: 可以分批次处理不同目录

## 性能数据

基于 390,000+ 文件的测试：
- **扫描时间**: 约 2-5 分钟
- **MD5 计算**: 约 5-15 分钟（取决于重复文件数量）
- **数据库大小**: 约 38 MB
- **内存占用**: 约 100-200 MB
