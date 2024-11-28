package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
	fmt.Println(findTheLongestSubstring("leetcodeisgreat"))
	fmt.Println(findTheLongestSubstring("bcbcbc"))
}

func findTheLongestSubstring(s string) int {
    indices := [32]int{}
    for i := 1; i < 32; i++  {
        indices[i] = -2
    }
    indices[0] = -1
    maxLen, occurrence := 0, 0
    for i := 0; i < len(s); i++ {
        switch s[i] {
            case 'a':
                occurrence ^= 1
            case 'e':
                occurrence ^= (1 << 1)
            case 'i':
                occurrence ^= (1 << 2)
            case 'o':
                occurrence ^= (1 << 3)
            case 'u':
                occurrence ^= (1 << 4)
        }
        if indices[occurrence] == -2 {
            indices[occurrence] = i
        } else {
            maxLen = max(maxLen, i - indices[occurrence])
        }
    }
    return maxLen
}