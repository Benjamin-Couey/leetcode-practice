package add_two_numbers

import (
	"leetcode/utils"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	testcases := []struct {
		l1, l2   []int
		solution []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, testcase := range testcases {
		result := AddTwoNumbers(utils.SliceToLinkedList(testcase.l1), utils.SliceToLinkedList(testcase.l2))

		if !utils.IsSameLinkedList(result, utils.SliceToLinkedList(testcase.solution)) {
			t.Errorf(
				"AddTwoNumbers: %v and %v returned %v want %v",
				testcase.l1,
				testcase.l2,
				utils.LinkedListToSlice(result),
				testcase.solution,
			)
		}
	}
}
