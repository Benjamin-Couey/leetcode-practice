package minimum_number_of_arrows

/* Assumes that:
A balloon with xstart and xend is burst by an arrow shot at x if xstart <= x <= xend.
1 <= points.length <= 105
points[i].length == 2
-23^1 <= xstart < xend <= 2^31 - 1 */
func FindMinArrowShots(points [][]int) int {
	arrow_targets := make([][]int, 0)

	for _, point := range points {
		found_target := false

		// Check for overlap with existing targets
		for _, target := range arrow_targets {
			if point[0] <= target[1] && point[1] >= target[1] ||
				point[1] >= target[0] && point[0] <= target[0] {
				found_target = true
				// Tighten the target so it will hit the point
				if point[0] > target[0] {
					target[0] = point[0]
				}
				if point[1] < target[1] {
					target[1] = point[1]
				}
			}
		}

		if !found_target {
			append_to_intervals(&arrow_targets, point)
		}
	}

	return len(arrow_targets)
}

func append_to_intervals(intervals *[][]int, interval []int) {
	interval_to_append := make([]int, len(interval))
	copy(interval_to_append, interval)
	*intervals = append(*intervals, interval_to_append)
}
