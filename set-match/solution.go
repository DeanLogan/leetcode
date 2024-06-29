package main

import (
	"fmt"
)

func main(){
	fmt.Println(findErrorNums([]int{1,2,2,4}))
	fmt.Println(findErrorNums([]int{1,1}))
}

func findErrorNums(nums []int) []int {
    n := len(nums)
    expectedSum := n * (n + 1) / 2
    expectedSquareSum := n * (n + 1) * (2*n + 1) / 6
    actualSum, actualSquareSum := 0, 0

    for _, num := range nums {
        actualSum += num
        actualSquareSum += num * num
    }

    sumDiff := actualSum - expectedSum
    squareSumDiff := actualSquareSum - expectedSquareSum

    missing := (squareSumDiff/sumDiff - sumDiff) / 2
    duplicate := sumDiff + missing

    return []int{duplicate, missing}
}