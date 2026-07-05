package main

import (
	"fmt"
)

func main() {
	fmt.Println(minPartitions("32"))
	fmt.Println(minPartitions("82734"))
	fmt.Println(minPartitions("27346209830709182346"))
}

func minPartitions(n string) int {
    max_digit := 0;
	for _, char := range n {
		digit := int(char)-48;
		if digit > max_digit {
			max_digit = digit;
		}
	}
	return max_digit
}