package main

import (
	"fmt"
)

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
	fmt.Println(findNumbers([]int{555,901,482,1771}))
}

func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		digitCount := 0
		for num != 0 {
			num /= 10
			digitCount++
		}
		if digitCount%2 == 0 {
			count++
		}
	}
	return count
}