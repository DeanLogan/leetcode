package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBadPairs([]int{4,1,3,3}))
	fmt.Println(countBadPairs([]int{1,2,3,4,5}))
}

func countBadPairs(nums []int) int64 {
    countMap := make(map[int]int)
    badPairs := 0
    numsLen := len(nums)

    for i := 0; i < numsLen; i++ {
        diff := nums[i] - i
        if count, exists := countMap[diff]; exists {
            badPairs += count
        }
        countMap[diff]++
    }

    totalPairs := int64(numsLen * (numsLen - 1) / 2)
    return totalPairs - int64(badPairs)
}