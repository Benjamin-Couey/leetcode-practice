package summary_ranges

import (
	"testing"
	"leetcode/utils"
)

func TestSummaryRanges(t *testing.T) {
	testcases := []struct {
		nums []int
		solution []string
	}{
		{ []int{ 0, 1, 2, 4, 5, 7 }, []string{ "0->2", "4->5", "7" } },
		{ []int{ 0, 2, 3, 4, 6, 8, 9 }, []string{ "0", "2->4", "6", "8->9" } },
		{ []int{}, []string{} },
	}

	for _, testcase := range testcases {
		result := SummaryRanges( testcase.nums )

		if !utils.SliceEqual( result, testcase.solution ) {
			t.Errorf(
				"SummaryRanges: %v returned %v, want %v",
				testcase.nums,
				result,
				testcase.solution,
			)
		}
	}
}
