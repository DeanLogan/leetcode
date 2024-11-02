package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfPairs([]int{1,3,4}, []int{1,3,4}, 1))
	fmt.Println(numberOfPairs([]int{1,2,4,12}, []int{2,4}, 3))
}

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
    goodPairs := int64(0)
    occurrences := make(map[int]int)
    
    for i := range nums2 {
        occurrences[nums2[i]*k]++
    }
    
    for i := range nums1 {
        for j := 1; j*j <= nums1[i]; j++ {
            if nums1[i]%j == 0 {
                goodPairs += int64(occurrences[j])
                complement := nums1[i] / j
                if complement != j && occurrences[complement] != 0 {
                    goodPairs += int64(occurrences[complement])
                }
            }
        }
    }
    return goodPairs
}