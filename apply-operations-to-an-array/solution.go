package main

import (
	"fmt"
)

func main(){
	fmt.Println(applyOperations([]int{1,2,2,1,1,0}))
	fmt.Println(applyOperations([]int{0,1}))
}

func applyOperations(nums []int) []int {
    numLen := len(nums)
    
    for i := 0; i < numLen-1; i++ {
        if nums[i] == nums[i+1] {
            nums[i] *= 2
            nums[i+1] = 0
        }
    }
    
    result := make([]int, numLen)
    index := 0
    
    for _, num := range nums {
        if num != 0 {
            result[index] = num
            index++
        }
    }
    
    return result
}