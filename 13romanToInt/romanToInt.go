package main

import (
	"fmt"
	"time"
)

func romanToInt1(s string) int {
	start := time.Now().UnixNano()
	romanMap := make(map[string]int)
	romanMap["I"] = 1
	romanMap["V"] = 5
	romanMap["X"] = 10
	romanMap["L"] = 50
	romanMap["C"] = 100
	romanMap["D"] = 500
	romanMap["M"] = 1000
	n := len(s)
	sum := 0

	for i := 0; i < n; i++ {
		if i == n-1 {
			sum += romanMap[string(s[i])]
			break
		}

		switch s[i] {
		case 'I':
			if s[i+1] == 'V' || s[i+1] == 'X' {
				sum = sum - romanMap["I"]
			} else {
				sum += romanMap["I"]
			}
		case 'X':
			if s[i+1] == 'L' || s[i+1] == 'C' {
				sum = sum - romanMap["X"]
			} else {
				sum += romanMap["X"]
			}
		case 'C':
			if s[i+1] == 'D' || s[i+1] == 'M' {
				sum = sum - romanMap["C"]
			} else {
				sum += romanMap["C"]
			}
		default:
			sum += romanMap[string(s[i])]

		}
	}
	end := time.Now().UnixNano()
	fmt.Println(end-start, sum)
	return sum
}

func romanToInt(s string) int {
	romanMap := make(map[string]int, 7)
	romanMap["I"] = 1
	romanMap["V"] = 5
	romanMap["X"] = 10
	romanMap["L"] = 50
	romanMap["C"] = 100
	romanMap["D"] = 500
	romanMap["M"] = 1000
	if len(s) == 1 {
		return romanMap[s]
	}
	sum := 0
	pre := romanMap[string(s[0])]
	for i := 1; i < len(s); i++ {
		cur := romanMap[string(s[i])]
		if pre < cur {
			sum -= pre
		} else {
			sum += pre
		}
		//fmt.Println(pre, cur, sum)
		pre = cur
	}
	sum += pre
	fmt.Println("sum: ", sum)
	return sum
}

func main() {
	romanToInt("III")
	romanToInt("IV")
	romanToInt("IX")
	romanToInt("LVIII")
	romanToInt("MCMXCIV")
}
