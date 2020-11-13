package main

import (
	"fmt"
	"strconv"
)

func addBinary1(a string, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	var sliceA []int
	var sliceB []int
	var res = ""

	for i := len(a) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(string(a[i]))
		sliceA = append(sliceA, num)
	}

	for i := len(b) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(string(b[i]))
		sliceB = append(sliceB, num)
	}

	var forLen = 0
	var tmp = 0
	la := len(a)
	lb := len(b)
	if la >= lb {
		forLen = lb
	} else {
		forLen = la
	}

	for i := 0; i < forLen; i++ {
		if sliceA[i]+sliceB[i]+tmp >= 2 {
			res = fmt.Sprintf("%d%s", (sliceA[i]+sliceB[i]+tmp)%2, res)
			tmp = 1
		} else {
			res = fmt.Sprintf("%d%s", sliceA[i]+sliceB[i]+tmp, res)
			tmp = 0
		}
	}

	if la > forLen {

		for i := forLen; i < la; i++ {
			if sliceA[i]+tmp >= 2 {
				res = fmt.Sprintf("%d%s", (sliceA[i]+tmp)%2, res)
				tmp = 1
			} else {
				res = fmt.Sprintf("%d%s", sliceA[i]+tmp, res)
				tmp = 0
			}
		}
	}
	if lb > forLen {
		for i := forLen; i < lb; i++ {
			if sliceB[i]+tmp >= 2 {
				res = fmt.Sprintf("%d%s", (sliceB[i]+tmp)%2, res)
				tmp = 1
			} else {
				res = fmt.Sprintf("%d%s", sliceB[i]+tmp, res)
				tmp = 0
			}
		}
	}
	if tmp == 1 {
		res = fmt.Sprintf("%d%s", tmp, res)
	}
	return res
}

func addBinary2(a, b string) string {
	an, _ := strconv.ParseInt(a, 2, 64)

	bn, _ := strconv.ParseInt(b, 2, 64)

	fmt.Printf("%b\n", an+bn)

	return ""
}

func main() {
	//a := "111"
	//b := "100"
	a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	addBinary1(a, b)
}
