package garray

import (
	"fmt"
	"testing"
)

func TestLifeOfGame(t *testing.T) {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	GameOfLife(board)
	for _, row := range board {
		fmt.Println(row)
	}
	x := 0x12345678

	fmt.Printf("%.8x\n", x>>8)
	//fmt.Printf("%.8x\n", x&0x000000ff)
	fmt.Printf("%x\n", (x>>0)&0x000000ff)
	fmt.Printf("%.8x\n", (x>>8)&0x000000ff)
	fmt.Printf("%.8x\n", (x>>16)&0x000000ff)
	fmt.Printf("%.8x\n", (x>>24)&0x000000ff)

}
