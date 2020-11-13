package main

import "fmt"

func reverseString1(s []byte) {
	fmt.Println("s: ", s)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	fmt.Println("s: ", s)
}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	reverseString([]byte("hello"))
	reverseString([]byte("Hannah"))
}
