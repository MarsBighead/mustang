package gstack

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	s := "abbaca"
	s = removeDuplicates(s)
	fmt.Println(s)
}

func TestIsValid(t *testing.T) {
	s := "{[]}"
	//s = "()))"
	//s = "[[[]"
	//s = "()"
	//s = "()[]{}"
	ok := isValid(s)
	fmt.Println(ok)
}

func TestCalculte(t *testing.T) {
	s := " 0-2147483647"
	r := calculate(s)
	fmt.Println(s, r)
}

func TestCalcultev2(t *testing.T) {
	s := " 0-2147483647"
	r := calculatev2(s)
	fmt.Println(s, r)
}

func TestDailyTemperatures(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans := dailyTemperatures(temperatures)
	fmt.Println(ans)
}
