package main

import (
    "fmt"
)

func main() {
	fmt.Println(check([]int{3,4,5,1,2}))
	fmt.Println(check([]int{2,1,3,4}))
	fmt.Println(check([]int{1,2,3}))
}

func check(nums []int) bool {
    count := 0
    numsLen := len(nums)

    for i := 0; i < numsLen; i++ {
        if nums[i] > nums[(i+1)%numsLen] {
            count++
            if count > 1 {
                return false
            }
        }
    }

    return count <= 1
}