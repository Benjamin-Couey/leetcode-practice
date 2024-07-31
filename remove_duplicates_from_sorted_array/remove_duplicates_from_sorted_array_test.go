package remove_duplicates_from_sorted_array

import (
	"leetcode/utils"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testcases := []struct {
		nums, solution []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
	}

	for _, testcase := range testcases {
		k := RemoveDuplicates(testcase.nums)

		if k != len(testcase.solution) {
			t.Errorf("Remove element: Returned %v, want %v", k, len(testcase.solution))
		}

		first_k := testcase.nums[0:k]
		if !utils.SliceEqual(first_k, testcase.solution) {
			t.Errorf("Remove element: %v, want %v", first_k, testcase.solution)
		}

	}
}
