package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 3))
	fmt.Println(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 13))
}

func searchMatrix(matrix [][]int, target int) bool {
    for i := 0; i<len(matrix); i++ {
        if matrix[i][0] <= target && target <= matrix[i][len(matrix[i])-1] {
            return binarySearch(matrix[i], target) 
        }
    }
    return false
}

func binarySearch(nums []int, target int) bool {
    var middle, last int
    for {
        middle = len(nums) / 2
        last = len(nums)-1
        switch {
        case target < nums[0]:
            return false
        case target == nums[0]:
            return true 
        case target > nums[last]:
            return false
        case target == nums[last]:
            return true
        case target < nums[middle]:
            nums = nums[:middle]
            continue
        case target > nums[middle]:
            nums = nums[middle:]
            continue
        case nums[middle] == target:
            return true
        }
    }
}
