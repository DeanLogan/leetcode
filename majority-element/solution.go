package main

import (
	"fmt"
)

func main(){
	fmt.Println(majorityElement([]int{3,2,3}))
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}

func majorityElement(nums []int) int {
    var candidate, count int
    for _, num := range nums {
        if count == 0 { 
            candidate = num
        }
        if num == candidate {
            count++
        } else {
            count--
        }
    }
    return candidate
}