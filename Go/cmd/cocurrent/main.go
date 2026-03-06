package main

import (
	"fmt"
	"sync"
	"time"
)

/*
- 交替打印数字和字母：如 12AB34CD...。创建两个无缓冲 Channel 作为“令牌”，分别控制数字和字母协程的执行顺序。数字协程打印后向字母令牌发送信号，字母协程打印后向数字令牌发送信号。

- 睡眠排序（Sleep Sort）：为数组每个元素启动一个协程，协程内 time.Sleep(elem * time.Millisecond)，然后将元素发送到同一个 Channel。主协程从 Channel 接收，接收顺序即为元素升序。

- 打印 0 与奇偶数（LeetCode 1116 变体）：创建三个 Channel，分别控制打印零、奇数、偶数的协程。核心逻辑是：zero 协程打印后，根据下一个要打印的数是奇数还是偶数，向对应的 odd 或 even 令牌 Channel 发送信号。
*/
// 1. 交替打印数字和字母：12AB34CD... （每组两个数字、两个字母）
func alternatePrint() {
	fmt.Println("=== 交替打印数字和字母 ===")

	numToken := make(chan struct{})    // 数字令牌
	letterToken := make(chan struct{}) // 字母令牌

	// 数字协程
	go func() {
		for i := 1; i <= 26; i += 2 {
			<-numToken // 等待数字令牌
			// 打印两个数字
			fmt.Printf("%d%d", i, i+1)
			letterToken <- struct{}{} // 交给字母
		}
	}()

	// 字母协程
	go func() {
		for ch := 'A'; ch <= 'Z'; ch += 2 {
			<-letterToken // 等待字母令牌
			// 打印两个字母
			fmt.Printf("%c%c", ch, ch+1)
			numToken <- struct{}{} // 交回数字
		}
	}()

	numToken <- struct{}{} // 启动数字协程

	// 等待足够时间让打印完成（实际可用 sync.WaitGroup，但这里简单 sleep）
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

// 2. 睡眠排序（Sleep Sort）
func sleepSort(arr []int) {
	fmt.Println("=== 睡眠排序 ===")

	ch := make(chan int)
	var wg sync.WaitGroup

	// 为每个元素启动一个协程
	for _, v := range arr {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			time.Sleep(time.Duration(val) * time.Millisecond)
			ch <- val
		}(v)
	}

	// 等待所有发送完成，关闭 channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 主协程接收并打印，顺序即为升序
	for v := range ch {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// 3. 打印 0 与奇偶数（LeetCode 1116 变体）
func zeroOddEven(n int) {
	fmt.Printf("=== 打印 0 与奇偶数（n=%d） ===\n", n)

	zeroCh := make(chan struct{})
	oddCh := make(chan struct{})
	evenCh := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(3) // 三个协程

	// zero 协程
	go func() {
		defer wg.Done()
		for i := 1; i <= n; i++ {
			<-zeroCh // 等待启动信号（初始由 main 发送）
			fmt.Print(0)
			if i%2 == 1 {
				oddCh <- struct{}{} // 下一个打印奇数
			} else {
				evenCh <- struct{}{} // 下一个打印偶数
			}
			// 等待奇数或偶数协程完成
			<-zeroCh
		}
		// 所有数字打印完毕，关闭 odd/even 通道让它们退出
		close(oddCh)
		close(evenCh)
	}()

	// odd 协程
	go func() {
		defer wg.Done()
		for range oddCh {
			fmt.Print(1) // 此处实际应打印具体奇数，但为了通用需知道当前数字
			// 改进：用共享变量传递具体数字（这里简化，仅演示顺序）
			// 正确实现应让 zero 协程把要打印的数字通过 channel 传过来
			// 但题目主要考察顺序控制，我们按标准模式实现，实际需传递数字
			zeroCh <- struct{}{}
		}
	}()

	// even 协程
	go func() {
		defer wg.Done()
		for range evenCh {
			fmt.Print(2) // 同上，简化
			zeroCh <- struct{}{}
		}
	}()

	zeroCh <- struct{}{} // 启动 zero 协程
	wg.Wait()
	fmt.Println()
}

// 修正版 zeroOddEven：传递具体数字，按真实 0,1,0,2,0,3... 打印
func zeroOddEvenCorrect(n int) {
	fmt.Printf("=== 打印 0 与奇偶数（正确版，n=%d） ===\n", n)

	signal := make(chan struct{})
	oddCh := make(chan int)  // 传递要打印的奇数
	evenCh := make(chan int) // 传递要打印的偶数

	var wg sync.WaitGroup
	wg.Add(3)
	// zero 协程
	go func() {
		defer wg.Done()
		for i := 1; i <= n; i++ {
			<-signal // 等待启动信号
			fmt.Print(0)
			if i%2 == 1 {
				oddCh <- i
			} else {
				evenCh <- i
			}
		}

		close(oddCh)
		close(evenCh)
	}()

	// odd 协程
	go func() {
		defer wg.Done()
		for num := range oddCh {
			fmt.Print(num)
			if num < n {
				signal <- struct{}{}
			}
		}
	}()

	// even 协程
	go func() {
		defer wg.Done()
		for num := range evenCh {
			fmt.Print(num)
			if num < n {
				signal <- struct{}{}
			}
		}
	}()

	signal <- struct{}{} // 启动 zero
	wg.Wait()
	fmt.Println()
}

func main() {
	// 1. 交替打印数字和字母
	alternatePrint()

	// 2. 睡眠排序
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
	sleepSort(arr)

	// 3. 打印 0 与奇偶数
	zeroOddEvenCorrect(10)
}
