package h_index

import (
	"testing"
)

type Testcase struct {
	citations []int
	solution  int
}

var testcases []Testcase = []Testcase{
	{[]int{3, 0, 6, 1, 5}, 3},
	{[]int{9, 5, 6, 7, 0, 5}, 5},
	{[]int{1, 3, 1}, 1},
	{[]int{0, 0, 0}, 0},
	{[]int{5, 5, 5}, 3},
}

func TestHIndex(t *testing.T) {
	for _, testcase := range testcases {
		result := HIndex(testcase.citations)
		if result != testcase.solution {
			t.Errorf(
				"HIndex: %v returned %v, want %v",
				testcase.citations,
				result,
				testcase.solution,
			)
		}
	}
}

func TestAltHIndex(t *testing.T) {
	for _, testcase := range testcases {
		result := AltHIndex(testcase.citations)
		if result != testcase.solution {
			t.Errorf(
				"AltHIndex: %v returned %v, want %v",
				testcase.citations,
				result,
				testcase.solution,
			)
		}
	}
}
