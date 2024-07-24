package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(frequencySort([]int{1,1,2,2,2,3}))
	fmt.Println(frequencySort([]int{1,3,3,2,2}))
	fmt.Println(frequencySort([]int{-1,1,-6,4,5,-6,1,4,1}))
}

func frequencySort(nums []int) []int {
    freqMap := make(map[int]int)
    for _, num := range nums {
        freqMap[num]++
    }

    sort.Slice(nums, func(i, j int) bool {
        if freqMap[nums[i]] == freqMap[nums[j]] {
            return nums[i] > nums[j] 
        }
        return freqMap[nums[i]] < freqMap[nums[j]] 
    })

    return nums
}