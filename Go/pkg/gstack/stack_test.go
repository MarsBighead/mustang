package gstack

import (
	"fmt"
	"strings"
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

func TestMultiply(t *testing.T) {
	num1 := "123"
	num2 := "456"
	ans := multiply(num1, num2)
	fmt.Println(ans)
}

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	ans := isPalindrome(s)
	fmt.Println(ans)
}

func TestSimplifyPath(t *testing.T) {
	s := "///"
	fmt.Printf("path: %#v\n", strings.Split(s, "/"))
	path := simplifyPath(s)
	fmt.Println(path)
}
