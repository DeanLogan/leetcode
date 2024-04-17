package main 

import (
	"fmt"
	"math"
)

func main() {
    fmt.Println(minWindow("ADOBECODEBANC", "ABC")) // Expected output: "BANC"
    fmt.Println(minWindow("a", "a")) // Expected output: "a"
    fmt.Println(minWindow("a", "aa")) // Expected output: ""
}

func minWindow(s string, t string) string {
    if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
        return ""
    }

    sCount := make([]int, 128)
    count := len(t)
    left, right := 0, 0
    minLen, startIndex := math.MaxInt64, 0

    for _, char := range t {
        sCount[char]++
    }

    for right < len(s) {
        if sCount[s[right]] > 0 {
            count--
        }
        sCount[s[right]]--
        right++

        for count == 0 {
            if right-left < minLen {
                startIndex = left
                minLen = right - left
            }

            if sCount[s[left]] == 0 {
                count++
            }
            sCount[s[left]]++
            left++
        }
    }

    if minLen == math.MaxInt64 {
        return ""
    }

    return s[startIndex : startIndex+minLen]
}
