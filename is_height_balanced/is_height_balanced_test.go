package is_height_balanced

import (
	"leetcode/utils"
	"testing"
)

func TestIsHeightBalanced(t *testing.T) {

	root_a := &utils.TreeNode{
		1,
		&utils.TreeNode{2, nil, nil},
		&utils.TreeNode{3, nil, nil},
	}

	root_b := &utils.TreeNode{
		1,
		&utils.TreeNode{
			2,
			&utils.TreeNode{3, nil, nil},
			nil,
		},
		nil,
	}

	root_c := &utils.TreeNode{
		1,
		&utils.TreeNode{2, nil, nil},
		nil,
	}

	root_d := &utils.TreeNode{
		1,
		nil,
		&utils.TreeNode{2, nil, nil},
	}

	root_e := &utils.TreeNode{
		1,
		nil,
		nil,
	}

	testcases := []struct {
		root     *utils.TreeNode
		solution bool
		summary  string
	}{
		{root_a, true, "root_a"},
		{root_b, false, "root_b"},
		{root_c, true, "root_c"},
		{root_d, true, "root_d"},
		{root_e, true, "root_d"},
		{nil, true, "nil"},
	}

	for _, testcase := range testcases {
		same := IsHeightBalanced(testcase.root)

		if same != testcase.solution {
			t.Errorf("IsHeightBalanced: for %v returned %v, want %v", testcase.summary, same, testcase.solution)
		}
	}
}
