package main

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	sLen := len(s)
	lastIndex := -1

	for i := sLen - 1; i >= 0; i-- {
		if s[i] == ' ' && lastIndex != -1 {
			//fmt.Printf("%s, %d, %d, %s, %d\n", s, i, lastIndex, s[i+1:lastIndex+1], len(s[i+1:lastIndex+1]))
			return len(s[i+1 : lastIndex+1])
		}
		if s[i] != ' ' && lastIndex == -1 {
			lastIndex = i
			continue
		}
	}
	if lastIndex != -1 {
		return len(s[0 : lastIndex+1])
	}
	return 0
}

func main() {
	lengthOfLastWord("Hello World")
	lengthOfLastWord("hehe   ")
	lengthOfLastWord("  hh hh    ")
	lengthOfLastWord("   hh    ")
}
