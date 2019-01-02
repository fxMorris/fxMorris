package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"case 7",
			args{"ccc"},
			"ccc",
		},
		{
			"case 6",
			args{"abacab"},
			"bacab",
		},
		{
			"case 5",
			args{"a"},
			"a",
		},
		{
			"case 4",
			args{""},
			"",
		},
		{
			"case 3",
			args{"aaaabcbasssaa"},
			"abcba",
		},
		{
			"case 3",
			args{"abba"},
			"abba",
		},
		{
			"case 2",
			args{"ab"},
			"a",
		},
		{
			"case 1",
			args{"aba"},
			"aba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
