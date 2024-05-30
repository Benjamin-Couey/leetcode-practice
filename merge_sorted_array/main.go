package merge_sorted_array

func Merge(nums1 []int, m int, nums2 []int, n int) {
		if n < 1 {
				return
		}
		if m < 1 {
				for i := range(n) {
						nums1[i] = nums2[i]
				}
				return
		}

		nums1_index := m - 1
		nums2_index := n - 1
		for insert_index := m + n - 1; insert_index >= 0; insert_index-- {
				if nums1_index < 0 {
						nums1[ insert_index ] = nums2[ nums2_index ]
						nums2_index--
				} else if nums2_index < 0 {
						nums1[ insert_index ] = nums1[ nums1_index ]
						nums1_index--
				} else {
						if nums1[ nums1_index ] > nums2[ nums2_index ] {
								nums1[ insert_index ] = nums1[ nums1_index ]
								nums1_index--
						} else {
								nums1[ insert_index ] = nums2[ nums2_index ]
								nums2_index--
						}
				}
		}
}
