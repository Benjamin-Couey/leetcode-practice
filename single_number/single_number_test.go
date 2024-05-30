package single_number

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	testcases := []struct {
		nums []int
		solution int
	}{
		{ []int{ 2, 2, 1 }, 1 },
		{ []int{ 4, 1, 2, 1, 2 }, 4 },
		{ []int{ 1 }, 1 },
	}

	for _, testcase := range testcases {
		result := SingleNumber( testcase.nums )

		if result != testcase.solution {
			t.Errorf(
				"SingleNumber: %v returned %v, want %v",
				testcase.nums,
				result,
				testcase.solution,
			)
		}
	}
}
