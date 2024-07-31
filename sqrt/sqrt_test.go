package sqrt

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	testcases := []struct {
		x, solution int
	}{
		{4, 2},
		{8, 2},
		{16, 4},
		{17, 4},
		{1, 1},
		{0, 0},
	}

	for _, testcase := range testcases {
		result := MySqrt(testcase.x)
		if result != testcase.solution {
			t.Errorf(
				"MySqrt: %v returned %v, want %v",
				testcase.x,
				result,
				testcase.solution,
			)
		}
	}
}
