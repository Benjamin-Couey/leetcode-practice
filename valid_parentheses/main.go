/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/valid-parentheses/description/
*/
package valid_parentheses

/*
IsValid reports whether the s contains valid parentheses. A string is valid if
all open parentheses are closed by the same type of parenthesis in the order they
were opened, and there are no open parentheses missing a closing parenthesis.
IsValid assumes that:
1 <= s.length <= 10^4
s consists of parentheses only '()[]{}'.
*/
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
		_, exists := left_parentheses[rune]
		if exists {
			stack = append(stack, rune)
		} else {
			if len(stack) < 1 {
				return false
			}
			last := len(stack) - 1
			if right_to_left_parentheses[rune] != stack[last] {
				return false
			} else {
				stack = stack[:last]
			}
		}
	}

	return true
}
