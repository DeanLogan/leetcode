package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestLocal([][]int{{9,9,8,1},{5,6,2,6},{8,2,6,4},{6,2,2,2}}))
	fmt.Println(largestLocal([][]int{{1,1,1,1,1},{1,1,1,1,1},{1,1,2,1,1},{1,1,1,1,1},{1,1,1,1,1}}))
}

func largestLocal(grid [][]int) [][]int {
    n := len(grid)
    for i := 0; i < n-2; i++ {
        for j := 0; j < n-2; j++ {
            maxVal := grid[i][j]
            for k := 0; k < 3; k++ {
                for l := 0; l < 3; l++ {
                    if grid[i+k][j+l] > maxVal {
                        maxVal = grid[i+k][j+l]
                    }
                }
            }
            grid[i][j] = maxVal
        }
    }

    result := make([][]int, n-2)
    for i := 0; i < n-2; i++ {
        result[i] = grid[i][:n-2]
    }

    return result
}