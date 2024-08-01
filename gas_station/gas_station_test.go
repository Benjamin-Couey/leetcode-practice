package gas_station

import (
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	testcases := []struct {
		gas, cost []int
		solution  int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{1, 2, 3, 4, 5}, []int{3, 2, 5, 2, 3}, 3},
		{[]int{1, 2, 3, 4, 5}, []int{3, 1, 5, 3, 3}, 3},
		{[]int{1, 2, 3, 4, 5}, []int{3, 1, 2, 2, 7}, 1},
		{[]int{3, 2, 4, 4, 1}, []int{1, 3, 1, 1, 8}, 0},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
		{[]int{2, 3, 4}, []int{6, 8, 6}, -1},
	}

	for _, testcase := range testcases {
		result := CanCompleteCircuit(testcase.gas, testcase.cost)

		if result != testcase.solution {
			t.Errorf(
				"CanCompleteCircuit: %v and %v returned %v, want %v",
				testcase.gas,
				testcase.cost,
				result,
				testcase.solution,
			)
		}
	}
}
