package first_occurrence_in_string

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	testcases := []struct {
		haystack, needle string
		solution         int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
	}

	for _, testcase := range testcases {
		result := StrStr(testcase.haystack, testcase.needle)

		if result != testcase.solution {
			t.Errorf(
				"StrStr: %v and %v returned %v, want %v",
				testcase.haystack,
				testcase.needle,
				result,
				testcase.solution,
			)
		}
	}
}
