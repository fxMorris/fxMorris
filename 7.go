package leetcode

import "strconv"

func reverse(x int) int {
	s := strconv.Itoa(x)
	isNegative := false
	if s[0] == '-' {
		s = s[1:len(s)]
		isNegative = true
	}
	n := len(s)
	chs := []byte(s)
	reversedChs := make([]byte, n)
	for i, v := range chs {
		if v == byte('-') {
			reversedChs[i] = v
		}
		reversedChs[n-i-1] = v
	}
	result, err := strconv.Atoi(string(reversedChs))
	if err != nil {
		result = 0
	}
	if isNegative {
		result = -result
	}
	if result != int(int32(result)) {
		result = 0
	}
	return result
}
