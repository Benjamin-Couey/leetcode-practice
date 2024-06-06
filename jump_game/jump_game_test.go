package jump_game

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	testcases := []struct {
		nums []int
		solution bool
	}{
		{ []int{ 2, 3, 1, 1, 4 }, true },
		{ []int{ 3, 2, 1, 0, 4 }, false },
		{ []int{ 0, 3, 1, 1, 4 }, false },
		{ []int{ 0, }, true },
		{ []int{ 4, 0, 0, 0, 0, }, true },
	}

	for _, testcase := range testcases {
		result := CanJump( testcase.nums )

		if result != testcase.solution {
			t.Errorf(
				"CanJump: %v returned %v, want %v",
				testcase.nums,
				result,
				testcase.solution,
			)
		}
	}
}
