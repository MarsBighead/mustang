package hash

import (
	"strconv"
	"strings"
)

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
