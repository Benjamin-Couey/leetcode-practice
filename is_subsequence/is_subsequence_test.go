package is_subsequence

import (
		"testing"
)

func TestIsSubsequence(t *testing.T) {
		testcases := []struct {
				s, t string
				solution bool
		}{
				{ "abc", "ahbgdc", true },
				{ "axc", "ahbgdc", false },
		}

		for _, testcase := range testcases {
				result := IsSubsequence( testcase.s, testcase.t )

				if result != testcase.solution {
						t.Errorf(
							"IsSubsequence: %v and %v returned %v, want %v",
							testcase.s,
							testcase.t,
							result,
							testcase.solution,
						)
				}
		}
}
