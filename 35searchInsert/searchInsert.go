package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums)

	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid
		}
		if nums[mid] == target {
			fmt.Println(nums, target, mid)

			return mid
		}
		if left >= right {
			fmt.Println(nums, target, right)
			return right
		}
	}
	return 0
}

func main() {

	searchInsert([]int{1, 3, 5, 6}, 5)
	searchInsert([]int{1, 3, 5, 6}, 2)
	searchInsert([]int{1, 3, 5, 6}, 7)
	searchInsert([]int{1, 3, 5, 6}, 0)
}
