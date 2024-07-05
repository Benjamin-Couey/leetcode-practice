package clone_graph

import (
	"testing"
	"leetcode/utils"
)

func TestCloneGraph(t *testing.T) {
	testcases := []struct {
		node [][]int
	}{
		{ [][]int{} },
		{ [][]int{ []int{} } },
		{ [][]int{ []int{ 2 }, []int{ 1 } } },
		{ [][]int{ []int{ 2, 3 }, []int{ 1 }, []int{ 1 } } },
		{ [][]int{ []int{ 2, 4 }, []int{ 1, 3 }, []int{ 2, 4 }, []int{ 1, 3 } } },
	}

	for _, testcase := range testcases {
		node, error := utils.AdjacencyToGraph( testcase.node )
		if error != nil {
			t.Errorf(
				"TestCloneGraph: Error when converting %v, AdjacencyToGraph raised %v",
				testcase.node,
				error,
			)
		}

		result := CloneGraph( node )

		if !utils.IsSameGraph( result, node ) {
			t.Errorf(
				"TestCloneGraph: %v returned %v, want %v",
				testcase.node,
				result,
				testcase.node,
			)
		}

		if node != nil && result == node {
			t.Errorf(
				"TestCloneGraph: returned pointer to same graph",
			)
		}
	}
}
