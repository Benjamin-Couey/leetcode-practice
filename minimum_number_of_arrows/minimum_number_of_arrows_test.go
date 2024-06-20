package minimum_number_of_arrows

import (
	"testing"
)

func TestFindMinArrowShots(t *testing.T) {
	testcases := []struct {
		points [][]int
		solution int
	}{
		{ [][]int{ { 10, 16 }, { 2, 8, },	{ 1, 6, }, { 7, 12, }, }, 2 },
		{ [][]int{ { 1, 2 }, { 3, 4, },	{ 5, 6, }, { 7, 8, }, }, 4 },
		{ [][]int{ { 1, 2 }, { 2, 3, },	{ 3, 4, }, { 4, 5, }, }, 2 },
		{ [][]int{ { 1, 2 }, }, 1 },
	}

	for _, testcase := range testcases {
		result := FindMinArrowShots( testcase.points )

		if result != testcase.solution {
			t.Errorf(
				"FindMinArrowShots: %v returned %v, want %v",
				testcase.points,
				result,
				testcase.solution,
			)
		}
	}
}
