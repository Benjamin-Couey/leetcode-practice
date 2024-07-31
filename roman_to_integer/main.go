package roman_to_integer

var rune_to_value = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

/* Assumes that:
1 <= s.length <= 15.
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
t is guaranteed that s is a valid roman numeral in the range [1, 3999]. */
func RomanToInt(s string) int {

	value := 0
	num_runes := 1
	current_rune := ' '

	for _, rune := range s {
		if current_rune == ' ' {
			current_rune = rune
		} else {
			if rune == current_rune {
				num_runes++
			} else {
				// Value has gone down, indicates the special case of a rune being
				// subtracted from the following rune (for example in IX, I is
				// subtracted from X to form 9).
				if rune_to_value[current_rune] < rune_to_value[rune] {
					value -= num_runes * rune_to_value[current_rune]
				} else {
					value += num_runes * rune_to_value[current_rune]
				}
				current_rune = rune
				num_runes = 1
			}
		}
	}
	value += num_runes * rune_to_value[current_rune]

	return value
}
