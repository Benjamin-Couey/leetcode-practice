package reverse_linked_list_ii

import (
	"testing"
	"leetcode/utils"
)

func TestReverseBetween(t *testing.T) {
	testcases := []struct {
		head []int
		left, right int
		solution []int
	}{
		{ []int{ 1, 2, 3, 4, 5 }, 2, 4, []int{ 1, 4, 3, 2, 5 } },
		{ []int{ 1, 2, 3, 4, 5 }, 1, 3, []int{ 3, 2, 1, 4, 5 } },
		{ []int{ 1, 2, 3, 4, 5 }, 1, 5, []int{ 5, 4, 3, 2, 1 } },
		{ []int{ 5 }, 1, 1, []int{ 5 } },
	}

	for _, testcase := range testcases {
		result := ReverseBetween( utils.SliceToLinkedList( testcase.head ), testcase.left, testcase.right )

		if !utils.IsSameLinkedList( result, utils.SliceToLinkedList( testcase.solution ) ) {
			t.Errorf(
				"ReverseBetween: reversing %v between %v and %v returned %v want %v",
				testcase.head,
				testcase.left,
				testcase.right,
				utils.LinkedListToSlice( result ),
				testcase.solution,
			)
		}
	}
}
