package gstack

import "fmt"

// 125. 验证回文串
// https://leetcode.cn/problems/valid-palindrome/description/
func isPalindrome(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		fmt.Println(string(ch), ch, byte('a'))
		if ch >= 'A' && ch <= 'Z' {
			ch = ch + 32
			stack = append(stack, ch)
		} else if ch >= 'a' && ch <= 'z' {
			stack = append(stack, ch)
		} else if ch >= '0' && ch <= '9' {
			stack = append(stack, ch)
		}
	}
	length := len(stack)
	ok := true
	for i := 0; i < length/2; i++ {
		if stack[i] != stack[length-i-1] {
			ok = false
			break
		}
	}
	return ok
}

func isPalindromev2(s string) bool {
	length := len(s)
	if length <= 1 {
		return true
	}
	i, j := 0, length-1
	for i <= j {
		for i < j && !isAlphanumeric(s[i]) {
			i++
		}
		for j >= i && !isAlphanumeric(s[j]) {
			j--
		}
		if i >= j {
			return true
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func toLower(r byte) byte {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func isAlphanumeric(r byte) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

type MinStack struct {
	Stack []int
	// Diff=Min-PrevMin
	Min  []int
	size int
}

// https://leetcode.cn/problems/min-stack/description/
// 155. 最小栈
func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		Min:   []int{},
	}

}

func (this *MinStack) Push(val int) {
	if len(this.Stack) == 0 {
		this.Stack = append(this.Stack, val)
		this.Min = append(this.Min, val)
	} else {
		min := this.Min[len(this.Min)-1]
		if val < min {
			min = val
		}
		this.Stack = append(this.Stack, val)
		this.Min = append(this.Min, min)
	}
	this.size += 1
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return
	} else {
		this.Stack = this.Stack[:this.size-1]
		this.Min = this.Min[:this.size-1]
		this.size -= 1
	}
}

func (this *MinStack) Top() int {
	if len(this.Stack) == 0 {
		return 0
	} else {
		return this.Stack[this.size-1]
	}

}

func (this *MinStack) GetMin() int {
	if len(this.Stack) == 0 {
		return 0
	}
	return this.Min[this.size-1]
}
