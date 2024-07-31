package two_sum

import (
	"leetcode/utils"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testcases := []struct {
		nums     []int
		target   int
		solution []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, testcase := range testcases {
		result := TwoSum(testcase.nums, testcase.target)

		if !utils.SliceEqual(result, testcase.solution) {
			t.Errorf(
				"TwoSum: %v and %v returned %v, want %v",
				testcase.nums,
				testcase.target,
				result,
				testcase.solution,
			)
		}
	}
}
