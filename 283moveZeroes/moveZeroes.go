package main

import "fmt"

func moveZeroes1(nums []int) {
	var (
		ln    int
		n     int
		count int
	)
	ln = len(nums)
	count = 0

	for n = 0; n < ln; n++ {
		if nums[n] == 0 {
			nums = append(nums[:n], nums[n+1:]...)
			count++
			ln--
			n--
		}
	}
	if count > 0 {
		nums = append(nums, make([]int, count)...)
	}
	fmt.Println("nums: ", nums)
}

func moveZeroes(nums []int) {
	var (
		ln   int
		fast int
		slow int
	)
	ln = len(nums)
	slow, fast = 0, 0

	for fast = 0; fast < ln; fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	if slow != fast {
		nums = append(nums[:slow], make([]int, fast-slow)...)
	}
	fmt.Println("nums: ", nums)
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0, 0, 1})
}
