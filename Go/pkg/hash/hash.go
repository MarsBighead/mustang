package hash

import (
	"strconv"
	"strings"
)

const base = 1024

type MyHashSet struct {
	data [][]int
}

func Constructor() MyHashSet {
	return MyHashSet{make([][]int, base)}

}

func (this *MyHashSet) Add(key int) {
	idx := this.hash(key)
	if !this.Contains(key) {
		this.data[idx] = append(this.data[idx], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	idx := this.hash(key)
	for i := 0; i < len(this.data[idx]); i++ {
		if key == this.data[idx][i] {
			suffix := this.data[idx][i+1:]
			this.data[idx] = this.data[idx][:i]
			this.data[idx] = append(this.data[idx], suffix...)
		}
	}

}

func (this *MyHashSet) hash(key int) int {
	return key % base
}

func (this *MyHashSet) Contains(key int) bool {
	idx := this.hash(key)
	for i := 0; i < len(this.data[idx]); i++ {
		if key == this.data[idx][i] {
			return true
		}
	}
	return false
}

// 811. 子域名访问计数
// https://leetcode.cn/problems/subdomain-visit-count/description/
// 分割的方法，无法处理数量大的情况；采用byte索引才能有效解决该问题
func subdomainVisits(cpdomains []string) []string {
	hash := make(map[string]int)
	for _, ss := range cpdomains {
		idx := strings.IndexByte(ss, ' ')
		cnt, _ := strconv.Atoi(ss[:idx])
		for ; idx < len(ss); idx++ {
			if ss[idx] == ' ' || ss[idx] == '.' {
				hash[ss[idx+1:]] += cnt
			}
		}
	}

	ans := make([]string, 0, len(hash))
	for domain, cnt := range hash {
		ans = append(ans, strconv.Itoa(cnt)+" "+domain)
	}
	return ans

}
