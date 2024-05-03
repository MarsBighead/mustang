package gstack

import "fmt"

type MinStack struct {
	Stack []int
	// Diff=Min-PrevMin
	Min  []int
	size int
}

// https://leetcode.cn/problems/min-stack/description/
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

// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/description/
// 删除字符串中相邻重复的字符
func removeDuplicates(s string) string {
	if len(s) <= 1 {
		return s
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

// https://leetcode.cn/problems/valid-parentheses/description/
func isValid(s string) bool {

	n := len(s)
	if n%2 == 1 {
		return false
	}
	var (
		stack = []byte(s)
		top   = stack[n-1]
		tmp   = []byte{top}
	)

	var valid bool
	//fmt.Println(string(stack))
	stack = stack[:n-1]

	for i := n - 2; i >= 0; i-- {
		current := stack[len(stack)-1]
		if len(tmp) > 0 {
			top = tmp[len(tmp)-1]
			tmp = tmp[:len(tmp)-1]
			stack = stack[:len(stack)-1]
		} else {
			if len(stack) > 1 {
				top = current
				stack = stack[:len(stack)-1]
				current = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				i--
			} else {
				break
			}
		}
		fmt.Println("stack", string(stack), "tmp:", string(tmp))
		p := string([]byte{current, top})
		fmt.Printf("n=%d, %v, %v, pair: %s\n", i, string(stack), string(tmp), p)
		if !(p == "()" || p == "{}" || p == "[]") {
			tmp = append(tmp, top, current)
		}

	}
	fmt.Println("End:", string(stack), "tmp:", string(tmp))
	if len(tmp) == 0 && len(stack) == 0 {
		valid = true
	}
	return valid

}
