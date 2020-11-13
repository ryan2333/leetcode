package main

import (
	"fmt"
	"time"
)

// 暴力破解法
func twoSum1(nums []int, target int) []int {
	start := time.Now()
	if len(nums) < 2 {
		return []int{}
	}
	for index, i := range nums {
		for j := index + 1; j < len(nums); j++ {
			if nums[j] == target-i {
				fmt.Println("time cost: ", time.Now().Sub(start))
				return []int{index, j}
			}
		}
	}
	fmt.Println("time cost: ", time.Now().Sub(start))
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	// map
	//start := time.Now()
	if len(nums) < 2 {
		return nil
	}
	m := make(map[int]int, 2)

	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			//fmt.Println("time cost: ", time.Now().Sub(start))
			return []int{k, i}
		}
		m[v] = i
	}
	//fmt.Println("time cost: ", time.Now().Sub(start))
	//fmt.Println("map: %v", m)
	return nil
}

func twoSum(nums []int, target int) []int {
	var (
		nmap map[int]int
		i    int
		n    int
		ok   bool
	)
	nmap = make(map[int]int)

	for i, n = range nums {
		if _, ok = nmap[target-n]; ok {
			return []int{nmap[target-n], i}
		}
		nmap[n] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
