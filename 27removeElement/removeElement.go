package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return len(nums)
	}
	var startIndex = 0
	var lenS = len(nums)
	for i := 0; i < lenS; i++ {
		if nums[i] == val {
			startIndex = i
			nums = append(nums[:startIndex], nums[startIndex+1:]...)
			startIndex = 0
			lenS = len(nums)
			i--
		}
	}
	fmt.Println(len(nums))

	return len(nums)
}

func main() {
	removeElement([]int{3, 2, 2, 3}, 3)
	removeElement([]int{0, 1, 2, 2, 2, 3, 0, 4, 2}, 2)
}
