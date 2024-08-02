/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/length-of-last-word/description/
*/
package length_of_last_word

/*
LengthOfLastWord returns the length of the last word in s.
LengthOfLastWord assumes that:
1 <= s.length <= 10^4,
s consists of only English letters and spaces ' ',
a word is a maximal substring containing only non-space characters,
and there will be at least one word in s.
*/
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
