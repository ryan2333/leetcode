package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var prefix string = ""

	if len(strs) == 0 {
		return prefix
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var minLen = len(strs[0])
	var minStr = strs[0]

	for _, str := range strs {
		if str == "" {
			return ""
		}
		if len(str) < minLen {
			minLen = len(str)
			minStr = str
		}
	}
	for i := 0; i < minLen; i++ {
		tag := true
		for _, sub := range strs {
			if sub[i] != minStr[i] {
				tag = false
				break
			}
		}
		if !tag {
			break
		}
		prefix += string(minStr[i])
	}
	fmt.Println(prefix)
	return prefix
}

func main() {
	longestCommonPrefix([]string{"flower", "flow", "flight"})
	longestCommonPrefix([]string{"dog", "racecar", "car"})
	longestCommonPrefix([]string{"dog", "", "car"})

}
