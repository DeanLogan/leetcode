package main 

import (
    "fmt"
)

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{2,3,4}, 6))
	fmt.Println(twoSum([]int{-1,0}, -1))
}

func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1

    for left < right {
        sumOfPointers := numbers[left] + numbers[right]
        if sumOfPointers == target {
            return []int{left+1, right+1}
        }
        if sumOfPointers > target  {
            right--
        } else {
            left++
        }
    }

    return []int{0,0}
}