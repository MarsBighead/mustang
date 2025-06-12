package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()
	go task(ctx)
	var ch chan int
	// read from uninitialized chan
	// x := <-ch /// fatal error: all goroutines are asleep - deadlock!
	// fmt.Println("x from chan", x)
	// write to uninitialized chan
	//ch <- 1             /// fatal error:: all goroutines are asleep - deadlock!
	ch = make(chan int) // init channel
	close(ch)
	// close(ch) // panic: close of closed channel
	// ch <- 1 // panic: send on closed channel
	v := <-ch // 缓冲取完，就是0（int零值）
	fmt.Println("value after channel closed", v)
	time.Sleep(time.Second * 6)
}

func task(ctx context.Context) {
	ch := make(chan struct{}, 0)
	go func() {
		// mock a task which costs 4s
		time.Sleep(time.Second * 5)
		fmt.Println("task accomplished")
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}

// practice
func practice() {
	// 不佳：每次循环都阻塞
	ch := make(chan int)
	for {
		data := <-ch // 频繁 gopark/goready
		process(data)
	}
	buffer := []int{}
	batchSize := 10
	// 优化：批量处理
	for {
		select {
		case data := <-ch:
			buffer = append(buffer, data)
			if len(buffer) >= batchSize {
				processBatch(buffer)
				buffer = buffer[:0]
			}
		case <-time.After(time.Millisecond):
			if len(buffer) > 0 {
				processBatch(buffer)
				buffer = buffer[:0]
			}
		}
	}
}

func process(data int) {
	fmt.Println("data", data)
}

func processBatch(buffer []int) {
	fmt.Println("data", buffer)
}
