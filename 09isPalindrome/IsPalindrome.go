package main

import "fmt"

func IsPalindrome(s int) bool {
	var n int
	tmp := s
	for tmp > 0 {
		n = n*10 + tmp%10
		tmp /= 10
	}
	fmt.Println(s, n)
	return n == s
}

func main() {
	fmt.Println(IsPalindrome(121))
	fmt.Println(IsPalindrome(-121))
	fmt.Println(IsPalindrome(0))

}
