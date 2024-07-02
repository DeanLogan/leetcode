package main

import (
	"fmt"
)

func main(){
	fmt.Println(intersection([][]int{{3,1,2,4,5},{1,2,3,4},{3,4,5,6}}))
	fmt.Println(intersection([][]int{{1,2,3},{4,5,6}}))
}

func intersection(nums [][]int) []int {
    numOfNums := make([]int, 1001)
	for _, numArr := range nums {
		for _, num := range numArr {
			numOfNums[num]++
		}
	}
	dupNums := []int{}
	length := len(nums)
	for num, count := range numOfNums {
		if count == length {
			dupNums = append(dupNums, num)
		}
	}
	return dupNums
}