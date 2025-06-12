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
