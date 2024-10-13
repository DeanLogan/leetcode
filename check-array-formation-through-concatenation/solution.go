package main

import (
	"fmt"
)

func main() {
	fmt.Println(canFormArray([]int{15,18}, [][]int{{88}, {15}}))
	fmt.Println(canFormArray([]int{49,18,16}, [][]int{{16,18,49}}))
	fmt.Println(canFormArray([]int{91,4,64,78}, [][]int{{78}, {4,64}, {91}}))
}

func canFormArray(arr []int, pieces [][]int) bool {
    for arrIndex := 0; arrIndex < len(arr); {
        pieceIndex := 0
        for pieceIndex < len(pieces) && pieces[pieceIndex][0] != arr[arrIndex] {
            pieceIndex++
        }
        if pieceIndex == len(pieces) {
            return false
        }
        pieceElementIndex := 0
        for pieceElementIndex < len(pieces[pieceIndex]) && arr[arrIndex] == pieces[pieceIndex][pieceElementIndex] {
            arrIndex, pieceElementIndex = arrIndex+1, pieceElementIndex+1
        }
    }
    return true
}