package gstring

// 6. Z 字形变换
// https://leetcode.cn/problems/zigzag_conversion/description
/*
方法三：直接构造
我们来研究方法一中矩阵的每个非空字符会对应到 s 的哪个下标（记作 idx），从而直接构造出答案。
由于 Z 字形变换的周期为 t=2r−2，因此对于矩阵第一行的非空字符，其对应的 idx 均为 t 的倍数，即 idx≡0(modt)；同理，对于矩阵最后一行的非空字符，应满足 idx≡r−1(modt)。
对于矩阵的其余行（行号设为 i），每个周期内有两个字符，第一个字符满足 idx≡i(modt)，第二个字符满足 idx≡t−i(modt)
*/
func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || numRows >= n {
		return s
	}
	t := numRows*2 - 2
	result := make([]byte, 0, n)
	for i := 0; i < numRows; i++ { // 枚举矩阵的行
		for j := 0; j+i < n; j += t { // 枚举每个周期的起始下标
			result = append(result, s[j+i]) // 当前周期的第一个字符
			if i > 0 && i < numRows-1 && j+t-i < n {
				result = append(result, s[j+t-i]) // 当前周期的第二个字符
			}
		}
	}
	return string(result)

}
