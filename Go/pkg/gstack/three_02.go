package gstack

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/basic-calculator/description/?envType=study-plan-v2&envId=top-interview-150
// 224. 基本计算器
func calculatev0(s string) int {
	// stack 存储数字前面的符号
	stack := []int{1}
	ans := 0
	sign := 1
	n := len(s)

	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return ans
}

// https://leetcode.cn/problems/basic-calculator-ii/description/
// 227. 基本计算器 II
func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	tokens := make([]string, 0)
	//current := ""
	length := len(s)
	fmt.Printf("s=%s\n", s)
	var isPre bool
	for i := 0; i < length; i++ {
		t := string(s[length-i-1])
		switch t {
		case "*", "/", "-", "+":
			tmp := s[length-i:]
			tokens = append(tokens, tmp, t)
			s = s[:length-i-1]
			//fmt.Printf("x: n=%d, t=%s,tmp=%s,length: %d\n", i, t, tmp, length)
			length = len(s)
			i = 0
			if t == "*" || t == "/" {
				isPre = true
			}
		}
	}
	tokens = append(tokens, s)
	var (
		stack  []string
		result int
	)
	if len(tokens) == 1 {
		result, _ := strconv.Atoi(tokens[0])
		return result
	}
	fmt.Printf("tokens=%v\n", tokens)

	if isPre {
		stack = append(stack, tokens[len(tokens)-1])
		tokens = tokens[:len(tokens)-1]
		length = len(tokens)
		for len(tokens) > 0 {
			t := tokens[length-1]
			if t == "*" || t == "/" {
				a, _ := strconv.Atoi(stack[len(stack)-1])
				//a := 1
				b, _ := strconv.Atoi(tokens[length-2])
				stack[len(stack)-1] = fmt.Sprintf("%d", cal(a, b, t))
				tokens = tokens[:length-2]
				//fmt.Printf("t=%s,a=%d,b=%d,length: %d\n", t, a, b, len(tokens))
			} else {
				stack = append(stack, t)
				tokens = tokens[:length-1]
			}
			length = len(tokens)
		}
	} else {
		length = len(tokens)
		for len(tokens) > 0 {
			t := tokens[length-1]
			stack = append(stack, t)
			tokens = tokens[:length-1]
			length = len(tokens)
		}
	}
	fmt.Println("stack=", stack)
	if len(stack) > 1 {
		length = len(stack)
		a, _ := strconv.Atoi(stack[0])
		fmt.Println("a vs b:", a)
		for i := 1; i < length; i++ {
			if stack[i] == "+" || stack[i] == "-" {
				//fmt.Println("a", result, a)
				b, _ := strconv.Atoi(stack[i+1])
				fmt.Println("a vs b:", a, b)
				a = cal(a, b, stack[i])
				i++
			}
		}
		result = a
	} else {
		result, _ = strconv.Atoi(stack[0])
	}

	return result
}

func cal(a, b int, char string) int {
	switch char {
	case "*":
		return a * b
	case "/":
		return a / b
	case "+":
		return a + b
	case "-":
		return a - b
	}
	return 0
}

func calculatev2(s string) int {
	stack := []int{}
	sign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			sign = ch
			num = 0
		}
	}
	result := 0
	for _, v := range stack {
		result += v
	}
	//fmt.Println("result=", stack)
	return result
}
