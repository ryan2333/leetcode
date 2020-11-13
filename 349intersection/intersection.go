package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	var (
		tmpNums  []int
		nextNums []int
		num      int
		nums     []int
		ok       bool
		numsMap  map[int]bool
	)

	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	numsMap = make(map[int]bool)

	if len(nums1) > len(nums2) {
		tmpNums = nums2
		nextNums = nums1
	} else {
		tmpNums = nums1
		nextNums = nums2
	}

	for _, num = range tmpNums {
		numsMap[num] = true
	}

	for _, num = range nextNums {
		if _, ok = numsMap[num]; ok {
			nums = append(nums, num)
			delete(numsMap, num)
		}
	}

	fmt.Println("nums: ", nums)
	return nums
}

func main() {
	intersection([]int{1, 2, 2, 1}, []int{2, 2})
	intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
}
