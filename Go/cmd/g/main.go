package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	fmt.Println("Go channel test")
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	//r := fn(100000)
	c1, c2 := make(chan int, 10), make(chan int, 10)
	go func() {

		for i := 0; i < 1000; i++ {
			c1 <- i
			time.Sleep(time.Microsecond)
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			c2 <- i
			time.Sleep(time.Microsecond)
		}
	}()
	for {
		select {
		case data := <-c1:
			fmt.Println("ch 1:", data)
		case data := <-c2:
			fmt.Println("ch 2", data)
		}
	}
}
