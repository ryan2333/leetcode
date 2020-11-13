package main

import (
	"fmt"
	"time"
)

func firstUniqChar1(s string) int {
	start := time.Now()
	defer func() {
		fmt.Println(time.Now().Sub(start))
	}()
	if len(s) == 0 {
		return -1
	}
	if len(s) == 1 {
		return 0
	}

	mp := make(map[string]int, len(s))

	for i, ss := range s {
		isDup := false
		if _, ok := mp[string(ss)]; ok {
			continue
		}
		for _, sss := range s[i+1:] {
			if sss == ss {
				isDup = true
				mp[string(ss)] = 1
			}
		}
		if !isDup {
			return i
		}
	}
	fmt.Println(s, -1)
	return -1
}

func firstUniqChar(s string) int {
	start := time.Now()
	defer func() {
		fmt.Println(time.Now().Sub(start))
	}()
	if s == "" {
		return -1
	}

	mp := make(map[string]int, len(s))
	for _, ss := range s {
		if _, ok := mp[string(ss)]; ok {
			mp[string(ss)]++
			continue
		}
		mp[string(ss)] = 1
	}
	for i, ss := range s {
		if mp[string(ss)] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	firstUniqChar("leetcode")
	firstUniqChar("loveleetcode")
	firstUniqChar("aadadaad")
	firstUniqChar("a")
	firstUniqChar("")
}
