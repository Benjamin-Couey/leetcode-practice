package word_pattern

import (
	"testing"
)

func TestWordPattern(t *testing.T) {
	testcases := []struct {
		pattern, s string
		solution bool
	}{
		{ "abba", "dog cat cat dog", true },
		{ "abba", "dog cat cat fish", false },
		{ "aaaa", "dog cat cat dog", false },
	}

	for _, testcase := range testcases {
		result := WordPattern( testcase.pattern, testcase.s )

		if result != testcase.solution {
			t.Errorf(
				"WordPattern: %v and %v returned %v, want %v",
				testcase.pattern,
				testcase.s,
				result,
				testcase.solution,
			)
		}
	}
}
