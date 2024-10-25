package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubarrays([]int{1,1,1}, 1))
	fmt.Println(countSubarrays([]int{1,1,2}, 1))
	fmt.Println(countSubarrays([]int{1,2,3}, 2))
}

func countSubarrays(nums []int, k int) int64 {
    countByValue := make(map[int]int)
    totalSubarrays := int64(0)
    for _, num := range nums {
        newCountByValue := make(map[int]int)
        if num & k == k {
            newCountByValue[num] = 1
            for value, count := range countByValue {
                newValue := value & num
                if newValue & k == k {
                    newCountByValue[newValue] += count
                }
            }
            totalSubarrays += int64(newCountByValue[k])
        }
        countByValue = newCountByValue
    }
    return totalSubarrays
}