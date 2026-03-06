package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rw   sync.RWMutex
	data = 0
)

func main() {
	var wg sync.WaitGroup

	// 启动 5 个读协程，它们会长期持有读锁
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 尝试获取读锁
			rw.RLock()
			fmt.Printf("[Reader %d] 获取读锁成功 (当前时间: %v)\n", id, time.Now().Format("15:04:05.000"))

			// 模拟长耗时读操作 (持锁 2 秒)
			time.Sleep(2 * time.Second)

			fmt.Printf("[Reader %d] 释放读锁 (当前时间: %v)\n", id, time.Now().Format("15:04:05.000"))
			rw.RUnlock()
		}(i)
	}

	// 稍微延迟，确保读协程先启动并拿到锁
	time.Sleep(100 * time.Millisecond)

	// 启动 1 个写协程
	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Printf("[Writer ] 尝试获取写锁... (当前时间: %v)\n", time.Now().Format("15:04:05.000"))

		// 【关键点】这里会阻塞，直到上面 5 个 Reader 全部释放
		rw.Lock()

		fmt.Printf("[Writer ] >>> 获取写锁成功！(当前时间: %v) <<<\n", time.Now().Format("15:04:05.000"))
		data = 100 // 修改数据
		time.Sleep(500 * time.Millisecond)

		fmt.Printf("[Writer ] 释放写锁 (当前时间: %v)\n", time.Now().Format("15:04:05.000"))
		rw.Unlock()
	}()

	// 启动几个“迟到”的读协程，测试它们是否会被写锁阻塞
	time.Sleep(200 * time.Millisecond) // 在 Writer 等待期间启动
	for i := 6; i <= 8; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("[Reader %d] (迟到者) 尝试获取读锁... (当前时间: %v)\n", id, time.Now().Format("15:04:05.000"))

			rw.RLock() // 这里会被阻塞，因为 Writer 已经在排队了
			fmt.Printf("[Reader %d] (迟到者) 获取读锁成功 (当前时间: %v)\n", id, time.Now().Format("15:04:05.000"))
			rw.RUnlock()
		}(i)
	}

	wg.Wait()
}
