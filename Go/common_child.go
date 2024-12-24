package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
Common Child

https://www.hackerrank.com/challenges/common-child/problem
*/

// Complete the commonChild function below.

func commonChild(s1 string, s2 string) int32 {
	length := len(s2)
	matrix := make([][]int32, length+1)
	//matrix[length][length] = 0
	for i := 0; i <= length; i++ {
		matrix[i] = make([]int32, length+1)
	}
	a := []byte(s1)
	b := []byte(s2)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if a[i] == b[j] {
				matrix[i+1][j+1] = matrix[i][j] + 1
			} else {
				matrix[i+1][j+1] = max(matrix[i+1][j], matrix[i][j+1])
			}
		}
		//fmt.Printf("%#v\n", matrix[i])
	}
	return matrix[length][length]
}
func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b

}

func common() {

	s1 := "SHINCHAN"

	s2 := "NOHARAAA"

	result := commonChild(s1, s2)

	fmt.Printf("result=%d\n", result)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
