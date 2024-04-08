package main

import (
	"fmt"
)

func main(){
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trap([]int{4,2,0,3,2,5}))
    fmt.Println(trap([]int{2,0,2}))
}

// solution with memory usage O(1)
func trap(height []int) int {
	left, right := 0, len(height) - 1
    maxLeft, maxRight := height[left], height[right]
    trapped := 0

    for left < right {
        if maxLeft < maxRight {
            left++
            heightNum := height[left]
            if heightNum > maxLeft {
                maxLeft = heightNum
            }
            trapped += maxLeft - heightNum
        } else {
            right--
            heightNum := height[right]
            if heightNum > maxRight {
                maxRight = heightNum
            }
            trapped += maxRight - heightNum
        }
        
    }
	return trapped
}