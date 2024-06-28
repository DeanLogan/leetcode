package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
	fmt.Println(largestSumAfterKNegations([]int{3,-1,0,2}, 3))
	fmt.Println(largestSumAfterKNegations([]int{2,-3,-1,5,-4}, 2))
}

func largestSumAfterKNegations(nums []int, k int) int {
    sort.Ints(nums)
    smallestAbs := abs(nums[0])
    for i := range nums {
        if nums[i] < 0 && k > 0 {
            nums[i] = -nums[i]
            k--
        }
        if abs(nums[i]) < smallestAbs {
            smallestAbs = abs(nums[i])
        }
    }
    if k%2 != 0 {
        return sum(nums) - 2*smallestAbs 
    }
    return sum(nums)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func sum(nums []int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}