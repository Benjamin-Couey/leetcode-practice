package length_of_last_word

/* Assumes that:
1 <= s.length <= 104
s consists of only English letters and spaces ' '.
There will be at least one word in s. */
func LengthOfLastWord(s string) int {

	length_of_current_word := 0
	space := true

	for _, rune := range s {
		if rune == ' ' {
			space = true
		} else {
			if space {
				space = false
				length_of_current_word = 1
			} else {
				length_of_current_word += 1
			}
		}
	}

	return length_of_current_word
}
