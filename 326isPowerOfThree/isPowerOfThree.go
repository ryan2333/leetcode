package main

import "fmt"

func isPowerOfThree(n int) bool {

	if n < 1 {
		return false
	}

	for n > 1 && n%3 == 0 {
		n /= 3
	}

	return n == 1
}

func main() {
	fmt.Println(isPowerOfThree(0))
	fmt.Println((1 << 3) - 1 - 5)
}
