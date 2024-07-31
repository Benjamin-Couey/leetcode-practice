package reverse_polish_notation

import (
	"testing"
)

func TestEvalRPN(t *testing.T) {
	testcases := []struct {
		tokens   []string
		solution int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
		{[]string{"-1"}, -1},
	}

	for _, testcase := range testcases {
		result := EvalRPN(testcase.tokens)

		if result != testcase.solution {
			t.Errorf(
				"EvalRPN: %v returned %v, want %v",
				testcase.tokens,
				result,
				testcase.solution,
			)
		}
	}
}
