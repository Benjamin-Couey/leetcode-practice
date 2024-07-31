package climbing_stairs

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	testcases := []struct {
		n, solution int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
	}

	for _, testcase := range testcases {
		result := ClimbStairs(testcase.n)
		if result != testcase.solution {
			t.Errorf(
				"ClimbStairs: %v returned %v, want %v",
				testcase.n,
				result,
				testcase.solution,
			)
		}
	}
}
