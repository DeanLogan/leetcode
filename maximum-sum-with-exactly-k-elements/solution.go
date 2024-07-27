package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximizeSum([]int{1,2,3,4,5}, 3))
	fmt.Println(maximizeSum([]int{5,5,5}, 2))
}

func maximizeSum(nums []int, k int) int {
	score := 0
	maxElement := 0

    for i := 0; i < len(nums); i++ {
        if maxElement < nums[i] {
            maxElement = nums[i]
        }
    }

    for i := 0; i < k; i++ {
        score += (maxElement+i)
    }

    return score
}