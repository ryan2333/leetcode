package main

import "fmt"

func maxProfit(prices []int) int {
	var (
		profit int
		lp     int
	)
	profit = 0
	lp = len(prices)

	for i := 1; i < lp; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	fmt.Println("profit: ", profit)
	return profit
}
func main() {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
	maxProfit([]int{7, 6, 4, 3, 1})
	maxProfit([]int{1, 2, 3, 4, 5})
}
