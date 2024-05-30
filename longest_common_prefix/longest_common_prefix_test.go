package longest_common_prefix


import (
		"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
		testcases := []struct {
				strs []string
				solution string
		}{
				{ []string{ "flower","flow","flight" }, "fl" },
				{ []string{ "dog","racecar","car" }, "" },
		}

		for _, testcase := range testcases {
				prefix := LongestCommonPrefix( testcase.strs )

				if prefix != testcase.solution {
						t.Errorf("LongestCommonPrefix: %v, want %v", prefix, testcase.solution)
				}
		}
}
