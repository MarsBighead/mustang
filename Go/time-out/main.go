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
