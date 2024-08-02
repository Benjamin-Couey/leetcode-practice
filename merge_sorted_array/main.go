/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/merge-sorted-array/description/
*/
package merge_sorted_array

/*
Merge merges nums2 into nums1 while maintaining non-decreasing order.
Merge assumes that:
nums1 and nums2 are sorted in non-decreasing order,
nums1.length == m + n,
nums2.length == n,
0 <= m, n <= 200,
1 <= m + n <= 200,
and -10^9 <= nums1[i], nums2[j] <= 10^9.
*/
func Merge(nums1 []int, m int, nums2 []int, n int) {
	if n < 1 {
		return
	}
	if m < 1 {
		for i := range n {
			nums1[i] = nums2[i]
		}
		return
	}

	nums1_index := m - 1
	nums2_index := n - 1
	for insert_index := m + n - 1; insert_index >= 0; insert_index-- {
		if nums1_index < 0 {
			nums1[insert_index] = nums2[nums2_index]
			nums2_index--
		} else if nums2_index < 0 {
			nums1[insert_index] = nums1[nums1_index]
			nums1_index--
		} else {
			if nums1[nums1_index] > nums2[nums2_index] {
				nums1[insert_index] = nums1[nums1_index]
				nums1_index--
			} else {
				nums1[insert_index] = nums2[nums2_index]
				nums2_index--
			}
		}
	}
}
