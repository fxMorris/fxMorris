package leetcode

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
*/
func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	m := make(map[int]int)
	c := 0   //最长回文子串中心
	len := 1 //覆盖半径
	m[0] = 1
	for i := 1; i < 2*n-1; i++ {
		var l, r int
		l = i >> 1
		if float64(i>>1) == float64(i)/2 {
			r = l + 1
			l = l - 1
		} else {
			r = l + 1
		}
		tmp := 0
		if i < c+len {
			if i+m[2*c-i]>>1 < c+len {
				tmp = m[2*c-i] >> 1
			} else {
				tmp = c + len - i - 1
			}
		}
		for l-tmp >= 0 && r+tmp <= n-1 && s[l-tmp] == s[r+tmp] {
			tmp++
		}
		m[i] = tmp*2 + r - l - 1
		if m[i] > len {
			len = m[i]
			c = i
		}
	}
	t := len >> 1
	left := (c+1)>>1 - t
	return s[left : left+len]
}
