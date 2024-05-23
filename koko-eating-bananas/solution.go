package main

import (
	"fmt"
)

func main(){ 
	fmt.Println(minEatingSpeed([]int{3,6,7,11}, 8))
	fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 5))
	fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 6))
}

func minEatingSpeed(piles []int, h int) int {
    maxPile := 0
	for _, pile := range piles {
		if pile > maxPile {
			maxPile = pile
		}
	}

	// Binary search for the minimum eating speed
	left, right := 1, maxPile
	for left < right {
		mid := left + (right-left)/2
		totalHours := 0
		for _, pile := range piles {
			totalHours += (pile + mid - 1) / mid
		}
		if totalHours <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
