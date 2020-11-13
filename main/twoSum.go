package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	var (
		index, num int
	)
	indexs := make([]int, 2)

	for index, num = range nums {
		subNum := target - num
		for subindex, subnum := range nums[index+1:] {

			if subnum == subNum {
				fmt.Println(index, subindex)
				indexs[0], indexs[1] = index, index+1+subindex

			}
		}
	}

	return indexs
}
