package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{-1,0,3,5,9,12}, 9))
	fmt.Println(search([]int{-1,0,3,5,9,12}, 2))
	fmt.Println(searchNormalApproach([]int{-1,0,3,5,9,12}, 2))
	fmt.Println(searchNormalApproach([]int{-1,0,3,5,9,12}, 2))

}

func searchNormalApproach(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

// In an attempt to make the binary a bit faster for the question I decided to take more of an unusual 
// approach compared to the tradional one of adjuscting index pointers seen below.

// The array is directly sliced instead of adjusting index pointers (left and right)
// An offset variable is used to track the adjustments made to the index during the search process.
// According to some googling and reading some stack overflow pages, by slicing the array and using an offset it 
// makes for better cache utilisation improving memory usage and cpu performance as there is less 
// overhead needed when accessing memory.
func search(nums []int, target int) int {
    var offset, middle, last int
    for {
        middle = len(nums) / 2
        last = len(nums)-1
        switch {
        case target < nums[0]:
            return -1
        case target == nums[0]:
            return offset 
        case target > nums[last]:
            return -1
        case target == nums[last]:
            return offset + last
        case target < nums[middle]:
            nums = nums[:middle]
            continue
        case target > nums[middle]:
            nums = nums[middle:]
            offset += middle
            continue
        case nums[middle] == target:
            return offset + middle
        }
    }
}
