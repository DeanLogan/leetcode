package main

import (
	"fmt"
)

func main(){
	fmt.Println(rowAndMaximumOnes([][]int{{0,1},{1,0}}))
	fmt.Println(rowAndMaximumOnes([][]int{{0,0,0},{0,1,1}}))
	fmt.Println(rowAndMaximumOnes([][]int{{0,0},{1,1},{0,0}}))
}

func rowAndMaximumOnes(mat [][]int) []int {
    result := make([]int, 2)
	
	for i, binaryArr := range mat {
		oneCounter := 0
		for _, i := range binaryArr {
			if i == 1 {
				oneCounter++
			}
		}
		if oneCounter > result[1] {
			result = []int{i, oneCounter}
		}
	}

	return result
}