package main

import "C"
import (
	"fmt"
)

//export printMessage
func printMessage() {
	fmt.Println("A go funcation!")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

func main() {}
