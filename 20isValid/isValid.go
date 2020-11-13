package main

import "fmt"

func isValid(s string) bool {
	var ss []int32
	sMap := make(map[int32]int32)
	sMap[')'] = '('
	sMap['}'] = '{'
	sMap[']'] = '['

	for _, l := range s {
		switch l {
		case '(', '[', '{':
			ss = append(ss, l)
		case ')', ']', '}':
			if len(ss) == 0 {
				return false
			}
			last := ss[len(ss)-1]
			if sMap[l] != last {
				return false
			}
			ss = ss[:len(ss)-1]
		}

	}
	if len(ss) != 0 {
		return false
	}

	return true
}
func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
