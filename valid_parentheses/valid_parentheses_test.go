package valid_parentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testcases := []struct {
		s string
		solution bool
	}{
		{ "()", true },
		{ "()[]{}", true },
		{ "([{}])", true },
		{ "())", false },
		{ "(]", false },
	}

	for _, testcase := range testcases {
		is_valid := IsValid( testcase.s )

		if is_valid != testcase.solution {
			t.Errorf(
				"IsValid: for string %v returned %v, want %v",
				testcase.s,
				is_valid,
				testcase.solution,
			)
		}
	}
}
