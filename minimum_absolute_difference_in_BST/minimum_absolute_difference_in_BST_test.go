package minimum_absolute_difference_in_BST

import (
	"testing"
	"leetcode/utils"
)

func TestGetMinimumDifference(t *testing.T) {
	testcases := []struct {
		tree_list []utils.NilInt
		solution int
	}{
		{ []utils.NilInt{ utils.NewInt(4), utils.NewInt(2), utils.NewInt(6), utils.NewInt(1), utils.NewInt(3) }, 1 },
		{ []utils.NilInt{ utils.NewInt(1), utils.NewInt(0), utils.NewInt(48), utils.NewNil(), utils.NewNil(), utils.NewInt(12), utils.NewInt(49) }, 1 },
		{ []utils.NilInt{ utils.NewInt(6), utils.NewInt(0), utils.NewInt(48), utils.NewNil(), utils.NewNil(), utils.NewInt(12), utils.NewInt(49) }, 1 },
	}

	for _, testcase := range testcases {
		result := GetMinimumDifference( utils.NillableLevelOrderToTree( testcase.tree_list ) )

		if result != testcase.solution {
			t.Errorf(
				"GetMinimumDifference: %v returned %v, want %v",
				testcase.tree_list,
				result,
				testcase.solution,
			)
		}
	}
}
