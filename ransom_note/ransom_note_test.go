package ransom_note

import (
		"testing"
)

func TestCanConstruct(t *testing.T) {
		testcases := []struct {
				ransomNote, magazine string
				solution bool
		}{
				{ "a", "b", false },
				{ "aa", "ab", false },
				{ "aa", "aab", true },
		}

		for _, testcase := range testcases {
				result := CanConstruct( testcase.ransomNote, testcase.magazine )

				if result != testcase.solution {
						t.Errorf(
							"CanConstruct: %v and %v returned %v, want %v",
							testcase.ransomNote,
							testcase.magazine,
							result,
							testcase.solution,
						)
				}
		}
}
