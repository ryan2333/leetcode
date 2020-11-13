package main

import "fmt"

func isPowerOfFour(num int) bool {
	if num < 1 {
		return false
	}

	var check = 1

	for num >= check {
		if num == check {
			return true
		}
		check <<= 2
	}

	return false
}

func isPowerOfFour1(num int) bool {
	if num < 1 {
		return false
	}

	if num&(num-1) == 0 && (num&0x55555555) == num {
		return true
	}

	return false
}

func isPowerOfFour2(num int) bool {
	if num < 1 {
		return false
	}

	for num > 1 && num%4 == 0 {
		num /= 4
	}

	return num == 1
}

func main() {
	fmt.Println(isPowerOfFour1(4))
	fmt.Println(4 & 3)
	fmt.Println(isPowerOfFour1(64))

}
