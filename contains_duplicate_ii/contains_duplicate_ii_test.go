package contains_duplicate_ii

import (
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	testcases := []struct {
		nums []int
		k int
		solution bool
	}{
		{ []int{ 1, 2, 3, 1 }, 3, true },
		{ []int{ 1, 0, 1, 1 }, 1, true },
		{ []int{ 1, 2, 3, 1, 2, 3 }, 2, false },
	}

	for _, testcase := range testcases {
		result := ContainsNearbyDuplicate( testcase.nums, testcase.k )

		if result != testcase.solution {
			t.Errorf(
				"ContainsNearbyDuplicate: %v and %v returned %v, want %v",
				testcase.nums,
				testcase.k,
				result,
				testcase.solution,
			)
		}
	}
}
