package gstring

import "fmt"

// 459. 重复的子字符串
// https://leetcode.cn/problems/repeated-substring-pattern/description/
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n <= 1 {
		return false
	}
	ds := s + s
	ok := false
	for i := 1; i < n; i++ {
		if ds[i:i+n] == s {
			ok = true
			fmt.Printf("%d, n=%d, sub string: %s\n", i, n, ds[i:i+n])
			break
		}
	}

	return ok
}

func stringMatching(words []string) []string {
	n := len(words)
	if n <= 1 {
		return nil
	}
	min := len(words[0])
	max := min
	hash := make(map[int][]string)
	for _, word := range words {
		m := len(word)
		if max < m {
			max = m
		}
		if min > m {
			min = m
		}
		hash[m] = append(hash[m], word)
	}
	lengths := []int{}
	for i := min; i <= max; i++ {
		if _, ok := hash[i]; ok {
			lengths = append(lengths, i)
		}
	}
	ans := []string{}
	for i, length := range lengths {
		arr := hash[length]
		if len(arr[0]) == max {
			break
		}
		ok := false
		for _, w := range arr {
			for j := i + 1; j < len(lengths); j++ {
				array := hash[lengths[j]]
				ok = matching(w, array)
				if ok {
					ans = append(ans, w)
					break
				}
			}
		}
	}
	return ans
}

func matching(w string, words []string) bool {
	length := len(w)
	ok := false
	for _, word := range words {
		for i := 0; i <= len(word)-length; i++ {
			if word[i:i+length] == w {
				ok = true
				break
			}
		}
		if ok {
			break
		}
	}
	return ok
}

func stringMatchingv1(words []string) []string {
	n := len(words)
	if n <= 1 {
		return nil
	}

	max := len(words[0])
	hash := make(map[int][]string)
	for _, word := range words {
		m := len(word)
		if max < m {
			max = m
		}
		hash[m] = append(hash[m], word)
	}
	ans := []string{}

	for i, arr := range hash {
		fmt.Printf("max=%d, i=%d\n", max, i)
		if i == max {
			continue
		}
		for _, w := range arr {
			ok := false
			for j, arrary := range hash {
				if j <= i {
					continue
				}
				ok = matching(w, arrary)
				if ok {
					ans = append(ans, w)
					break
				}
			}
		}
	}
	return ans
}
