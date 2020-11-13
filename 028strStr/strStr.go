package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if haystack == needle || needle == "" {
		fmt.Println(haystack, needle, 0)

		return 0
	}
	if haystack == "" || len(haystack) < len(needle) {
		fmt.Println(haystack, needle, -1)

		return -1
	}

	//h := []rune(haystack)
	//n := []rune(needle)

	rflag := true
	//pubstr := ""

	for index, l := range haystack {

		if string(l) == string(needle[0]) {
			rflag = true
			for nindex, nl := range needle {
				if index+nindex >= len(haystack) || string(haystack[index+nindex]) != string(nl) {
					rflag = false
				}
			}
			if rflag {
				fmt.Println(haystack, needle, index)
				return index
			}
		}

	}

	fmt.Println(haystack, needle, -1)
	return -1
}

func strStr2(haystack string, needle string) int {
	if haystack == needle || needle == "" {
		fmt.Println(haystack, needle, 0)
		return 0
	}

	ln := len(needle)

	for i := 0; i <= len(haystack)-ln; i++ {
		if haystack[i:i+ln] == needle {
			fmt.Println(haystack, needle, i)

			return i
		}
	}
	fmt.Println(haystack, needle, -1)

	return -1
}

func main() {
	strStr2("hello", "ll")
	strStr2("aaaaa", "bba")
	strStr2("aaa", "aaaa")
	strStr2("aaa", "aaa")
	strStr2("", "")
	strStr2("a", "a")
	strStr2("mississippi", "issipi")
	strStr2("mississippi", "issip")
	strStr2("mississippi", "pi")
}
