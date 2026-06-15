package best_time_to_buy_and_sell_stock

func MaxProfit(prices []int) int {
	seen := make(map[int]int, len(prices))
	pricesMinMax := map[string]int{
		"max": 0,
		"min": 0,
	}

	for k, v := range prices {
		if pricesMinMax["max"] < v {
			pricesMinMax["max"] = v
		}

		if pricesMinMax["min"] > v {
			pricesMinMax["min"] = v
		}

		seen[v] = k
	}

	if seen[pricesMinMax["max"]] > seen[pricesMinMax["min"]] {
		return pricesMinMax["max"] - pricesMinMax["min"]
	}

	return 0
}
