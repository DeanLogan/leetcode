package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(addedInteger([]int{2,6,4}, []int{9,7,5}))
	fmt.Println(addedInteger([]int{10}, []int{5}))
	fmt.Println(addedInteger([]int{1,1,1,1}, []int{1,1,1,1}))
}

func addedInteger(nums1 []int, nums2 []int) int {
	sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
	sort.Slice(nums2, func(i, j int) bool { return nums2[i] < nums2[j] })

	return nums2[0] - nums1[0]
}