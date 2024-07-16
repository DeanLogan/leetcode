package main

import (
	"fmt"
)

func main(){
	fmt.Println(findIntersectionValues([]int{2,3,2}, []int{1,2}))
	fmt.Println(findIntersectionValues([]int{4,3,2,3,1}, []int{2,2,5,2,3,6}))
	fmt.Println(findIntersectionValues([]int{3,4,2,3}, []int{1,5}))
}

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	nums1In2 := 0
	nums2In1 := 0
	for _, num := range nums1 {
		for _, num2 := range nums2 {
			if num == num2 {
				nums1In2++
				break
			}
		}
	}
	for _, num2 := range nums2 {
		for _, num1 := range nums1 {
			if num2 == num1 {
				nums2In1++
				break
			}
		}
	}

	return []int{nums1In2, nums2In1}
}