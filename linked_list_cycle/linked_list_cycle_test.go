package linked_list_cycle

import (
	"testing"
	"leetcode/utils"
)

func TestHasCycle(t *testing.T) {
	testcases := []struct {
		list []int
		cycle_index int
		solution bool
	}{
		{ []int{ 3, 2, 0, -4 }, 1, true },
		{ []int{ 1, 2 }, 0, true },
		{ []int{ 1 }, -1, false },
	}

	for _, testcase := range testcases {
		result := HasCycle( utils.SliceToLinkedListWithCycle( testcase.list, testcase.cycle_index ) )

		if result != testcase.solution {
			t.Errorf(
				"HasCycle: list %v with cycle index %v returned %v, want %v",
				testcase.list,
				testcase.cycle_index,
				result,
				testcase.solution,
			)
		}
	}
}
