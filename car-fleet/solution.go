package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(carFleet(12, []int{10,8,0,5,3}, []int{2,4,1,1,3}))
	fmt.Println(carFleet(10, []int{3}, []int{3}))
	fmt.Println(carFleet(100, []int{0,2,4}, []int{4,2,1}))
}

type car struct {
	pos   int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	stack := []float32{}
	fleets := []car{}
	for i := 0; i < len(position); i++ {
		fleets = append(fleets, car{position[i], speed[i]})
	}

	sort.Slice(fleets, func(i, j int) bool {
		return fleets[i].pos < fleets[j].pos
	})

	for i := len(fleets) -1; i >=0; i-- {
		stack = append(stack, float32(target-fleets[i].pos / fleets[i].speed))
		stackLen := len(stack)
		if stackLen >= 2 && stack[stackLen-1] <= stack[stackLen-2] {
			stack = stack[:stackLen-1]
		}
	}

	return len(stack)
}