package main

import (
	"fmt"
)

func main() {
	fmt.Println(pivotIndex([]int{1,7,3,6,5,6}))
	fmt.Println(pivotIndex([]int{1,2,3}))
	fmt.Println(pivotIndex([]int{2,1,-1}))
}

func pivotIndex(nums []int) int {
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