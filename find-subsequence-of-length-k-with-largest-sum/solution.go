package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(maxSubsequence([]int{2,1,3,3},2))
	fmt.Println(maxSubsequence([]int{-1,-2,3,4},3))
	fmt.Println(maxSubsequence([]int{3,4,3,3},2))
}

func maxSubsequence(nums []int, k int) []int {
    indexedNums := make([][2]int, len(nums))
    for i, num := range nums {
        indexedNums[i] = [2]int{num, i}
    }

    sort.Slice(indexedNums, func(i, j int) bool {
        return indexedNums[i][0] > indexedNums[j][0]
    })

    sort.Slice(indexedNums[:k], func(i, j int) bool {
        return indexedNums[i][1] < indexedNums[j][1]
    })

    result := make([]int, k)
    for i := 0; i < k; i++ {
        result[i] = indexedNums[i][0]
    }

    return result
}