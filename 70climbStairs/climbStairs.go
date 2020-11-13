package main

import "fmt"

func climbStairs(n int) int {

	if n <= 2 {
		return n
	}
	var a int = 1
	var b int = 2
	var ret int

	for i := 3; i <= n; i++ {
		ret = a + b
		a = b
		b = ret
	}
	fmt.Println(ret)
	return ret
}

func main() {
	climbStairs(5)
}
