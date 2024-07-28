package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(maximumStrongPairXor([]int{1,2,3,4,5}))
	fmt.Println(maximumStrongPairXor([]int{10,100}))
	fmt.Println(maximumStrongPairXor([]int{5,6,25,30}))
}

func maximumStrongPairXor(nums []int) int {
    sort.Ints(nums)
    maxXor := 0
    n := len(nums)

    for i := 0; i < n; i++ {
        for j := i; j < n && nums[j] - nums[i] <= nums[i]; j++ {
            currentXor := nums[i] ^ nums[j]
            if currentXor > maxXor {
                maxXor = currentXor
            }
        }
    }

    return maxXor
}