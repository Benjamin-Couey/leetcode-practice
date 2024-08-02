/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
*/
package reverse_polish_notation

import (
	"fmt"
)

// The string representations of the operators which EvalRPN supports.
const ADD string = "+"
const SUBTRACT string = "-"
const MULTIPLY string = "*"
const DIVIDE string = "/"

/*
EvalRPN returns the integer that results from evaluating the arithmetic expression
tokens.
EvalRPN assumes that:
each operand may be an integer or another expression,
the division between two integers always truncates toward zero,
there will not be any division by zero,
the input represents a valid arithmetic expression in a reverse polish notation,
the answer and all the intermediate calculations can be represented in a 32-bit integer,
1 <= tokens.length <= 10^4,
and tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].
*/
func EvalRPN(tokens []string) int {

	stack := make([]string, 0)

	for _, token := range tokens {
		if len(stack) > 1 {
			last_index := len(stack) - 1
			var right_operand int
			var left_operand int
			var result int
			performed_operation := false
			_, _ = fmt.Sscanf(stack[last_index], "%d", &right_operand)
			_, _ = fmt.Sscanf(stack[last_index-1], "%d", &left_operand)
			switch token {
			case ADD:
				result = left_operand + right_operand
				performed_operation = true
			case SUBTRACT:
				result = left_operand - right_operand
				performed_operation = true
			case MULTIPLY:
				result = left_operand * right_operand
				performed_operation = true
			case DIVIDE:
				result = left_operand / right_operand
				performed_operation = true
			default:
				stack = append(stack, token)
			}
			if performed_operation {
				stack = stack[:last_index-1]
				stack = append(stack, fmt.Sprint(result))
			}
		} else {
			stack = append(stack, token)
		}
	}
	last_index := len(stack) - 1
	var result int
	_, _ = fmt.Sscanf(stack[last_index], "%d", &result)
	return result
}
