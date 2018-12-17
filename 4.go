package leetcode

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。
*/
func findMedianSortedArrays(nums1, nums2 []int) float64 {
	var p1, p2 *[]int
	if len(nums1) < len(nums2) {
		p1 = &nums2
		p2 = &nums1
	} else {
		p1 = &nums1
		p2 = &nums2
	}
	m := len(*p1)
	n := len(*p2)
	var i, j, leftLen int
	maxLeft, minRight := 1, 0
	leftLen = (m + n) >> 1
	if leftLen == 0 {
		return float64((*p1)[0])
	}
	i = m >> 1
	lmtLeft := leftLen - n
	lmtRight := leftLen
	for maxLeft > minRight {
		j = leftLen - i
		if j == n || ((i < m) && (*p1)[i] < (*p2)[j]) {
			minRight = (*p1)[i]
		} else {
			minRight = (*p2)[j]
		}
		if j == 0 || (i >= 1 && (*p1)[i-1] > (*p2)[j-1]) {
			maxLeft = (*p1)[i-1]
			if i == lmtLeft+1 {
				i = lmtLeft
			} else {
				lmtRight = i
				i = (lmtLeft + i + 1) >> 1
			}
		} else {
			maxLeft = (*p2)[j-1]
			if i == lmtRight-1 {
				i = lmtRight
			} else {
				lmtLeft = i
				i = (i + lmtRight) >> 1
			}
		}
	}
	if leftLen*2 == m+n {
		return float64(maxLeft+minRight) / 2
	}
	return float64(minRight)
}
