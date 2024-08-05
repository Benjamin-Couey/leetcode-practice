package group_anagrams

import (
	"leetcode/utils"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testcases := []struct {
		strs     []string
		solution [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}

	for _, testcase := range testcases {
		result := GroupAnagrams(testcase.strs)

		if !utils.MatrixEqualAnyOrder(result, testcase.solution) {
			t.Errorf(
				"GroupAnagrams: %v returned %v, want %v",
				testcase.strs,
				result,
				testcase.solution,
			)
		}
	}
}
