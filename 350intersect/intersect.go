package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	var (
		tmpNum  []int
		NextNum []int
		nums    []int
		num     int
		numsMap map[int]int
		ok      bool
		value   int
	)

	if len(nums1) == 0 || len(nums2) == 0 {
		return nums
	}

	if len(nums1) > len(nums2) {
		tmpNum = nums2
		NextNum = nums1
	} else {
		tmpNum = nums1
		NextNum = nums2
	}
	numsMap = make(map[int]int, len(tmpNum))

	for _, num = range tmpNum {
		if _, ok = numsMap[num]; ok {
			numsMap[num]++
			continue
		}
		numsMap[num] = 1
	}

	for _, num = range NextNum {
		if value, ok = numsMap[num]; !ok || value == 0 {
			continue
		}
		nums = append(nums, num)
		numsMap[num]--
	}

	fmt.Println("nums: ", nums)
	return nums

}

func main() {
	intersect([]int{1, 2, 2, 1}, []int{2, 2})
	intersect([]int{4, 9, 5, 4}, []int{9, 4, 9, 8, 4})
}
