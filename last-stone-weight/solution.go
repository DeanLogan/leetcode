package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(lastStoneWeight([]int{2,7,4,1,8,1}))
	fmt.Println(lastStoneWeight([]int{1}))
}

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	numStones := len(stones)
	if numStones > 1 {
		stone1 := stones[numStones-1]
		stone2 := stones[numStones-2]
		if stone1 < stone2 {
			stones[numStones-2] = stone2 - stone1

		} else if stone1 > stone2 {
			stones[numStones-2] = stone1 - stone2
		} else {
			stones[numStones-1] = 0
			stones[numStones-2] = 0
		}
		lastStoneWeight(stones[:len(stones)-1])
	}
	return stones[0]
}