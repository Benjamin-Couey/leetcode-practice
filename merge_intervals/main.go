package merge_intervals

/* Assumes that:
1 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^4 */
func Merge(intervals [][]int) [][]int {

	merged_intervals := make( [][]int, 0 )

	for _, interval := range intervals {
		if len(merged_intervals) < 1 {
			new_merged_interval := make([]int, len(interval))
			copy( new_merged_interval, interval )
			merged_intervals = append( merged_intervals, new_merged_interval )
		} else {
			merged := false
			for index, merged_interval := range merged_intervals {
				if interval[0] <= merged_interval[1] && interval[1] >= merged_interval[1] ||
					interval[1] >= merged_interval[0] && interval[0] <= merged_interval[0] {
					if interval[0] < merged_interval[0]{
						merged_interval[0] = interval[0]
					}
					if interval[1] > merged_interval[1]{
						merged_interval[1] = interval[1]
					}
					merged_intervals[index] = merged_interval
					merged = true
				}
			}
			if !merged {
				new_merged_interval := make([]int, len(interval))
				copy( new_merged_interval, interval )
				merged_intervals = append( merged_intervals, new_merged_interval )
			}
		}
	}

	return merged_intervals
}
