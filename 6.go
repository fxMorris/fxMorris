package leetcode

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
*/
func convert(s string, numRows int) string {
	n := len(s)
	if n == 0 || numRows == 1 {
		return s
	}
	chs := make([]byte, n)
	index := 0
	partNum := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for j := 0; j < n/partNum+1; j++ {
				index1 := i + j*partNum
				if index1 < n {
					chs[index] = s[index1]
					index++
				}
			}
		} else {
			delta := 2 * (numRows - i - 1)
			for j := 0; j < n/partNum+1; j++ {
				index1 := i + j*partNum
				index2 := index1 + delta
				if index1 < n {
					chs[index] = s[index1]
					index++
				}
				if index2 < n {
					chs[index] = s[index2]
					index++
				}
			}
		}
	}
	return string(chs)
}
