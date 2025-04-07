package main

func fn(n int) int {
	var arr [100000]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	return fn(n-1) + 1
}
