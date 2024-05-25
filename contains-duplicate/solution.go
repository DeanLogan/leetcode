package main 

import (
	"fmt"
)

func main(){
	fmt.Println(containsDuplicate([]int{1,2,3,1}))
    fmt.Println(containsDuplicate([]int{1,2,3,4}))
    fmt.Println(containsDuplicate([]int{1,1,1,3,3,4,3,2,4,2}))
}

func containsDuplicate(nums []int) bool {
    contains := make(map[int]bool)
    for _, num := range nums {
        if !contains[num] {
            contains[num] = true
        } else {
            return true
        }
    }
    return false
}