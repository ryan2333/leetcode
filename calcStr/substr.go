package main

import "fmt"

func calcStr(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	var slen int
	slen = len(s)
	smap := make(map[string]int, 5)
	smap["srt"] = 0
	smap["rt"] = 0

	for i := 0; i < slen; i++ {
		//fmt.Printf("for: %d, rt: %d\n", i, smap["rt"])
		smap[string(s[i])] = 1
		smap["srt"]++
		for j := i + 1; j < slen; j++ {
			_, ok := smap[string(s[j])]
			if ok {
				break
			}
			fmt.Printf("%s, %v\n", string(s[j]), ok)
			fmt.Println(smap)
			smap["srt"]++
		}
		if smap["srt"] > smap["rt"] {
			//fmt.Printf("rt: %d, srt: %d\n", smap["rt"], smap["srt"])
			smap["rt"] = smap["srt"]
		}
		smap["srt"] = 0
	}
	return smap["rt"]
}

func main() {
	//fmt.Println("abcabcbbb", calcStr("abcabcbbb"))
	fmt.Println("pwwkew", calcStr("pwwkew"))

}
