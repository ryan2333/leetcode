package main

import (
	"fmt"
	"time"
)

func isAnagram(s string, t string) bool {
	start := time.Now()
	defer func() {
		fmt.Println("cost: ", time.Now().Sub(start))
	}()

	if len(s) != len(t) {
		return false
	}

	if s == t && len(s) <= 1 {
		return true
	}

	mps := make(map[string]int, len(t))
	for _, ss := range s {
		if _, ok := mps[string(ss)]; ok {
			mps[string(ss)]++
		} else {
			mps[string(ss)] = 1
		}
	}

	for _, ss := range t {
		mps[string(ss)]--
		if mps[string(ss)] == 0 {
			delete(mps, string(ss))
		}
	}
	if len(mps) == 0 {
		return true
	}

	return false

}

func main() {
	fmt.Println(isAnagram("hehe", "eheh"))
	fmt.Println(isAnagram("", ""))
	fmt.Println(isAnagram("a", "a"))
	fmt.Println(isAnagram("hh", ""))
	fmt.Println(isAnagram("car", "rat"))
	fmt.Println(isAnagram("anagram", "nagaram"))
}
