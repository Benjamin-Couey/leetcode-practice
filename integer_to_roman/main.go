/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/integer-to-roman/description/
*/
package integer_to_roman

import (
	"strings"
)

// Map that converts integers to Roman numerals.
var value_to_string = map[int]string{
	1000: "M",
	500:  "D",
	100:  "C",
	50:   "L",
	10:   "X",
	5:    "V",
	1:    "I",
}

// Powers of ten used by IntToRoman to deconstruct num.
var powers_of_ten = []int{
	1000,
	100,
	10,
	1,
}

/*
IntToRoman returns Roman numeral representation of num.
IntToRoman assumes that 1 <= num <= 3999.
*/
func IntToRoman(num int) string {
	roman_num := ""

	for _, value := range powers_of_ten {
		quotient := num / value

		if quotient < 5 {
			if quotient == 4 {
				roman_num += strings.Repeat(value_to_string[value], 1)
				roman_num += strings.Repeat(value_to_string[value*5], 1)
			} else {
				roman_num += strings.Repeat(value_to_string[value], quotient)
			}
		} else {
			if quotient == 9 {
				roman_num += strings.Repeat(value_to_string[value], 1)
				roman_num += strings.Repeat(value_to_string[value*10], 1)
			} else {
				roman_num += strings.Repeat(value_to_string[value*5], 1)
				roman_num += strings.Repeat(value_to_string[value], quotient-5)
			}
		}
		num = num % value
	}

	return roman_num
}
