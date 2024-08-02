/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
*/
package best_time_to_buy_and_sell_stock

/*
MaxProfit returns the maximum profit possible with prices, or 0 if no profit
can be achieved. prices[i] is the price of a given stock on the ith day. You can
choose a single day to buy one stock and choose a different day in the future to
sell that stock.
MaxProfit assumes that:
1 <= prices.length <= 10^5,
and 0 <= prices[i] <= 10^4.
*/
func MaxProfit(prices []int) int {
	buy_index := 0
	max_profit := 0

	/* If price is higher, store the profit. If price is lower, buy at that time
	instead. */
	for sell_index := 1; sell_index < len(prices); sell_index++ {
		if prices[buy_index] < prices[sell_index] {
			profit := prices[sell_index] - prices[buy_index]
			if profit > max_profit {
				max_profit = profit
			}
		} else {
			buy_index = sell_index
		}
	}

	return max_profit
}
