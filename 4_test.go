package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		//TODO: Add test cases.
		{
			"test case 8",
			args{[]int{1}, []int{2, 3, 4, 5, 6, 7, 8, 9}},
			float64(5),
		},
		{
			"test case 7",
			args{[]int{1, 2, 5}, []int{3, 4, 6, 7, 8, 9}},
			float64(5),
		},
		{
			"test case 7",
			args{[]int{1}, []int{2, 3, 4, 5, 6, 7}},
			float64(4),
		},
		{
			"test case 7",
			args{[]int{4}, []int{1, 2, 3, 5, 6}},
			float64(3.5),
		},
		{
			"test case 6",
			args{[]int{}, []int{1}},
			float64(1),
		},
		{
			"test case 1",
			args{[]int{1, 3}, []int{2}},
			float64(2),
		},
		{
			"test case 2",
			args{[]int{1, 2}, []int{3, 4}},
			float64(2.5),
		},
		{
			"test case 3",
			args{[]int{1, 3}, []int{2, 5, 6, 7}},
			float64(4),
		},
		{
			"test case 4",
			args{[]int{3, 4}, []int{1, 2}},
			float64(2.5),
		},
		{
			"test case 5",
			args{[]int{2, 3}, []int{}},
			float64(2.5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
