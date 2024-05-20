package main

import (
	"fmt"
)

func main(){
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
	fmt.Println(largestRectangleArea([]int{2,4}))
}

func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }
    stack := make([]int, 0, len(heights))
    maxArea := 0

    for i, h := range heights {
        for len(stack) != 0 && h < heights[stack[len(stack)-1]] {
            height := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            width := i
            if len(stack) != 0 {
                width = i - stack[len(stack)-1] - 1
            }
            area := height * width
            if area > maxArea {
                maxArea = area
            }
        }
        stack = append(stack, i)
    }

    for len(stack) != 0 {
        height := heights[stack[len(stack)-1]]
        stack = stack[:len(stack)-1]
        width := len(heights)
        if len(stack) != 0 {
            width = len(heights) - stack[len(stack)-1] - 1
        }
        area := height * width
        if area > maxArea {
            maxArea = area
        }
    }

    return maxArea
}
