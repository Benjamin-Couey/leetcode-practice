package best_time_to_buy_and_sell_stock_ii

/* Assumes that:
1 <= prices.length <= 3 * 10^4
0 <= prices[i] <= 10^4 */
func MaxProfit(prices []int) int {
	sum_profit := 0
	buy_index := 0
	sell_index := 0
	/* Sell as soon as value starts decreasing and don't buy until value stops
	decreasing. */
	for index := 1; index < len(prices); index++ {
		if prices[index] < prices[sell_index] {
			sum_profit += prices[sell_index] - prices[buy_index]
			buy_index = index
			sell_index = buy_index
		} else {
			sell_index = index
		}
	}
	sum_profit += prices[sell_index] - prices[buy_index]

	return sum_profit
}
