package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("leetcode", "leeto"))
}

func strStr(haystack string, needle string) int {
	lenOfNeedle := len(needle)
    for i:=0; i<len(haystack); i++ {
		if (lenOfNeedle+i) > len(haystack) {
			break
		}
		if haystack[i:lenOfNeedle+i] == needle {
			return i
		}
	}
	return -1
}