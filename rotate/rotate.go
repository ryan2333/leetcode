package main

import "fmt"

func rotate1(nums []int, k int) {
	lenN := len(nums)
	nums = append(nums[lenN-k:], nums[:lenN]...)
	nums = nums[:lenN]
	fmt.Println(nums)
}

func rotate2(nums []int, k int) {
	//lenN := len(nums)
	for _, n := range nums[:k] {
		nums = append(nums, n)
	}

	nums = nums[k:]
	fmt.Println(nums)
}

func rotate3(nums []int, k int) {
	for i := 0; i < k; i++ {
		nums = append([]int{nums[len(nums)-1]}, nums[:len(nums)-1]...)
	}
	fmt.Println(nums)
}
func rotate4(nums []int, k int) {
	lp := len(nums)
	//if lp > 1 {
	k %= lp
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
	//}

	fmt.Println(nums)
}

func main() {
	rotate4([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate4([]int{-1, 100, 3, 99}, 2)
	rotate4([]int{-1}, 2)
	rotate4([]int{1, 2}, 3)
}
