package main

import "fmt"

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	numMap := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			return true
		}
		numMap[num] = struct{}{}
	}
	numMap = nil
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 3, 4}))
}
