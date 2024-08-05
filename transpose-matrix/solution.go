package main

import (
	"fmt"
)

func main(){
	fmt.Println(transpose([][]int{{1,2,3},{4,5,6},{7,8,9}}))
	fmt.Println(transpose([][]int{{1,2,3},{4,5,6}}))
}

func transpose(matrix [][]int) [][]int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return [][]int{}
    }
    
    rows := len(matrix)
    cols := len(matrix[0])
    result := make([][]int, cols)
    
    for i := range result {
        result[i] = make([]int, rows)
    }
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            result[j][i] = matrix[i][j]
        }
    }
    
    return result
}