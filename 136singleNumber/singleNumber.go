package main

import "fmt"

func singleNumber1(nums []int) int {
	var (
		snum   int
		num    int
		numMap map[int]int
	)
	numMap = make(map[int]int)
	for _, num = range nums {
		if _, ok := numMap[num]; !ok {
			numMap[num] = 1
		} else {
			numMap[num]++
		}
	}
	for k, v := range numMap {
		if v == 1 {
			snum = k
			break
		}
	}
	fmt.Println("num: ", snum)
	numMap = nil
	return snum
}

func singleNumber2(nums []int) int {
	var (
		ln   int
		nmap map[int]bool
		ok   bool
	)
	ln = len(nums)
	nmap = make(map[int]bool, ln)
	for i := 0; i < ln; i++ {
		if _, ok = nmap[nums[i]]; ok {
			continue
		}
		for j := i + 1; j < ln; j++ {
			if nums[j] == nums[i] {
				ok = true
				nmap[nums[i]] = true
				break
			}
		}
		if !ok {
			fmt.Println("snum: ", nums[i])
			return nums[i]
		}
	}
	return 0
}

/**
位运算
任何数与0做异或运算，结果为原来数本身, 即： a ^ 0 = a
任何数和自身做运算，结果为0, 即: a ^ a = 0
异或运算满足交换律和结合律，即：a ^ b ^ a == a ^ a ^ b ==  (a ^ a) ^ b  == 0 ^ b == b
*/
func singleNumber(nums []int) int {
	var one, two = 0, 0
	for _, n := range nums {
		fmt.Printf("^%d & (%d ^ %d) == %d & %d \n", ^one, two, n, one, one^n)
		one = ^two & (one ^ n)
		fmt.Println("one: ", one)
		fmt.Printf("^%d & (%d ^ %d) == %d & %d \n", ^two, one, n, two, two^n)
		two = ^one & (two ^ n)
		fmt.Println("two: ", two)

	}
	fmt.Println("snum: ", one)

	return one
}

func main() {
	singleNumber([]int{2, 2, 2, 1})
	//singleNumber([]int{4, 1, 2, 1, 1, 2, 2})
}
