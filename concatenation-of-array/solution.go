package main

import (
	"fmt"
)

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 1}))
	fmt.Println(getConcatenation([]int{1, 3, 2, 1}))
}

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}