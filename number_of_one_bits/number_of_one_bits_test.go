package number_of_one_bits

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	testcases := []struct {
		n, solution int
	}{
		{11, 3},
		{128, 1},
		{2147483645, 30},
	}

	for _, testcase := range testcases {
		result := HammingWeight(testcase.n)

		if result != testcase.solution {
			t.Errorf(
				"HammingWeight: %v returned %v, want %v",
				testcase.n,
				result,
				testcase.solution,
			)
		}
	}
}
