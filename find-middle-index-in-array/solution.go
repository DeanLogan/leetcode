package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMiddleIndex([]int{2,3,-1,8,4}))
	fmt.Println(findMiddleIndex([]int{1,-1,4}))
	fmt.Println(findMiddleIndex([]int{2,5}))
}

func findMiddleIndex(nums []int) int {
    totalSum := 0
    for _, num := range nums {
        totalSum += num 
    }
    leftSum := 0
    for index, num := range nums {
        totalSum -= num
        if totalSum == leftSum { 
            return index 
        }
        leftSum += num
    }
    return -1
}