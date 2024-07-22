package main

import (
	"fmt"
)

func main(){
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
    start := 0
    end := len(s) - 1
    newStringArr := []byte(s)

    for start < end {
        if isVowel(newStringArr[start]) && isVowel(newStringArr[end]) {
            newStringArr[start], newStringArr[end] = newStringArr[end], newStringArr[start]
            start++
            end--
        }
        if !isVowel(newStringArr[start]) {
            start++
        }
        if !isVowel(newStringArr[end]) {
            end--
        }
    }
    return string(newStringArr)
}

func isVowel(c byte) bool {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}