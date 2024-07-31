package remove_duplicates_from_sorted_array_ii

import (
	"leetcode/utils"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testcases := []struct {
		nums, solution []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, []int{1, 1, 2, 2, 3}},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, []int{0, 0, 1, 1, 2, 3, 3}},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3, 4}, []int{0, 0, 1, 1, 2, 3, 3, 4}},
	}

	for _, testcase := range testcases {
		k := RemoveDuplicates(testcase.nums)

		if k != len(testcase.solution) {
			t.Errorf("RemoveDuplicates: returned %v, want %v", testcase.nums, testcase.solution)
		}

		first_k := testcase.nums[0:k]
		if !utils.SliceEqual(first_k, testcase.solution) {
			t.Errorf("RemoveDuplicates: returned %v, want %v", first_k, testcase.solution)
		}

	}
}
