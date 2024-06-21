package main

import (
	"fmt"
)

func main(){ 
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 3))
	fmt.Println(search([]int{1}, 0))
}

func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    
    for left <= right {
        mid := (left + right) / 2
        if target == nums[mid] {
            return mid
        }
        
        if nums[left] <= nums[mid] {
            if target > nums[mid] || target < nums[left] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        } else {
            if target < nums[mid] || target > nums[right] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        }
    }
    return -1
}