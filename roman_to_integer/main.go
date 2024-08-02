/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/roman-to-integer/description/
*/
package roman_to_integer

// Map that converts Roman numerals to integers.
var rune_to_value = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

/*
RomanToInt returns integer represented by Roman numeral string s.
RomanToInt assumes that:
1 <= s.length <= 15,
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M'),
and t is guaranteed that s is a valid roman numeral in the range [1, 3999].
*/
func RomanToInt(s string) int {

	value := 0
	num_runes := 1
	last_rune := ' '

	for _, rune := range s {
		if last_rune == ' ' {
			last_rune = rune
		} else {
			if rune == last_rune {
				num_runes++
			} else {
				/*
				If the value of the current rune is higher than the last rune, this
				is the special case of a rune being subtracted from the following rune
				( for example in IX, I is subtracted from X to form 9 ).
				*/
				if rune_to_value[last_rune] < rune_to_value[rune] {
					value -= num_runes * rune_to_value[last_rune]
				} else {
					value += num_runes * rune_to_value[last_rune]
				}
				last_rune = rune
				num_runes = 1
			}
		}
	}
	value += num_runes * rune_to_value[last_rune]

	return value
}
