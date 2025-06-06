package gstack

import (
	"fmt"
	"strconv"
	"strings"
)

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

func reverseWords(s string) string {
	if len(s) <= 1 {
		return s
	}
	words := strings.Split(s, " ")
	//fmt.Println(words)
	for j, word := range words {
		length := len(word)
		stack := []byte(word)
		for i := 0; i < length/2; i++ {
			stack[i], stack[length-i-1] = stack[length-i-1], stack[i]
		}
		//fmt.Println(string(stack))
		words[j] = string(stack)

	}

	return strings.Join(words, " ")
}

func addStrings(num1 string, num2 string) string {
	length := len(num1)
	if len(num2) < length {
		length = len(num2)
		num1, num2 = num2, num1
	}
	count := len(num2)
	carry := 0
	stack := make([]string, count)
	for i := 0; i < length; i++ {
		sum := int(num1[length-i-1]-'0') + int(num2[count-i-1]-'0') + carry
		sum, carry = sum%10, sum/10
		fmt.Println(sum)
		stack[count-i-1] = fmt.Sprintf("%d", sum)
	}
	fmt.Println("pre stack=", stack, carry)
	for i := 0; i < count-length; i++ {
		sum := int(num2[count-length-i-1]-'0') + carry
		sum, carry = sum%10, sum/10
		stack[count-length-i-1] = fmt.Sprintf("%d", sum)
	}
	fmt.Println("stack=", stack, carry)
	suffix := strings.Join(stack, "")
	if carry == 0 {
		return suffix
	}
	return fmt.Sprintf("%d", carry) + suffix
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	length := len(num1)
	if len(num2) < length {
		length = len(num2)
		num1, num2 = num2, num1
	}
	count := len(num2)
	stack := make([]string, length)
	for i := length - 1; i >= 0; i-- {
		a := int(num1[i] - '0')
		carry := 0
		//result := ""
		array := make([]string, count)
		for j := count - 1; j >= 0; j-- {
			b := int(num2[j] - '0')
			multi := a*b + carry
			multi, carry = multi%10, multi/10
			array[j] = fmt.Sprintf("%d", multi)
		}
		val := strings.Join(array, "")
		if carry > 0 {
			val = fmt.Sprintf("%d", carry) + val
		}
		fmt.Println("val=", val)
		nZero := length - i - 1
		if nZero > 0 {
			zeros := make([]string, nZero)
			for i := 0; i < nZero; i++ {
				zeros[i] = "0"
			}
			val = val + strings.Join(zeros, "")
		}
		stack[i] = val
	}
	fmt.Println(stack)
	if len(stack) == 1 {
		return stack[0]
	}
	str := stack[0]
	for i := 1; i < len(stack); i++ {
		str = addStrings(str, stack[i])
		fmt.Println(str)
	}
	return str
}
