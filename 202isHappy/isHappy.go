package main

import (
	"fmt"
)

func getNum(n int) []int {
	var nums []int
	for n > 9 {
		nums = append([]int{n % 10}, nums...)
		n /= 10
	}
	nums = append([]int{n}, nums...)
	return nums
}

func isHappy(n int) bool {
	var sum int
	var nmap map[int]int
	var ok bool
	var count int
	nmap = make(map[int]int)
	nmap[n] = 0
	for {
		count++
		nums := getNum(n)
		for _, num := range nums {
			sum += num * num
		}
		if sum == 1 {
			return true
		}
		if _, ok = nmap[sum]; ok {
			return false
		}
		n = sum
		sum = 0
		nmap[n] = 0
	}
	return false
}

func getSum(n int) int {
	var sum int

	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	return sum
}

func isHappy2(n int) bool {
	m := map[int]bool{}
	for {
		n = getSum(n)
		if n == 1 {
			return true
		}
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = false
	}
}

func main() {
	//fmt.Println(isHappy(19))
	fmt.Println(isHappy2(2))
	fmt.Println(isHappy2(19))
}
