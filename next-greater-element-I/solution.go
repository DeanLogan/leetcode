package main

import (
	"fmt"
)

func main() {
	fmt.Println(nextGreaterElement([]int{4,1,2}, []int{1,3,4,2}))
	fmt.Println(nextGreaterElement([]int{2,4}, []int{1,3,4,2}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    nextGreaterMap := make(map[int]int)
    var stack []int

    for _, num := range nums2 {
        for len(stack) > 0 && num > stack[len(stack)-1] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            nextGreaterMap[top] = num
        }
        stack = append(stack, num)
    }

    for _, num := range stack {
        nextGreaterMap[num] = -1
    }
    result := make([]int, len(nums1))
    for i, num := range nums1 {
        result[i] = nextGreaterMap[num]
    }

    return result
}