package group_anagrams

import (
	"testing"
	"leetcode/utils"
	"fmt"
)

func TestGroupAnagrams(t *testing.T) {
	testcases := []struct {
		strs []string
		solution [][]string
	}{
		{ []string{ "eat", "tea", "tan", "ate", "nat", "bat" },
			[][]string{
				{ "bat" },
				{ "nat", "tan" },
				{ "ate", "eat", "tea" },
			},
		},
		{ []string{ "" }, [][]string{	{ "" }, } },
		{ []string{ "a" }, [][]string{	{ "a" }, } },
	}

	for _, testcase := range testcases {
		result := GroupAnagrams( testcase.strs )

		fmt.Printf("%v\n", result)

		if !utils.MatrixEqualAnyOrder( result, testcase.solution ) {
			t.Errorf(
				"GroupAnagrams: %v returned %v, want %v",
				testcase.strs,
				result,
				testcase.solution,
			)
		}
	}
}
