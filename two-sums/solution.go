package main;

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
}

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        containsValue := target - num
        if index, exists := numMap[containsValue]; exists {
            return []int{index, i}
        }
        numMap[num] = i
    }
    return []int{}
}