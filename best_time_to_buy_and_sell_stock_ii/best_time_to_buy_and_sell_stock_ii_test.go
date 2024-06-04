package best_time_to_buy_and_sell_stock_ii

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	testcases := []struct {
		prices []int
		solution int
	}{
		{ []int{ 7, 1, 5, 3, 6, 4 }, 7 },
		{ []int{ 1, 2, 3, 4, 5 }, 4 },
		{ []int{ 7, 6, 4, 3, 1 }, 0 },
	}

	for _, testcase := range testcases {
		result := MaxProfit( testcase.prices )

		if result != testcase.solution {
			t.Errorf(
				"MaxProfit: %v returned %v, want %v",
				testcase.prices,
				result,
				testcase.solution,
			)
		}
	}
}
