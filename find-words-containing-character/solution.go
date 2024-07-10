package main

import (
	"fmt"
)

func main(){
	fmt.Println(findWordsContaining([]string{"leet","code"}, 'e'))
	fmt.Println(findWordsContaining([]string{"abc","bcd","aaaa","cbc"}, 'a'))
	fmt.Println(findWordsContaining([]string{"abc","bcd","aaaa","cbc"}, 'z'))
}

func findWordsContaining(words []string, x byte) []int {
    xRune := rune(x)
    occurrences := make([]int, 0, len(words)) 
    for i, word := range words {
        for _, char := range word {
            if char == xRune {
                occurrences = append(occurrences, i)
                break
            }
        }
    }
    return occurrences
}