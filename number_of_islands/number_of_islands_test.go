package number_of_islands

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	testcases := []struct {
		grid [][]byte
		solution int
	}{
		{ [][]byte{
				{ byte('1'), byte('1'), byte('1'), byte('1'), byte('0'), },
				{ byte('1'), byte('1'), byte('0'), byte('1'), byte('0'), },
				{ byte('1'), byte('1'), byte('0'), byte('0'), byte('0'), },
				{ byte('0'), byte('0'), byte('0'), byte('0'), byte('0'), },
			}, 1,
		},
		{ [][]byte{
				{ byte('1'), byte('1'), byte('0'), byte('0'), byte('0'), },
				{ byte('1'), byte('1'), byte('0'), byte('0'), byte('0'), },
				{ byte('0'), byte('0'), byte('1'), byte('0'), byte('0'), },
				{ byte('0'), byte('0'), byte('0'), byte('1'), byte('1'), },
			}, 3,
		},
		{ [][]byte{
				{ byte('1'), byte('0'), byte('1'), },
				{ byte('0'), byte('1'), byte('0'), },
				{ byte('1'), byte('0'), byte('1'), },
			}, 5,
		},
		{ [][]byte{ { byte('0'), byte('0'), byte('0'), }, }, 0, },
	}

	for _, testcase := range testcases {
		result := NumIslands( testcase.grid )

		if result != testcase.solution {
			t.Errorf(
				"NumIslands: %v returned %v, want %v",
				testcase.grid,
				result,
				testcase.solution,
			)
		}
	}
}
