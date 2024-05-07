package gstack

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

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

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
			fmt.Println("stack=", stack)
		}
		stack = append(stack, i)
		fmt.Println("ans=", i, stack)
	}
	return ans
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			hash[num] = stack[len(stack)-1]
		} else {
			hash[num] = -1
		}

		stack = append(stack, num)
	}
	for i, num := range nums1 {
		nums1[i] = hash[num]
	}
	return nums1
}

func nextGreaterElementv2(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		hash[nums2[i]] = i
	}
	for i, num := range nums1 {
		index := hash[num]
		found := false
		for j := index + 1; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				nums1[i] = nums2[j]
				found = true
				break
			}
		}
		if !found {
			nums1[i] = -1
		}

	}
	return nums1
}

func simplifyPath(path string) string {
	stack := []string{}
	for _, dir := range strings.Split(path, "") {
		if dir == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

		} else if dir != "" && dir != "." {
			stack = append(stack, dir)
		}
	}
	return "/" + strings.Join(stack, "/")
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

// https://leetcode.cn/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	count := len(strs)
	if count == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < count; i++ {
		prefix := lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := len(str1)
	if len(str2) < length {
		length = len(str2)
	}
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
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

// https://leetcode.cn/problems/group-anagrams/description/
func groupAnagrams(strs []string) [][]string {
	mp := make(map[string]int)
	ans := [][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		if idx, ok := mp[sortedStr]; ok {
			ans[idx] = append(ans[idx], str)
		} else {
			ans = append(ans, []string{str})
			mp[sortedStr] = len(ans) - 1
		}
	}

	return ans
}

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		// 只delete一个，后面的可以跳过已经筛选的字串，直接奔向出现重复之后的字串
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

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
