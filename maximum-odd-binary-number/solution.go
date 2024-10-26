package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maximumOddBinaryNumber("010"))
	fmt.Println(maximumOddBinaryNumber("0101"))
}

func maximumOddBinaryNumber(s string) string {
	count := strings.Count(s, "1")
	length := len(s)

	result := make([]byte, length)
	for i := range result {
		if i < count-1 {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}
	
	result[length-1] = '1'
	return string(result)
}