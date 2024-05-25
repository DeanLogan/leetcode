package main

import (
	"fmt"
)

func main(){
	fmt.Println(longestConsecutive([]int{100,4,200,1,3,2}))
	fmt.Println(longestConsecutive([]int{0,3,7,2,5,8,4,6,0,1}))
}

func longestConsecutive(nums []int) int {
    numsMap := make(map[int]bool)

    for _, num := range nums {
        numsMap[num] = true
    }
    
    maxSequence := 0 
    for num := range numsMap {
        if !numsMap[num-1] {
            sequenceCount := 1

            for numsMap[num+1] {
                num++
                sequenceCount++
            }

            if maxSequence < sequenceCount {
                maxSequence = sequenceCount
            }
        }
    }
    return maxSequence
}