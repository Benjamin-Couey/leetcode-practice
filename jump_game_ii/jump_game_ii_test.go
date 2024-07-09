package jump_game_ii

import (
	"testing"
)

func TestJump(t *testing.T) {
	testcases := []struct {
		nums []int
		solution int
	}{
		{ []int{ 2, 3, 1, 1, 4 }, 2 },
		{ []int{ 2, 3, 0, 1, 4 }, 2 },
		{ []int{ 1, 2, 3, 4, 5 }, 3 },
		{ []int{ 1, 2, 1, 4, 1, 1, 1, 5 }, 3 },
		{ []int{ 3, 2, 1, 1, 3, 1, 1, 5 }, 3 },
		{ []int{ 1, }, 0 },
	}

	for _, testcase := range testcases {
		result := Jump( testcase.nums )

		if result != testcase.solution {
			t.Errorf(
				"Jump: %v returned %v, want %v",
				testcase.nums,
				result,
				testcase.solution,
			)
		}
	}
}
