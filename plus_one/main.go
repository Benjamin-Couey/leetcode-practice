package plus_one

/* Assumes that:
1 <= digits.length <= 100
0 <= digits[i] <= 9
digits does not contain any leading 0's. */
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
