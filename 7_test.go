package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"case 5",
			args{1534236469},
			0,
		},
		{
			"case 4",
			args{-123},
			-321,
		},
		{
			"case 3",
			args{1234567899999999999},
			0,
		},
		{
			"case 2",
			args{0},
			0,
		},
		{
			"case 1",
			args{123},
			321,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
