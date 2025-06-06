package gstring

// 1668. 最大重复子字符串
// https://leetcode.cn/problems/maximum_repeating_substring/description/
func maxRepeating(sequence string, word string) int {
	m, n := len(sequence), len(word)
	if n > m {
		return 0
	}
	if n == m && sequence == word {
		return 1
	}
	//fmt.Printf("m=%d, n=%d\n", m, n)
	s := word
	k := 0
	for i := 0; i <= m-n; i++ {
		//fmt.Printf("idx=%d, k=%d in: %s\n", i, k, sequence[i:i+n])
		if sequence[i:i+n] == s {
			k++
			s = s + word
			i--
			n = len(s)
		}
	}

	return k
}

func maxRepeatingv1(sequence string, word string) int {
	m, n := len(sequence), len(word)
	if n > m {
		return 0
	}
	if n == m && sequence == word {
		return 1
	}

	k, i, j := 0, 0, 0
	up := n

	for ; i <= m-up; i++ {
		j = 0
		for ; j <= k; j++ {
			if i+(j+1)*n > m {
				break
			}
			if sequence[i+j*n:i+(j+1)*n] != word {
				break
			}
		}
		if j == k+1 {
			k++
			i--
		}
	}

	return k
}
