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

}
