package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(rotateString("abcde", "cdeab"))
	fmt.Println(rotateString("abcde", "abced"))
}

func rotateString(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }
    concatenated := s + s
    return strings.Contains(concatenated, goal)
}