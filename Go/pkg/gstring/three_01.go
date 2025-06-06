package gstring

import "fmt"

// https://leetcode.cn/problems/repeated_dna_sequences/description/
// 187. 重复的DNA序列
func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	fmt.Println("length of DNA s:", n)
	if n <= 10 {
		return []string{}
	}
	ans := []string{}
	seq := s[:10]
	hash := map[string]int{
		seq: 1,
	}
	fmt.Println(len(seq))
	for i := 1; i <= n-10; i++ {
		seq = s[i : i+10]
		hash[seq]++
		if hash[seq] == 2 {
			ans = append(ans, seq)
		}
	}
	return ans
}
