package rotate_array

import (
	"testing"
	"leetcode/utils"
)

type Testcase struct {
	nums, solution []int
	k int
}

var testcases []Testcase = []Testcase {
	{ []int{ 1, 2, 3, 4, 5, 6, 7 }, []int{ 7, 1, 2, 3, 4, 5, 6 }, 1 },
	{ []int{ 1, 2, 3, 4, 5, 6, 7 }, []int{ 6, 7, 1, 2, 3, 4, 5 }, 2 },
	{ []int{ 1, 2, 3, 4, 5, 6, 7 }, []int{ 5, 6, 7, 1, 2, 3, 4 }, 3 },
	{ []int{ -1, -100, 3, 99 }, []int{ 3, 99, -1, -100 }, 2 },
}

func TestRotate(t *testing.T) {
	for _, testcase := range testcases {
		/* Since Rotate mutates the array of the slice passed to it, use a local
		copy of the testcase to avoid contaminating other tests. */
		local_nums := make( []int, len(testcase.nums) )
		copy( local_nums, testcase.nums )

		Rotate( local_nums, testcase.k )

		if !utils.SliceEqual( local_nums, testcase.solution ) {
			t.Errorf("Rotate: returned %v, want %v", local_nums, testcase.solution)
		}
	}
}

func TestAltRotate(t *testing.T) {
	for _, testcase := range testcases {
		local_nums := make( []int, len(testcase.nums) )
		copy( local_nums, testcase.nums )

		AltRotate( local_nums, testcase.k )

		if !utils.SliceEqual( local_nums, testcase.solution ) {
			t.Errorf("AltRotate: returned %v, want %v", local_nums, testcase.solution)
		}
	}
}
