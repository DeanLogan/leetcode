package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(alternateDigitSum(521))
	fmt.Println(alternateDigitSum(111))
	fmt.Println(alternateDigitSum(886996))
}

func alternateDigitSum(n int) int {
    strNum := strconv.Itoa(n)
	n = 0
	for i, charNum := range strNum {
		num := int(charNum - '0')
		if i % 2 != 0 {
			num *= -1
		}
		n += num
	}
	return n
}