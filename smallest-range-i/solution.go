package main

import (
	"fmt"
)

func main(){
	fmt.Println(smallestRangeI([]int{1}, 0))
	fmt.Println(smallestRangeI([]int{0,10}, 2))
	fmt.Println(smallestRangeI([]int{1,3,6}, 3))
}

func smallestRangeI(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	maxNum, minNum := nums[0], nums[0]

	for _, num := range nums[1:] {
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}

	maxVal := maxNum - k
	minVal := minNum + k
	score := maxVal - minVal

	if score < 0 {
		return 0
	}
	return score
}