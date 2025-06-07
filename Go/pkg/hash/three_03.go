package hash

// https://leetcode.cn/problems/ransom-note/description
// 383. 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	m, n := len(ransomNote), len(magazine)
	if m > n {
		return false
	}
	hash := make(map[rune]int, 26)
	for _, e := range ransomNote {
		hash[e] += 1
	}
	for _, e := range magazine {
		if _, ok := hash[e]; ok {
			hash[e]--
		}
	}
	for _, val := range hash {
		if val > 0 {
			return false
		}
	}
	return false
}

func canConstructv1(ransomNote string, magazine string) bool {
	m, n := len(ransomNote), len(magazine)
	if m > n {
		return false
	}
	check := [26]int{}
	for _, e := range ransomNote {
		check[e-'a']++
	}
	for _, e := range magazine {
		check[e-'a']--
	}
	for _, val := range check {
		if val > 0 {
			return false
		}
	}
	return true
}
