package product_of_array_except_self

import (
	"leetcode/utils"
	"testing"
)

type Testcase struct {
	nums, solution []int
}

var testcases []Testcase = []Testcase{
	{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	{[]int{0, 0}, []int{0, 0}},
}

func TestProductExceptSelf(t *testing.T) {
	for _, testcase := range testcases {
		result := ProductExceptSelf(testcase.nums)

		if !utils.SliceEqual(result, testcase.solution) {
			t.Errorf(
				"ProductExceptSelf: %v returned %v, want %v",
				testcase.nums,
				result,
				testcase.solution,
			)
		}
	}
}

func TestAltProductExceptSelf(t *testing.T) {
	for _, testcase := range testcases {
		/* Since AltProductExceptSelf mutates the array of the slice passed to it,
		use a local copy of the testcase to avoid contaminating other tests. */
		local_nums := make([]int, len(testcase.nums))
		copy(local_nums, testcase.nums)
		result := AltProductExceptSelf(local_nums)

		if !utils.SliceEqual(result, testcase.solution) {
			t.Errorf(
				"ProductExceptSelf: %v returned %v, want %v",
				testcase.nums,
				result,
				testcase.solution,
			)
		}
	}
}
