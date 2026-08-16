package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDuplicate([]int{1,3,4,2,2}))
	fmt.Println(findDuplicate([]int{3,1,3,4,2}))
	fmt.Println(findDuplicate([]int{3,3,3,3,3}))
}

func findDuplicate(nums []int) int {
    slowPointer, fastPointer := 0, 0
	for {
		slowPointer = nums[slowPointer]
		fastPointer = nums[nums[fastPointer]]
		if slowPointer == fastPointer {
			break
		}
	}
	slowPointer2 := 0
    for slowPointer2 != slowPointer {
        slowPointer2 = nums[slowPointer2]
        slowPointer = nums[slowPointer]
    }
	return slowPointer
}