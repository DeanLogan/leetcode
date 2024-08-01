package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTheDifference([]string{"adc","wzy","abc"}))
	fmt.Println(findTheDifference([]string{"aaa","bob","ccc","ddd"}))
}

func findTheDifference(words []string) string {
    if len(words) <= 1 {
        return ""
    }

    calculateDifference := func(word string) []int {
        difference := make([]int, len(word)-1)
        for i := 0; i < len(word)-1; i++ {
            difference[i] = int(word[i+1]) - int(word[i])
        }
        return difference
    }
    differenceMap := make(map[string][]string)
    for _, word := range words {
        diffArray := calculateDifference(word)
        diffKey := fmt.Sprint(diffArray) 
        differenceMap[diffKey] = append(differenceMap[diffKey], word)
    }
    for _, group := range differenceMap {
        if len(group) == 1 {
            return group[0]
        }
    }

    return ""
}