package h_index

import (
	"sort"
)

/* O(n^2log(n)) time complexity and O(n) space complexity ( O(1) if you don't
care about mutating the original array ).
Assumes that:
1 <= citations.length <= 5000
0 <= citations[i] <= 1000 */
func HIndex(citations []int) int {
	// Make a local copy to avoid mutating the original slice
	local_citations := make([]int, len(citations))
	copy(local_citations, citations)
	// TODO: Replace with my own sort implementation
	sort.Ints( local_citations )
	h_index := len(local_citations)
	index := 0
	for index < len(local_citations) && local_citations[ index ] < h_index {
		h_index--
		index++
	}
	return h_index
}

/* O(n^2) time complexity and O(n) space complexity.
Trades space for speed.*/
func AltHIndex(citations []int) int {
	counts := make([]int, 1001)
	max_num := 0
	for _, num := range (citations) {
		counts[ num ]++
		if num > max_num {
			max_num = num
		}
	}

	index := max_num
	num_higher := 0
	for index > 0 {
		if (counts[ index ] + num_higher) >= index {
			return index
		} else {
			num_higher += counts[index]
			index--
		}
	}
	return 0
}
