package roman_to_integer

import (
		"testing"
)

func TestRomanToInteger(t *testing.T) {
		testcases := []struct {
				s string
				solution int
		}{
				{ "III", 3 },
				{ "LVIII", 58 },
				{ "MCMXCIV", 1994 },
		}

		for _, testcase := range testcases {
				value := RomanToInt( testcase.s )

				if value != testcase.solution {
						t.Errorf("TestRomanToInteger: %v, want %v", value, testcase.solution)
				}

		}
}
