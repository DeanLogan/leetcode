package main

import (
	"fmt"
)

func main(){
	fmt.Println(characterReplacement("ABAB", 2)) 
	fmt.Println(characterReplacement("AABABBA", 1)) 
}

func characterReplacement(s string, k int) int {
    count := make([]int, 26)
    maxCount, maxLen := 0, 0

    for left, right := 0, 0; right < len(s); right++ {
        count[s[right]-'A']++
        maxCount = max(maxCount, count[s[right]-'A'])
        if (right-left) + (1-maxCount) > k {
            count[s[left]-'A']--
            left++
        }
        maxLen = max(maxLen, right-left+1)
    }
    return maxLen
}