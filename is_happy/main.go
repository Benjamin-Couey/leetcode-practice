/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/happy-number/description/
*/
package is_happy

import (
	"math"
)

/*
IsHappy reports whether n is happy.

A happy number is a number defined by the following process:
Starting with any positive integer, replace the number by the sum of the squares
of its digits.

Repeat the process until the number equals 1 (where it will stay), or it loops
endlessly in a cycle which does not include 1.

Those numbers for which this process ends in 1 are happy.

The number will not increase forever, because for a set of n digits, the can sum
to at most n*81 while you need at least 10^(n-1) to not decrease n. Since the
latter scales faster, the sets for n = 4 or higher always decrease.

IsHappy assumes that 1 <= n <= 2^31 - 1.
*/
func IsHappy(n int) bool {
	encountered_nums := make(map[int]bool)

	for n != 1 {
		_, exists := encountered_nums[n]
		if exists {
			return false
		}
		encountered_nums[n] = true
		n = sumSquaredDigits(n)
	}

	return true
}

func sumSquaredDigits(n int) int {

	sum_of_squared_digits := 0

	for n > 0 {
		quotient := n / 10
		remainder := n % 10
		sum_of_squared_digits += int(math.Pow(float64(remainder), float64(2)))
		n = quotient
	}

	return sum_of_squared_digits
}
