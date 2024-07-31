package zigzag_conversion

/* Assumes that:
1 <= s.length <= 1000
s consists of English letters (lower-case and upper-case), ',' and '.'.
1 <= numRows <= 1000 */
func Convert(s string, numRows int) string {

	rune_s := []rune(s)
	return_runes := make([]rune, 0)

	step := numRows
	if numRows > 2 {
		step += numRows - 2
	}

	for start_index := 0; start_index < numRows; start_index++ {
		left_step := step - 2*start_index
		right_step := step - left_step

		working_index := start_index
		use_left := true
		for working_index < len(rune_s) {
			return_runes = append(return_runes, rune_s[working_index])
			if use_left {
				if left_step > 0 {
					working_index += left_step
				} else {
					working_index += right_step
				}
				use_left = false
			} else {
				if right_step > 0 {
					working_index += right_step
				} else {
					working_index += left_step
				}
				use_left = true
			}
		}
	}

	return string(return_runes)
}
