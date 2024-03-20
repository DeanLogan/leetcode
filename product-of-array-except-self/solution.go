package main

import (
	"fmt"
)

func main(){
	fmt.Println(productExceptSelf([]int{1,2,3,4}))
	fmt.Println(productExceptSelf([]int{-1,1,0,-3,3}))
}

func productExceptSelf(nums []int) []int {
    answer := []int{1}

    for i := 1; i < len(nums); i++ {
        answer = append(answer, answer[i-1] * nums[i-1])
    }

    rightProduct := 1
    for i := len(nums) - 1; i >= 0; i-- {
        answer[i] *= rightProduct
        rightProduct *= nums[i]
    }

    return answer
}
