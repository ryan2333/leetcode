package main

import (
	"fmt"
)

func countPrimes(n int) int {
	var count int
	var nmap = make([]bool, n)
	for i := 2; i < n; i++ {
		if nmap[i] {
			continue
		}
		count++
		for j := 2 * i; j < n; j += i {
			nmap[j] = true
		}
	}
	return count
}

func main() {
	fmt.Println("count: ", countPrimes(499979))
	fmt.Println("count: ", countPrimes(10))
	fmt.Println("count: ", countPrimes(2))
	fmt.Println("count: ", countPrimes(3))
}
