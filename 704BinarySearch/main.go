package main

import "fmt"

func BinarySearch(nums []int, target int) int {
	var start, end, mid int
	start = 0
	end = len(nums) - 1
	for start <= end {
		mid = (end-start)/2 + start
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func BinarySearch2(nums []int, target int) int {
	var start, end, mid int
	start = 0
	end = len(nums)
	for start < end {
		mid = (end-start)/2 + start
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	fmt.Println(BinarySearch2([]int{5}, 1))
	fmt.Println(BinarySearch2([]int{-1, 3, 5, 7, 9}, 5))
	fmt.Println(BinarySearch2([]int{2, 3, 4, 6, 8}, 1))
	fmt.Println(BinarySearch2([]int{2, 3}, 2))
	fmt.Println(BinarySearch2([]int{2}, 2))
}

func test() {
	fmt.Println("Hello")
}
