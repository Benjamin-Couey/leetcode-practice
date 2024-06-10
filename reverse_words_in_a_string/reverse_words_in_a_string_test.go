package reverse_words_in_a_string

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	testcases := []struct {
		s, solution string
	}{
		{ "the sky is blue", "blue is sky the" },
		{ "  hello world  ", "world hello" },
		{ "a good   example", "example good a" },
		{ "word", "word" },
	}

	for _, testcase := range testcases {
		result := ReverseWords( testcase.s )

		if result != testcase.solution {
			t.Errorf(
				"ReverseWords: %v returned %v, want %v",
				testcase.s,
				result,
				testcase.solution,
			)
		}
	}
}
