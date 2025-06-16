package garray

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	numbers := [][]int{{1, 2, 3}, {3, 2, 1}, {1, 1, 5}}
	for _, nums := range numbers {
		NextPermutation(nums)
		fmt.Println(nums)
	}
}
func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	ans := CombinationSum(candidates, target)
	fmt.Printf("%#v\n", ans)
}

func TestIsValidSudoku(t *testing.T) {
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	ok := isValidSudoku(board)
	fmt.Println("ok?=", ok)

}
func TestCombinationSum2(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	ans := CombinationSum2(candidates, target)
	fmt.Printf("%#v\n", ans)
}

func TestFirstMissingPositive(t *testing.T) {
	nums := []int{1, 2, 0}
	n := FirstMissingPositive(nums)
	fmt.Printf("output %#v\n", n)
	nums = []int{3, 4, -1, 1}
	n = FirstMissingPositiveV1(nums)
	fmt.Printf("output %#v\n", n)
}

func TestUniquePaths(t *testing.T) {
	m, n := 7, 3
	r := UniquePaths(m, n)
	fmt.Println("result=", r)
	r = UniquePathsV1(m, n)
	fmt.Println("result=", r)
}

func TestSpiralOrder(t *testing.T) {
	mx := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	nums := spiralOrder(mx)
	fmt.Println(nums)
}
func TestCanJump(t *testing.T) {
	mat := [][]int{
		{2, 3, 1, 1, 4},
		{2, 3, 1, 1, 0},
		{3, 2, 1, 0, 4},
	}
	for _, nums := range mat {
		ok := canJump(nums)
		fmt.Println(ok)
	}
}
func TestUniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	r := UniquePathsWithObstacles(obstacleGrid)
	fmt.Println("result=", r)
	r = UniquePathsWithObstaclesV1(obstacleGrid)
	fmt.Println("result=", r)

}
