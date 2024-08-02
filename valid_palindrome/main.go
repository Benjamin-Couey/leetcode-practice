/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/valid-palindrome/description/
*/
package valid_palindrome

import (
	"unicode"
)

/*
IsPalindrome reports whether s is a palindrome. That is, it reads the same
forwards and backwards.
IsPalindrome assumes that:
1 <= s.length <= 2 * 10^5
and s consists only of printable ASCII characters.
*/
func IsPalindrome(s string) bool {

	var rune_list []rune = []rune(s)

	left_index := 0
	right_index := len(rune_list) - 1

	for left_index <= right_index && left_index < len(rune_list) && right_index >= 0 {
		/*
		Skip any blank space and find the next character on the left and right
		side.
		*/
		for left_index < len(rune_list) && !unicode.IsLetter(rune_list[left_index]) {
			left_index += 1
		}
		for right_index >= 0 && !unicode.IsLetter(rune_list[right_index]) {
			right_index -= 1
		}
		// If the indices haven't crossed over, compare the characters.
		if left_index <= right_index && unicode.ToLower(rune_list[left_index]) != unicode.ToLower(rune_list[right_index]) {
			return false
		}
		left_index += 1
		right_index -= 1
	}
	return true
}
