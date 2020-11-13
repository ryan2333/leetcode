package main

import "fmt"

func plusOne1(digits []int) []int {
	dLen := len(digits)
	if dLen == 0 {
		digits[0] = 1
		fmt.Println(digits)
		return digits
	}

	if digits[dLen-1] < 9 {
		digits[dLen-1]++
		fmt.Println(digits)

		return digits
	}

	var j, s int
	for i := dLen - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + j
			fmt.Println(digits)
			return digits
		}
		if i == dLen-1 {
			s = digits[i] + 1
		} else {
			s = digits[i] + j
		}

		if s > 9 {
			digits[i] = s % 10
			j = s / 10
		} else {
			digits[i] = s
			j = 0
		}
	}
	if j == 1 {
		digits = append([]int{j}, digits...)
	}

	fmt.Println(digits)
	return digits
}

func plusOne2(digits []int) []int {
	var (
		flag int
		lp   int
	)

	lp = len(digits)
	flag = 1

	for i := lp - 1; i >= 0; i-- {
		flag, digits[i] = (digits[i]+flag)/10, (digits[i]+flag)%10
	}

	if flag == 1 {
		digits = append([]int{1}, digits...)
	}
	fmt.Println("ret: ", digits)

	return digits
}

func plusOne(digits []int) []int {
	var (
		lp  int
		n   int
		ret []int
	)
	lp = len(digits)
	if lp == 0 {
		return []int{1}
	}
	for n = lp - 1; n >= 0; n-- {
		if digits[n] != 9 {
			digits[n]++
			fmt.Println("digits: ", digits)
			return digits
		} else {
			digits[n] = 0
		}
	}
	ret = make([]int, len(digits)+1)
	ret[0] = 1
	fmt.Println("ret: ", ret)
	return ret
}

func main() {
	plusOne([]int{1, 2, 3})
	plusOne([]int{4, 3, 2, 1})
	plusOne([]int{9, 9, 9})
}
