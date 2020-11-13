package main

import "fmt"

func isPowerOfTwo(n int) bool {
	return n&(-n) == n
}

func main() {
	fmt.Println(isPowerOfTwo(8))
}
