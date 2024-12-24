package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("前】")
	readMem()

	const slen = 1_000_000

	s1 := make([]int, slen)
	fmt.Printf("s1: len=%d cap=%d\n", len(s1), cap(s1))
	time.Sleep(3 * time.Second) // 必须，否则，下面这个 readMem() 的结果和上面的相同

	fmt.Println("\n中】")
	readMem()

	// 填充数据
	// 填充后，内存使用才有变化
	for i := range [slen]struct{}{} {
		s1 = append(s1, i)
	}
	time.Sleep(3 * time.Second) // 非必须

	fmt.Println("\n后】")
	readMem()
}

func readMem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// fmt.Println(ms) // 太复杂了

	// 内存 通义千问
	// 打印一些关键的内存统计信息
	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)           // 分配但未释放的内存
	fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024) // 程序启动以来分配的总内存
	fmt.Printf("Sys = %v MiB\n", m.Sys/1024/1024)               // 从操作系统分配的总内存
	fmt.Printf("HeapAlloc = %v MiB\n", m.HeapAlloc/1024/1024)   // 从堆上分配但未释放的内存
	fmt.Printf("HeapSys = %v MiB\n", m.HeapSys/1024/1024)       // 由Go分配的堆内存的系统内存大小
	fmt.Printf("NumGC = %v\n", m.NumGC)                         // 进行的GC次数
	fmt.Printf("HeapObjs = %v\n", m.HeapObjects/1024/1024)      // 进行的GC次数
}

func m() {
	fmt.Printf("bool: %d\n", unsafe.Alignof(bool(true)))
	fmt.Printf("string: %d\n", unsafe.Alignof(string("a")))
	fmt.Printf("int: %d\n", unsafe.Alignof(int(0)))
	fmt.Printf("int32: %d\n", unsafe.Alignof(int32(0)))
	fmt.Printf("int64: %d\n", unsafe.Alignof(int64(0)))
	fmt.Printf("float64: %d\n", unsafe.Alignof(float64(0)))
	fmt.Printf("float32:%d\n", unsafe.Alignof(float32(0)))
}
