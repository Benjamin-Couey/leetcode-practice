package best_time_to_buy_and_sell_stock

import (
	"testing"
)

func TestMerge(t *testing.T) {
	testcases := []struct {
		prices   []int
		solution int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, testcase := range testcases {
		answer := MaxProfit(testcase.prices)

		if answer != testcase.solution {
			t.Errorf("MaxProfit: %v, want %v", answer, testcase.solution)
		}
	}
}
