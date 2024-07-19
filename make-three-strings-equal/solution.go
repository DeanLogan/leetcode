package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(findMinimumOperations("abc", "abb", "ab"))
	fmt.Println(findMinimumOperations("dac", "bac", "cac"))
	fmt.Println(findMinimumOperations("a", "a", "a"))
}

func findMinimumOperations(s1 string, s2 string, s3 string) int {
	shortestStr := ""
	tempStr := s1 + "#"
	i := 0
	for strings.HasPrefix(s1, shortestStr) && strings.HasPrefix(s2, shortestStr) && strings.HasPrefix(s3, shortestStr) {
		shortestStr += string(tempStr[i])
		i++
	}
	shortestStrLen := len(shortestStr) -1 
	if shortestStrLen == 0 {
		return -1
	}
	return (len(s1) - shortestStrLen) + (len(s2) - shortestStrLen) + (len(s3) - shortestStrLen)
}