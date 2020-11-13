package main

import "fmt"

func searchRange(nums []int, target int) []int {

	var start, end, mid int

	mid = binarySearch(nums, target)
	if mid == -1 {
		return []int{-1, -1}
	}
	start, end = mid, mid
	for start > 0 && nums[start-1] == target {
		start--
	}

	for end < len(nums)-1 && nums[end+1] == target {
		end++
	}

	return []int{start, end}
}

func binarySearch(nums []int, target int) int {
	var start, end, mid int
	start = 0
	end = len(nums) - 1
	for start <= end {
		mid = (end-start)/2 + start

		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func searchRange1(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := left_bond1(nums, target)
	// fmt.Println("left: ", left)
	right := right_bond1(nums, target)
	//fmt.Println("right: ", right)
	return []int{left, right}
}

func left_bond1(nums []int, target int) int {

	var start, end, mid int
	start, end = 0, len(nums)-1

	for start <= end {
		mid = start + (end-start)/2
		if nums[mid] == target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		}
	}
	if start >= len(nums) || nums[start] != target {
		return -1
	}
	return start
}

func right_bond1(nums []int, target int) int {
	var start, end, mid int
	start, end = 0, len(nums)-1
	for start <= end {
		mid = start + (end-start)/2
		if nums[mid] == target {
			start = mid + 1
		} else if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		}
		//fmt.Println("right_bond: ", start, end)
	}
	if end < 0 || nums[end] != target {
		return -1
	}
	return start - 1
}

func searchRange2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := left_bond2(nums, target)
	// fmt.Println("left: ", left)
	right := right_bond2(nums, target)
	//fmt.Println("right: ", right)
	if left > right {
		return []int{-1, -1}
	}
	return []int{left, right}
}

func left_bond2(nums []int, target int) int {

	var start, end, mid int
	start, end = 0, len(nums)

	for start < end {
		mid = start + (end-start)/2
		if nums[mid] == target {
			end = mid
		} else if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid
		}
	}

	return start
}

func right_bond2(nums []int, target int) int {
	var start, end, mid int
	start, end = 0, len(nums)
	for start < end {
		mid = start + (end-start)/2
		if nums[mid] == target {
			start = mid + 1
		} else if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid
		}
		//fmt.Println("right_bond: ", start, end)
	}

	return start - 1
}

func main() {
	fmt.Println(searchRange2([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange2([]int{5, 7, 7, 8, 8, 10}, 12))
	fmt.Println(searchRange2([]int{5}, 5))
	fmt.Println(searchRange2([]int{2, 2}, 2))
	fmt.Println(searchRange2([]int{1, 1, 2}, 1))
}

func test() {
	fmt.Println("Hello")
}
