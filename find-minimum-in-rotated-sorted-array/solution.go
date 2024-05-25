package main

import (
	"fmt"
)

func main(){ 
	fmt.Println(findMin([]int{3,4,5,1,2}))
	fmt.Println(findMin([]int{4,5,6,7,0,1,2}))
	fmt.Println(findMin([]int{11,13,15,17}))
}

func findMin(nums []int) int {
    left, right := 0, len(nums) - 1
    for left < right {
        point := (left + right) / 2
        if nums[point] > nums[right] {
            left = point + 1
        } else {
            right = point
        }
    }
    return nums[left]
}

