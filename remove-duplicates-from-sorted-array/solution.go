package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    uniqueElement := nums[0]
    uniqueCount := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != uniqueElement {
            nums[uniqueCount] = nums[i]
            uniqueCount++
        }
        uniqueElement = nums[i]
    }
    return uniqueCount
}