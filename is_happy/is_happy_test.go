package is_happy

import (
	"testing"
)

func TestIsHappy(t *testing.T) {
	testcases := []struct {
		n        int
		solution bool
	}{
		{19, true},
		{2, false},
	}

	for _, testcase := range testcases {
		result := IsHappy(testcase.n)

		if result != testcase.solution {
			t.Errorf(
				"IsHappy: %v returned %v, want %v",
				testcase.n,
				result,
				testcase.solution,
			)
		}
	}
}
