package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
	fmt.Println(maxProduct([]int{1,5,4,5}))
	fmt.Println(maxProduct([]int{3,7}))
}

func maxProduct(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}
