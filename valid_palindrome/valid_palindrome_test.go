package valid_palindrome

import (
		"testing"
)

func TestIsPalindrome(t *testing.T) {
		testcases := []struct {
				s string
				solution bool
		}{
				{ "A man, a plan, a canal: Panama", true },
				{ "race a car", false },
				{ " ", true },
		}

		for _, testcase := range testcases {
				value := IsPalindrome( testcase.s )

				if value != testcase.solution {
						t.Errorf("TestIsPalindrome: %v returned %v, want %v", testcase.s, value, testcase.solution)
				}

		}
}
