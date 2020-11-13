package main

import "fmt"

func fib(N int) int {
	if N < 2 {
		return N
	}
	var first, second int
	first = 0
	second = 1
	for i := 0; i < N-1; i++ {
		first, second = second, first+second
	}
	return second
}

func main() {

	fmt.Println(fib(10))
}
