/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/h-index/description/
*/
package h_index

import (
	"sort"
)

/*
HIndex returns the h-index of a researcher defined by citations. The h-index
is the maximum value of h such that a researcher has published at least h papers
that have each been cited at least h times. citations is an exhaustive list of the
number of times each of a researcher's papers has been cited.
HIndex assumes that:
1 <= citations.length <= 5000,
and 0 <= citations[i] <= 1000.

HIndex runs in O(n^2log(n)) time complexity and O(n) space complexity. It could
be modified to use O(1) space complexity by mutating the original array, but this
is an undesireable side effect.
*/
func HIndex(citations []int) int {
	// Make a local copy to avoid mutating the original slice
	local_citations := make([]int, len(citations))
	copy(local_citations, citations)
	// TODO: Replace with my own sort implementation
	sort.Ints(local_citations)
	h_index := len(local_citations)
	index := 0
	for index < len(local_citations) && local_citations[index] < h_index {
		h_index--
		index++
	}
	return h_index
}

/*
AltHIndex is an alternative implementation of HIndex which runs in O(n^2)
time complexity and O(n) space complexity, trading space for time.
*/
func AltHIndex(citations []int) int {
	counts := make([]int, 1001)
	max_num := 0
	for _, num := range citations {
		counts[num]++
		if num > max_num {
			max_num = num
		}
	}

	index := max_num
	num_higher := 0
	for index > 0 {
		if (counts[index] + num_higher) >= index {
			return index
		} else {
			num_higher += counts[index]
			index--
		}
	}
	return 0
}
