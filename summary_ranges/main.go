package summary_ranges

import (
	"fmt"
)

/* Assumes that:
0 <= nums.length <= 20
-2^31 <= nums[i] <= 2^31 - 1
All the values of nums are unique.
nums is sorted in ascending order. */
func SummaryRanges(nums []int) []string {

	if len(nums) < 1 {
		return []string{}
	}

	range_left, range_right := nums[0], nums[0]
	ranges := []string{}
	for _, value := range( nums[1:] ) {
		if value == range_right + 1 {
			range_right = value
		} else {
			if range_left == range_right {
				ranges = append( ranges, fmt.Sprintf("%v", range_right) )
			} else {
				ranges = append( ranges, fmt.Sprintf("%v->%v", range_left, range_right) )
			}
			range_left, range_right = value, value
		}
	}

	if range_left == range_right {
		ranges = append( ranges, fmt.Sprintf("%v", range_right) )
	} else {
		ranges = append( ranges, fmt.Sprintf("%v->%v", range_left, range_right) )
	}

	return ranges
}
