/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
*/
package best_time_to_buy_and_sell_stock_ii

/*
MaxProfit returns the maximum profit possible with prices, or 0 if no profit
can be achieved. prices[i] is the price of a given stock on the ith day. Each day
you can choose to buy stock and choose a different day in the future to sell that
stock. You can only hold at most one share of the stock at any time.
MaxProfit assumes that:
1 <= prices.length <= 10^5,
and 0 <= prices[i] <= 10^4.
*/
func MaxProfit(prices []int) int {
	sum_profit := 0
	buy_index := 0
	sell_index := 0
	/*
	Sell as soon as value starts decreasing and don't buy until value stops
	decreasing.
	*/
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
