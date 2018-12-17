package leetcode

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
*/
func longestPalindrome(s string) string {
	l := len(s)
	sl := 0
	var subStr string
	var lmt, tmp int
	for i := 0; i < l; i++ {
		if i >= (l >> 1) {
			lmt = l - i
		} else {
			lmt = i
		}
		tmp = 0
		for j := 1; j <= lmt; j++ {
			if s[i-j] == s[i+j] {
				tmp++
			} else {

			}
		}
	}
}
