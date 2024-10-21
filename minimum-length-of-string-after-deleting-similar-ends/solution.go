package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumLength("ca"))
	fmt.Println(minimumLength("cabaabac"))
	fmt.Println(minimumLength("aabccabba"))
}

func minimumLength(s string) int {
    left, right := 0, len(s) - 1

    for left < right {
        if s[left] == s[right] {
            char := s[left]
            for left <= right && s[left] == char {
                left++
            }
            for left <= right && s[right] == char {
                right--
            }
        } else {
            return right - left + 1
        }
    }

    if left == right {
        return 1
    }

    return 0
}