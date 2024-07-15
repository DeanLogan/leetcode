package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	shortestWord := 200 
    for _, word := range strs {
		if len(word) < shortestWord {
			shortestWord = len(word)
		}
	}
	for i := 0; i < shortestWord; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[j-1][i] {
				return strs[j][:i]
			}
		}
	}
	return strs[0][:shortestWord]
}