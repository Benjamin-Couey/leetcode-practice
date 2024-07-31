package remove_element

import (
	"leetcode/utils"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	testcases := []struct {
		nums, solution []int
		val            int
	}{
		{[]int{3, 2, 2, 3}, []int{2, 2}, 3},
		{[]int{2, 3, 3, 3}, []int{2}, 3},
		{[]int{3, 3, 3, 2}, []int{2}, 3},
	}

	for _, testcase := range testcases {
		k := RemoveElement(testcase.nums, testcase.val)

		if k != len(testcase.solution) {
			t.Errorf("Remove element: Returned %v, want %v", k, len(testcase.solution))
		}

		utils.SortFirstKInts(testcase.nums, k)
		first_k := testcase.nums[0:k]
		if !utils.SliceEqual(first_k, testcase.solution) {
			t.Errorf("Remove element: %v, want %v", first_k, testcase.solution)
		}

	}
}
