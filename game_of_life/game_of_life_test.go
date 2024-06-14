package game_of_life

import (
	"testing"
	"leetcode/utils"
)

func TestGameOfLife(t *testing.T) {
	testcases := []struct {
		board [][]int
		solution [][]int
	}{
		{ [][]int{
				{ 0, 1, 0, },
				{ 0, 0, 1, },
				{ 1, 1, 1, },
				{ 0, 0, 0, },
			}, [][]int{
				{ 0, 0, 0, },
				{ 1, 0, 1, },
				{ 0, 1, 1, },
				{ 0, 1, 0, },
			},
		},
		{ [][]int{
				{ 1, 1, },
				{ 1, 0, },
			}, [][]int{
				{ 1, 1, },
				{ 1, 1, },
			},
		},
		{ [][]int{
				{ 1, },
			}, [][]int{
				{ 0, },
			},
		},
	}

	for _, testcase := range testcases {
		GameOfLife( testcase.board )

		if !utils.MatrixEqual( testcase.board, testcase.solution ) {
			t.Errorf(
				"GameOfLife: returned %v, want %v",
				testcase.board,
				testcase.solution,
			)
		}
	}
}
