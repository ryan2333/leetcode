package main

import (
	"fmt"
	"strings"
)

func isPalindrome1(s string) bool {
	if s == "" || len(s) == 1 {
		return true
	}
	s = strings.ToLower(s)
	tmps := ""
	for _, ss := range s {
		if (ss >= 'a' && ss <= 'z') || (ss >= '0' && ss <= '9') {
			tmps += string(ss)
		}
	}
	l := len(tmps)
	for i := 0; i < l/2; i++ {
		if tmps[i] != tmps[l-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isalNum(s[l]) {
			l++
		}
		for l < r && !isalNum(s[r]) {
			r--
		}
		if l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}

	return true
}

func isalNum(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9'
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	//fmt.Println(isPalindrome(""))
	//fmt.Println(isPalindrome("a"))
	//fmt.Println(isPalindrome("ab"))
	//fmt.Println(isPalindrome("aa"))
	//fmt.Println(isPalindrome("0P"))
}
