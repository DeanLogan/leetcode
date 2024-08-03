package main

import (
	"fmt"
)

func main(){
	fmt.Println(leftRightDifference([]int{10,4,8,3}))
	fmt.Println(leftRightDifference([]int{1}))
}

func leftRightDifference(nums []int) []int {
    n := len(nums)
    leftSum := make([]int, n)
    rightSum := make([]int, n)
    answer := make([]int, n)

    for i := 1; i < n; i++ {
        leftSum[i] = leftSum[i-1] + nums[i-1]
    }
    for i := n - 2; i >= 0; i-- {
        rightSum[i] = rightSum[i+1] + nums[i+1]
    }
	
    for i := 0; i < n; i++ {
        answer[i] = abs(leftSum[i] - rightSum[i])
    }

    return answer
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}