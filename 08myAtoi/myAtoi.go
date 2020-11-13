package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	sig, l := 1, 0
	switch str[l] {
	case '-':
		sig = -1
		str = str[1:]
	case '+':
		str = str[1:]
		sig = 1
	default:
		sig = 1
	}

	ret := 0
	for _, v := range str {
		if v < '0' || v > '9' {
			break
		}
		ret = ret*10 + int(v-'0')
		if ret > math.MaxInt32 {
			if sig == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return ret * sig
}

func main() {
	fmt.Println(myAtoi("3.14"))
	fmt.Println(myAtoi("2147483648"))
	fmt.Println(myAtoi("-2147483648"))
	fmt.Println(myAtoi("+1"))
	fmt.Println(myAtoi(" 0000000000012345678"))
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193  with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("+-1"))
	fmt.Println(myAtoi("-+1"))
	fmt.Println(myAtoi("9223372036854775808"))
}
