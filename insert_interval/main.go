package insert_interval

/* Assumes that:
0 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^5
intervals is sorted by starti in ascending order.
intervals contains no overlapping intervals
newInterval.length == 2
0 <= start <= end <= 10^5 */
func Insert(intervals [][]int, newInterval []int) [][]int {

	new_intervals := make([][]int, 0)

	inserted := false
	for _, interval := range intervals {
		// If there is overlap, merge new interval
		if newInterval[0] <= interval[1] && newInterval[1] >= interval[1] ||
			newInterval[1] >= interval[0] && newInterval[0] <= interval[0] {
			if interval[0] < newInterval[0] {
				newInterval[0] = interval[0]
			}
			if interval[1] > newInterval[1] {
				newInterval[1] = interval[1]
			}
		} else {
			// Append new interval in order
			if newInterval[0] < interval[0] && newInterval[1] < interval[0] {
				inserted = true
				append_to_intervals(&new_intervals, newInterval)
			}

			append_to_intervals(&new_intervals, interval)
		}
	}

	if !inserted {
		append_to_intervals(&new_intervals, newInterval)
	}

	return new_intervals
}

func append_to_intervals(intervals *[][]int, interval []int) {
	interval_to_append := make([]int, len(interval))
	copy(interval_to_append, interval)
	*intervals = append(*intervals, interval_to_append)
}
