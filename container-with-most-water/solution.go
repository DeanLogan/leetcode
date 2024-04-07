package main

import (
	"fmt"
)

func main(){
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
	fmt.Println(maxArea([]int{1,1}))
}

func maxArea(height []int) int {
    left, right := 0, len(height) - 1
    max := 0

    for left < right {
        heightNum := 0
        if height[left] < height[right] {
            heightNum = height[left]
            left++
        } else {
            heightNum = height[right]
            right--
        }

        area := (right - left + 1) * heightNum

        if area > max {
            max = area
        }
    }

    return max
}