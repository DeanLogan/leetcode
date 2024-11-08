package main

import (
	"fmt"
)

func main() {
	fmt.Println(addToArrayForm([]int{1,2,0,0}, 34))
	fmt.Println(addToArrayForm([]int{2,7,4}, 181))
	fmt.Println(addToArrayForm([]int{2,1,5}, 806))
}

func addToArrayForm(num []int, k int) []int {
    num[len(num)-1] += k
    i := len(num) - 1
    for i >= 0 {
        if num[i] < 10 {
            return num
        }
        if i == 0 {
            num = append([]int{num[i] / 10}, num...)
            num[1] = num[1] % 10
        } else {
            num[i-1] += num[i] / 10
            num[i] = num[i] % 10
            i--
        }
    }
    return num
}