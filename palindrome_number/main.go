/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/palindrome-number/description/
*/
package palindrome_number

import (
	"math"
)

/*
IsPalindrome reports whether x is a palindrome.
IsPalindrome assumes that -2^31 <= x <= 2^31 - 1.
*/
func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	left_mod := int(math.Log10(float64(x)))
	right_mod := 1

	for left_mod >= right_mod {
		left_digit := x / int(math.Pow10(left_mod))
		right_digit := x % int(math.Pow10(right_mod))

		if left_digit != right_digit {
			return false
		}

		left_mod--
		right_mod++
	}

	return true
}
