package main

import (
	"fmt"
	"slices"
	"cmp"
)

func main() {
	fmt.Println(carFleet(12, []int{10,8,0,5,3}, []int{2,4,1,1,3}))
	fmt.Println(carFleet(10, []int{3}, []int{3}))
	fmt.Println(carFleet(100, []int{0,2,4}, []int{4,2,1}))
}

type Car struct {
    position int
    time     float32
}

func carFleet(target int, position []int, speed []int) int {
    cars := make([]Car, len(position))
    for i := range position {
        cars[i] = Car{position[i], float32(target-position[i]) / float32(speed[i])}
    }
    slices.SortFunc(cars, func(a, b Car) int {
        return cmp.Compare(a.position, b.position)
    })
    fleetStack := []Car{}
    for i := len(cars) - 1; i >= 0; i-- {
        if len(fleetStack) == 0 {
            fleetStack = append(fleetStack, cars[i])
            continue
        }
        if cars[i].time > fleetStack[len(fleetStack)-1].time {
            fleetStack = append(fleetStack, cars[i])
        }
    }

    return len(fleetStack)
}
