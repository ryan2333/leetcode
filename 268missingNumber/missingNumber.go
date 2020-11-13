package main

import (
	"fmt"
	"sort"
)

func missingNumber1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i && i != len(nums) {
			return i
		}
	}
	return len(nums)
}

func missingNumber(nums []int) int {
	mp := make(map[int]bool, len(nums))

	for _, v := range nums {
		mp[v] = true
	}
	for i := 0; i <= len(nums); i++ {
		if !mp[i] {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{}))
	fmt.Println(missingNumber([]int{0}))
}
