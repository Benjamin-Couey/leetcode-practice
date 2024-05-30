package best_time_to_buy_and_sell_stock

func MaxProfit(prices []int) int {
		buy_index := 0
		max_profit := 0

		// if index is higher, store the profit
		// if index is lower, switch to that index
		for sell_index := 1; sell_index < len( prices ); sell_index++ {
				if prices[ buy_index ] < prices[ sell_index ] {
						profit := prices[ sell_index ] - prices[ buy_index ]
						if profit > max_profit {
								max_profit = profit
						}
				} else {
						buy_index = sell_index
				}
		}

		return max_profit
}
