package length_of_last_word

import (
		"testing"
)

func TestLengthOfLastWord(t *testing.T) {
		testcases := []struct {
				s string
				solution int
		}{
				{ "Hello World", 5 },
				{ "   fly me   to   the moon  ", 4 },
				{ "luffy is still joyboy", 6 },
		}

		for _, testcase := range testcases {
				length := LengthOfLastWord( testcase.s )

				if length != testcase.solution {
						t.Errorf("TestLengthOfLastWord: %v, want %v", length, testcase.solution)
				}

		}
}
