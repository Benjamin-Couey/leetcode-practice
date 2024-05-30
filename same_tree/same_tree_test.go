package same_tree

import (
		"testing"
		"leetcode/utils"
)

func TestIsSameTree(t *testing.T) {
		testcases := []struct {
				p_list, q_list []int
				solution bool
		}{
				{ []int{ 1, 2, 3 }, []int{ 1, 2, 3 }, true },
				{ []int{ 1, 2 }, []int{ 1, 0, 2 }, false },
				{ []int{ 1, 2, 1 }, []int{ 1, 1, 2 }, false },
		}

		for _, testcase := range testcases {
				same := IsSameTree(
					utils.LevelOrderToTree( testcase.p_list ),
					utils.LevelOrderToTree( testcase.q_list ),
				)

				if same != testcase.solution {
						t.Errorf("IsSameTree: %v, want %v", same, testcase.solution)
				}
		}
}
