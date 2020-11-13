package main

import "fmt"

func fizzBuzz1(n int) []string {
	ret := make([]string, n)
	for i := 0; i < n; i++ {
		if (i+1)%15 == 0 {
			ret[i] = "FizzBuzz"
		} else if (i+1)%5 == 0 {
			ret[i] = "Buzz"
		} else if (i+1)%3 == 0 {
			ret[i] = "Fizz"
		} else {
			ret[i] = fmt.Sprintf("%d", i+1)
		}
	}
	fmt.Println("ret: ", ret)
	return ret
}

func fizzBuzz(n int) []string {
	ret := make([]string, n)

	for i := 1; i <= n; i++ {
		tmp := ""
		if i%3 == 0 {
			tmp += "Fizz"
		}
		if i%5 == 0 {
			tmp += "Buzz"
		}
		if tmp == "" {
			tmp = fmt.Sprintf("%d", i)
		}
		ret[i-1] = tmp
	}
	return ret
}

func main() {
	fizzBuzz(1)
	fizzBuzz(10)
	fizzBuzz(15)
}
