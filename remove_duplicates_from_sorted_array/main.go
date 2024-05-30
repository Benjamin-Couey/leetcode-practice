package remove_duplicates_from_sorted_array


func RemoveDuplicates(nums []int) int {
		if len( nums ) < 2 {
				return len( nums )
		}

		insert_index := 1
		last_val := nums[ 0 ]

		for index := 1; index < len( nums ); index++ {
				if nums[ index ] != last_val {
						nums[ insert_index ] = nums[ index ]
						insert_index += 1
						last_val = nums[ index ]
				}
		}

		return insert_index
}
