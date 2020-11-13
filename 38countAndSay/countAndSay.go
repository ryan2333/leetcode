package main

import "fmt"

func countAndSay(n int) string {
	base := "1"
	count := 1
	for count < n {
		base = calcList(base)
		count++
	}
	return base
}

func calcList(s string) string {
	ret := ""
	start := 0
	letter := s[0]
	for start < len(s) {
		if s[start] == letter {
			start++
		} else {
			ret += fmt.Sprintf("%d%s", start, string(letter))
			s = s[start:]
			start = 0
			letter = s[0]
		}
	}
	ret += fmt.Sprintf("%d%s", start, string(letter))
	//fmt.Println("ret: ", ret)
	return ret
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(countAndSay(i + 1))
	}
}
