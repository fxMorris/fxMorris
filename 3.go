package leetcode

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
*/
func lengthOfLongestSubstring(s string) int {
	re, tmp := 0, 1
	chs := []byte(s)
	for i, v := range []byte(s) {
		if len(s)-i < re {
			break
		}
		tmp = 1
		tmpChs := make(map[byte]int)
		tmpChs[v] = 1
		for j := i + 1; j < len(s); j++ {
			ch := chs[j]
			if _, ok := tmpChs[ch]; ok {
				break
			} else {
				tmpChs[ch] = 1
				tmp++
			}
		}
		if tmp > re {
			re = tmp
		}
	}
	return re
}
