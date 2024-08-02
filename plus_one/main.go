/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/plus-one/description/
*/
package plus_one

/*
PlusOne returns a slice representation of an integer after adding 1 to digits.

PlusOne assumes that:
1 <= digits.length <= 100,
0 <= digits[i] <= 9,
digits is ordered from most significant to least significant in left-to-right order,
and digits does not contain any leading 0's.
*/
func PlusOne(digits []int) []int {

	sum := make([]int, 0)
	sum = append(sum, digits...)
	carry := 1
	index := len(digits) - 1

	for carry == 1 && index >= 0 {
		digit_sum := digits[index] + 1
		sum[index] = digit_sum % 10
		carry = digit_sum / 10
		index--
	}

	if carry == 1 {
		sum = append([]int{1}, sum...)
	}

	return sum
}
