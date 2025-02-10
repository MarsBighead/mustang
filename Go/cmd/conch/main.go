package main

import (
	"fmt"
	"time"
)

func main() {
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
	for i := 0; i < 10; i++ {
		go func(n int) {
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
	time.Sleep(3 * time.Second)
}
