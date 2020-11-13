package main

import "fmt"

func maxProfit(prices []int) int {
	var (
		less    int
		current int
		preMin  int
		lp      int
		profit  int
	)
	lp = len(prices)
	if lp == 0 || lp == 1 {
		return 0
	}

	preMin = prices[0]
	profit = 0

	for i := 1; i < lp; i++ {
		current = prices[i]
		less = current - preMin
		if less > profit {
			profit = less
		}
		if current < preMin {
			preMin = current
		}
	}
	fmt.Println("profit: ", profit)
	return profit
}
func main() {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
	maxProfit([]int{7, 6, 4, 3, 1})
}
