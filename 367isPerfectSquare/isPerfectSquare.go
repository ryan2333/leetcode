package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	if num == 2 {
		return false
	}
	left := 2
	right := num

	for left < right {
		mid := (left + right) / 2
		fmt.Printf("left: %d, right: %d, mid: %d\n", left, right, mid)
		if mid*mid < num {
			left = mid + 1
		} else if mid*mid > num {
			right = mid
		} else {
			return true
		}

	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(8))
}
