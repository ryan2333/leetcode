package main

import "fmt"

func sortArray(nums []int) []int {

	nLen := len(nums)
	if nLen == 0 {
		return nums
	}
	l := 0
	//mid := nLen - 1
	pivot := nums[0]
	r := nLen - 1

	for l < r {
		for nums[r] >= pivot && l < r {
			r--
		}

		//fmt.Printf("befor swap right: %v, r: %d[%d], l: [%d], pivot: %d, nums: %v\n", nums, r, nums[r], l, pivot, nums)

		nums[l] = nums[r]
		//fmt.Printf("after swap right: %v, pivot: %d\n ", nums, pivot)

		for nums[l] <= pivot && l < r {
			l++
		}

		//fmt.Printf("befor swap left: %v, l: %d[%d], r: %d, pivot: %d, nums: %v\n", nums, l, nums[l], r, pivot, nums)

		nums[r] = nums[l]
		//fmt.Printf("after swap right: %v, pivot: %d\n ", nums, pivot)

		if l == r {
			nums[l] = pivot
		}
		//fmt.Printf("res:  %v\n", nums)

	}
	sortArray(nums[:l])
	sortArray(nums[l+1:])
	//fmt.Println(nums)

	return nums
}

func sortArray2(nums []int) []int {
	nLen := len(nums)
	if nLen < 2 {
		return nums
	}
	left := 0
	right := nLen - 1
	pindex := 0
	pivot := nums[pindex]
	//l := left
	//r := right

	for left < right {
		for nums[right] >= pivot && left < right {
			right--
		}

		for nums[left] <= pivot && left < right {
			left++
		}

		if left == right {
			nums[left], nums[pindex] = nums[pindex], nums[left]
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	//fmt.Println(l, r)
	sortArray2(nums[:left])
	sortArray2(nums[left+1:])

	return nums
}

func sortArray3(nums []int) []int {
	nLen := len(nums)
	if nLen < 2 {
		return nums
	}
	left := 0
	right := nLen - 1
	pindex := nLen - 1
	pivot := nums[pindex]

	for left < right {
		for nums[left] <= pivot && left < right {
			left++
		}
		for nums[right] >= pivot && left < right {
			right--
		}
		if left == right {
			nums[left], nums[pindex] = nums[pindex], nums[left]
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	sortArray3(nums[:left])
	sortArray3(nums[left+1:])

	return nums
}

func main() {
	nums := []int{3, 5, 20, 1, 2, 9, -10, 7, 6}

	//fmt.Println(sortArray(nums))
	fmt.Println(sortArray3(nums))

}
