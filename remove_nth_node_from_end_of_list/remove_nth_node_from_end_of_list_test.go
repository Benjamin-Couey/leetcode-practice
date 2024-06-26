package remove_nth_node_from_end_of_list

import (
	"testing"
	"leetcode/utils"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testcases := []struct {
		head []int
		n int
		solution []int
	}{
		{ []int{ 1, 2, 3, 4, 5 }, 2, []int{ 1, 2, 3, 5 } },
		{ []int{ 1, 2 }, 1, []int{ 1 } },
		{ []int{ 1 }, 1, []int{} },
	}

	for _, testcase := range testcases {
		result := RemoveNthFromEnd( utils.SliceToLinkedList( testcase.head ), testcase.n )

		if !utils.IsSameLinkedList( result, utils.SliceToLinkedList( testcase.solution ) ) {
			t.Errorf(
				"RemoveNthFromEnd: %v and %v returned %v want %v",
				testcase.head,
				testcase.n,
				utils.LinkedListToSlice( result ),
				testcase.solution,
			)
		}
	}
}
