package isomorphic_strings

import (
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	testcases := []struct {
		s, t     string
		solution bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}

	for _, testcase := range testcases {
		result := IsIsomorphic(testcase.s, testcase.t)

		if result != testcase.solution {
			t.Errorf(
				"IsIsomorphic: %v and %v returned %v, want %v",
				testcase.s,
				testcase.t,
				result,
				testcase.solution,
			)
		}
	}
}
