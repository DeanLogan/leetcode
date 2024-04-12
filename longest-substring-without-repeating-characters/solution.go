package main

import (
	"fmt"
)

func main(){
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) 
	fmt.Println(lengthOfLongestSubstring("bbbbb")) 
	fmt.Println(lengthOfLongestSubstring("pwwkew")) 
}

func lengthOfLongestSubstring(s string) int {
    longestSequence := 0
    sequenceArray := make([]int, 128)
    for i := range sequenceArray {
        sequenceArray[i] = -1
    }
    for left, right := 0, 0; right < len(s); right++ {
        if index := sequenceArray[s[right]]; index >= left {
            left = index + 1
        }
        sequenceArray[s[right]] = right
        if right - left + 1 > longestSequence {
            longestSequence = right - left + 1
        }
    }
    return longestSequence
}