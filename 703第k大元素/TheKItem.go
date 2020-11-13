package main

import (
	"fmt"
	"sort"
)

type KthLargest struct {
	Min int
	All []int
}

func Constructor(k int, nums []int) KthLargest {
	// 当k小于nums长度时，则先排序nums， 然后保留后k个元素，最后调整堆
	if k < len(nums) {
		sort.Ints(nums)
		nums = nums[len(nums)-k:]
	}

	return KthLargest{
		Min: k,
		All: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	if len(this.All) < 2 {
		this.Min = val
		this.All = append(this.All, val)
	} else if len(this.All) < 3 {

	}
	return this.Min
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	fmt.Println([]int{1, 5, 7, 9, 2, 3, 10, 7}, 9)
	a := []int{1, 5, 7, 9, 2, 3, 10, 7}
	sort.Ints(a)
	fmt.Println(a)
}
