package merge_sorted_array

import (
	"leetcode/utils"
	"testing"
)

func TestMerge(t *testing.T) {
	testcases := []struct {
		nums1, nums2, solution []int
		m, n                   int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}, []int{1, 2, 2, 3, 5, 6}, 3, 3},
		{[]int{1}, []int{}, []int{1}, 1, 0},
		{[]int{0}, []int{1}, []int{1}, 0, 1},
	}

	for _, testcase := range testcases {
		Merge(testcase.nums1, testcase.m, testcase.nums2, testcase.n)

		if !utils.SliceEqual(testcase.nums1, testcase.solution) {
			t.Errorf("Merge: %v, want %v", testcase.nums1, testcase.solution)
		}
	}
}
