package valid_parentheses

/* Assumes that:
1 <= s.length <= 104
s consists of parentheses only '()[]{}'. */
func IsValid(s string) bool {

	left_parentheses := map[rune]bool{
		'(': true,
		'[': true,
		'{': true,
	}

	right_to_left_parentheses := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}

	for _, rune := range s {
		_, exists := left_parentheses[ rune ]
		if exists {
			stack = append( stack, rune )
		} else {
			if len(stack) < 1 {
				return false
			}
			last := len(stack) - 1
			if right_to_left_parentheses[ rune ] != stack[ last ] {
				return false
			} else {
				stack = stack[:last]
			}
		}
	}

	return true
}
