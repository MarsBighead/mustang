package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	content := "ABC"
	sig := make(chan bool)
	ch := make(chan rune, 3)
	go func() {
		for i := 0; i < 10; i++ {
			for _, c := range content {
				ch <- c
			}
			sig <- true
		}
		close(ch)
	}()

	i, j := 1, 1
	now := time.Now()
	/*
		output: ABC
		for v := range ch {
			if i%3 == 0 {
				<-sig
				fmt.Printf("%s No.%d\n", string(v), j)
				j++
			} else {
				fmt.Printf("%s", string(v))
			}
			i++
		}
	*/
	// output: AB
	for v := range ch {
		if i%3 > 0 {
			fmt.Printf("%s", string(v))
		} else {
			fmt.Printf(" No.%d\n", j)
			<-sig
			j++
		}
		i++
	}
	//close(sig)
	d := time.Since(now)
	fmt.Printf("Running time: %v\n", d)
	first()
}

func first() {
	content := "ABC"
	sig := make(chan bool, 10)
	ch := make(chan rune)
	// 启动生成数据的协程
	go func() {
		for i := 0; i < 10; i++ {
			for _, c := range content {
				ch <- c
			}
		}
		close(ch)

	}()

	/*cnt := 0
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("%s", string(v))
		cnt++
		if cnt > 0 && cnt%3 == 0 {
			sig <- true
			fmt.Println()
		}
		//ch <- ele
	}
	*/
	sig <- true
	now := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			v := <-sig
			if v {
				fmt.Printf("Serial %d: ", n)
				for j := 0; j < 3; j++ {
					v, ok := <-ch
					if ok {
						fmt.Printf("%s", string(v))
					} else {
						return
					}
				}
				fmt.Println()
				sig <- true
			}
		}(i)
	}
	wg.Wait()
	d := time.Since(now)
	fmt.Printf("Running time: %v\n", d)
	time.Sleep(3 * time.Second)
}
