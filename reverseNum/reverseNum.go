package main

import (
	"fmt"
	"math"
)

func reverseNum(x int) int {
	if x == 0 {
		return x
	}

	var rt int64
	for x > 9 || x < -9 {
		rt = rt*10 + int64(x%10)
		x = x / 10
	}
	rt = rt*10 + int64(x)
	if rt > int64(math.Pow(2, 31)-1) || rt < int64(math.Pow(-2, 31)) {
		return 0
	}
	return int(rt)
}

func main() {
	fmt.Println(reverseNum(1234567891313424))
	fmt.Println(reverseNum(-123))
}
