package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("vim-go")
	var wg sync.WaitGroup
	concurrencyKey := make(chan int, 4)
	for n := 0; n < 5000; n++ {
		wg.Add(1)
		concurrencyKey <- n
		go func(i int) {
			defer func() {
				wg.Done()
				<-concurrencyKey
			}()
			rand.Seed(time.Now().UnixNano())

			rn := rand.Intn(200)
			fmt.Println("Operation number is", n, ",random", rn, "channel buffer length", len(concurrencyKey))
		}(n)
	}
	wg.Wait()
}
