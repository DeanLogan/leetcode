package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(numberOfEmployeesWhoMetTarget([]int{0,1,2,3,4}, 2))
	fmt.Println(numberOfEmployeesWhoMetTarget([]int{5,1,4,2,2}, 6))
}

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    sort.Ints(hours)

    hoursLen := len(hours)
    firstInd := 0
    lastInd := hoursLen - 1

	// binary search to "find" the target hours, if not found it is safe to assume anything after firstInd is larger than the target
    for firstInd <= lastInd {
        midInd := (firstInd + lastInd) / 2
        if hours[midInd] < target {
            firstInd = midInd + 1
        } else {
            lastInd = midInd - 1
        }
    }

    return hoursLen - firstInd
}