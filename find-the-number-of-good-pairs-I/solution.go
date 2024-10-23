package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfPairs([]int{1,3,4}, []int{1,3,4}, 1))
	fmt.Println(numberOfPairs([]int{1,2,4,12}, []int{2,4}, 3))
}

func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	goodPairs := 0

    for i := range nums1 {
		for j := range nums2 {
			if nums1[i] % (nums2[j] * k)  == 0 {
				goodPairs++
			}
		}
	}

	return goodPairs;
}
