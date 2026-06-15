package best_time_to_buy_and_sell_stock

func MaxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices[1:] {
		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}

		if price < minPrice {
			minPrice = price
		}
	}

	return maxProfit
}
