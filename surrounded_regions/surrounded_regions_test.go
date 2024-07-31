package surrounded_regions

import (
	"leetcode/utils"
	"testing"
)

func TestSolve(t *testing.T) {
	testcases := []struct {
		matrix   [][]byte
		solution [][]byte
	}{
		{[][]byte{
			{byte('X'), byte('X'), byte('X'), byte('X')},
			{byte('X'), byte('O'), byte('O'), byte('X')},
			{byte('X'), byte('X'), byte('O'), byte('X')},
			{byte('X'), byte('O'), byte('X'), byte('X')},
		}, [][]byte{
			{byte('X'), byte('X'), byte('X'), byte('X')},
			{byte('X'), byte('X'), byte('X'), byte('X')},
			{byte('X'), byte('X'), byte('X'), byte('X')},
			{byte('X'), byte('O'), byte('X'), byte('X')},
		},
		},
		{[][]byte{
			{byte('X'), byte('X'), byte('X'), byte('X')},
			{byte('X'), byte('O'), byte('O'), byte('O')},
			{byte('X'), byte('O'), byte('X'), byte('X')},
			{byte('X'), byte('X'), byte('X'), byte('X')},
		}, [][]byte{
			{byte('X'), byte('X'), byte('X'), byte('X')},
			{byte('X'), byte('O'), byte('O'), byte('O')},
			{byte('X'), byte('O'), byte('X'), byte('X')},
			{byte('X'), byte('X'), byte('X'), byte('X')},
		},
		},
		{[][]byte{
			{byte('O'), byte('X'), byte('O'), byte('X')},
			{byte('X'), byte('O'), byte('X'), byte('O')},
			{byte('O'), byte('X'), byte('O'), byte('X')},
			{byte('X'), byte('O'), byte('X'), byte('O')},
		}, [][]byte{
			{byte('O'), byte('X'), byte('O'), byte('X')},
			{byte('X'), byte('X'), byte('X'), byte('O')},
			{byte('O'), byte('X'), byte('X'), byte('X')},
			{byte('X'), byte('O'), byte('X'), byte('O')},
		},
		},
		{[][]byte{
			{byte('X'), byte('O'), byte('X'), byte('X')},
		}, [][]byte{
			{byte('X'), byte('O'), byte('X'), byte('X')},
		},
		},
		{[][]byte{{byte('O')}}, [][]byte{{byte('O')}}},
	}

	for _, testcase := range testcases {
		Solve(testcase.matrix)

		if !utils.MatrixEqual(testcase.matrix, testcase.solution) {
			t.Errorf(
				"Solve: returned %v, want %v",
				testcase.matrix,
				testcase.solution,
			)
		}
	}
}
