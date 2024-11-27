package main

import (
    "fmt"
)

func main() {
    fmt.Println(maxEqualRowsAfterFlips([][]int{{0,1},{1,1}}))
    fmt.Println(maxEqualRowsAfterFlips([][]int{{0,1},{1,0}}))
    fmt.Println(maxEqualRowsAfterFlips([][]int{{0,0,0},{0,0,1},{1,1,0}}))
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
    patternCount := make(map[string]int)
    maxRows := 0

    for _, row := range matrix {
        pattern := make([]byte, len(row))
        complement := make([]byte, len(row))
        for i, val := range row {
            if val == 0 {
                pattern[i] = '0'
                complement[i] = '1'
            } else {
                pattern[i] = '1'
                complement[i] = '0'
            }
        }
        
        patternStr := string(pattern)
        complementStr := string(complement)
        
        if patternStr > complementStr {
            patternStr = complementStr
        }
        
        patternCount[patternStr]++
        if patternCount[patternStr] > maxRows {
            maxRows = patternCount[patternStr]
        }
    }

    return maxRows
}