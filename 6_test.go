package leetcode

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"case 4",
			args{"A", 1},
			"A",
		},
		{
			"case 3",
			args{"", 4},
			"",
		},
		{
			"case 2",
			args{"LEETCODEISHIRING", 4},
			"LDREOEIIECIHNTSG",
		},
		{
			"case 1",
			args{"LEETCODEISHIRING", 3},
			"LCIRETOESIIGEDHN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
