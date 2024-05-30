package valid_anagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	testcases := []struct {
		s, t string
		solution bool
	}{
		{ "anagram", "nagaram", true },
		{ "rat", "car", false },
	}

	for _, testcase := range testcases {
		result := IsAnagram( testcase.s, testcase.t )

		if result != testcase.solution {
			t.Errorf(
				"IsAnagram: %v and %v returned %v, want %v",
				testcase.s,
				testcase.t,
				result,
				testcase.solution,
			)
		}
	}
}
