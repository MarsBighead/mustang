package gstack

// 3. 单调栈模板
// 3.1 单调递增栈模板
func monotoneIncreasingStack(nums []int) []int {
	stack := []int{}
	for _, num := range nums {
		n := len(stack)
		for n > 0 && num >= stack[n-1] {
			stack = stack[:n-1]
		}
		stack = append(stack, num)
	}
	return stack
}

// 3.2 单调递减栈模板
func monotoneDecreasingStack(nums []int) []int {
	stack := []int{}
	for _, num := range nums {
		n := len(stack)
		for n > 0 && num <= stack[n-1] {
			stack = stack[:n-1]
		}
		stack = append(stack, num)
	}
	return stack
}
