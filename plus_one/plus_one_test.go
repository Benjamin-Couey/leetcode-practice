package plus_one

import (
	"testing"
	"leetcode/utils"
)

func TestPlusOne(t *testing.T) {
	testcases := []struct {
		digits, solution []int
	}{
		{ []int{ 1, 2, 3 }, []int{ 1, 2, 4 } },
		{ []int{ 4, 3, 2, 1 }, []int{ 4, 3, 2, 2 } },
		{ []int{ 9 }, []int{ 1, 0 } },
		{ []int{ 0 }, []int{ 1 } },
	}

	for _, testcase := range testcases {
		result := PlusOne( testcase.digits )
		if !utils.SliceEqual( result, testcase.solution ) {
			t.Errorf(
				"PlusOne: %v returned %v, want %v",
				testcase.digits,
				result,
				testcase.solution,
			)
		}
	}
}
