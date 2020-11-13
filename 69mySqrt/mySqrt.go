package main

import "fmt"

func mysqrt1(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	var y = 0
	var flag = 1

	if x < 0 {
		flag = -1
		x *= flag
	}

	for i := 1; i <= x/2; i++ {
		if i*i > x {
			break
		}
		y = i
	}
	return y * flag
}

func mysqrt2(x int) int {

	if x < 2 {
		return x
	}

	var left = 2
	var right = x / 2

	for left <= right {
		fmt.Printf("left: %d, right: %d\n", left, right)
		mid := (left + right) / 2
		if mid*mid < x {
			left = mid + 1
		}
		if mid*mid > x {
			right = mid - 1
		}
		if mid*mid == x {
			return mid
		}
	}

	return right
}

func main() {
	fmt.Println(mysqrt2(20))
}
