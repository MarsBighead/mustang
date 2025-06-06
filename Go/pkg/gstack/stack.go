package gstack

import (
	"fmt"
	"strings"
)

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
