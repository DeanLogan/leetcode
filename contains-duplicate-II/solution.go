package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1,0,1,1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1,2,3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	lastIndexMap := make(map[int]int)

	for i, num := range nums {
		if lastIndex, ok := lastIndexMap[num]; ok && i-lastIndex <= k {
			return true
		}
		lastIndexMap[num] = i
	}

	return false
}