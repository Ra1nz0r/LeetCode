package main

import (
	"fmt"
	"leetcode/array/best_time_to_buy_and_sell_stock"
)

func main() {
	fmt.Println(best_time_to_buy_and_sell_stock.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(best_time_to_buy_and_sell_stock.MaxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(best_time_to_buy_and_sell_stock.MaxProfit([]int{1, 5}))
}
