package merge_two_sorted_lists

import (
	"testing"
	"leetcode/utils"
)

func TestMergeTwoLists(t *testing.T) {
	testcases := []struct {
		list_1, list_2 []int
		solution []int
	}{
		{ []int{ 1, 2, 4 }, []int{ 1, 3, 4 }, []int{ 1, 1, 2, 3, 4, 4 } },
		{ []int{}, []int{}, []int{} },
		{ []int{}, []int{ 0 }, []int{ 0 } },
	}

	for _, testcase := range testcases {
		result := MergeTwoLists(
			utils.SliceToLinkedList( testcase.list_1),
			utils.SliceToLinkedList( testcase.list_2),
		)

		if !utils.IsSameLinkedList( result, utils.SliceToLinkedList( testcase.solution ) ) {
			t.Errorf(
				"MergeTwoLists: did not correctly merge lists %v and %v",
				testcase.list_1,
				testcase.list_2,
			)
		}
	}
}
