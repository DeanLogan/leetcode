package main

import (
	"fmt"
    "sort"
)

func main(){
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
	fmt.Println(threeSum([]int{0,1,1}))
	fmt.Println(threeSum([]int{0,0,0}))
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	result := [][]int{}
	for num1Idx := 0; num1Idx < n-2; num1Idx++ {
		// Skip duplicates for num1
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}

		num2Idx, num3Idx := num1Idx+1, n-1
		for num2Idx < num3Idx {
			sum := nums[num1Idx] + nums[num2Idx] + nums[num3Idx]
			if sum == 0 {
				// Add triplet to result
				result = append(result, []int{nums[num1Idx], nums[num2Idx], nums[num3Idx]})
				num2Idx++

				// Skip duplicates for num2
				for num2Idx < num3Idx && nums[num2Idx] == nums[num2Idx-1] {
					num2Idx++
				}
			} else if sum > 0 {
				num3Idx--
			} else {
				num2Idx++
			}
		}
	}
	return result
}
