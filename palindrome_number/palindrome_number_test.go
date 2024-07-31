package palindrome_number

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testcases := []struct {
		x        int
		solution bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}

	for _, testcase := range testcases {
		result := IsPalindrome(testcase.x)
		if result != testcase.solution {
			t.Errorf(
				"IsPalindrome: %v returned %v, want %v",
				testcase.x,
				result,
				testcase.solution,
			)
		}
	}
}
