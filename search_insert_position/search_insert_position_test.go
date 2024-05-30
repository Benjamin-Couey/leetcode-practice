package search_insert_position

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	testcases := []struct {
		nums []int
		target, solution int
	}{
		{ []int{ 1, 3, 5, 6 }, 5, 2 },
		{ []int{ 1, 3, 5, 6 }, 2, 1 },
		{ []int{ 1, 3, 5, 6 }, 7, 4 },
		{ []int{}, 5, 0 },
	}

	for _, testcase := range testcases {
		result := SearchInsert( testcase.nums, testcase.target )
		if result != testcase.solution {
			t.Errorf(
				"SearchInsert: %v and %v returned %v, want %v",
				testcase.nums,
				testcase.target,
				result,
				testcase.solution,
			)
		}
	}
}
