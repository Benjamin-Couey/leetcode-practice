package integer_to_roman

import (
		"strings"
)

var value_to_string = map[int]string {
		1000: "M",
		500: "D",
		100: "C",
		50: "L",
		10: "X",
		5: "V",
		1: "I",
}

var values_to_check = []int {
	1000,
	100,
	10,
	1,
}

// Assumes that 1 <= num <= 3999.
func IntToRoman(num int) string {
		roman_num := ""

		for _, value := range values_to_check {
				quotient := num / value

				if quotient < 5 {
						if quotient == 4 {
								roman_num += strings.Repeat( value_to_string[ value ], 1)
								roman_num += strings.Repeat( value_to_string[ value * 5 ], 1 )
						} else {
								roman_num += strings.Repeat( value_to_string[ value ], quotient )
						}
				} else {
						if quotient == 9 {
								roman_num += strings.Repeat( value_to_string[ value ], 1)
								roman_num += strings.Repeat( value_to_string[ value * 10 ], 1 )
						} else {
								roman_num += strings.Repeat( value_to_string[ value * 5 ], 1 )
								roman_num += strings.Repeat( value_to_string[ value ], quotient - 5 )
						}
				}
				num = num % value
		}

		return roman_num
}
