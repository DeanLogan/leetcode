package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumBoxes([]int{1,3,2}, []int{4,3,1,5,2}))
	fmt.Println(minimumBoxes([]int{5,5,5}, []int{2,4,2,7}))
}

func minimumBoxes(apple []int, capacity []int) int {	
	apples := 0
    boxesUsed := 0
    bags := 0
    output := 0

	sort.Sort(sort.Reverse(sort.IntSlice(capacity)))
	
    for _, appleNum := range apple {
        apples += appleNum
    }

    for _, cap := range capacity {
        bags += cap
        boxesUsed++
        if bags >= apples {
            output = boxesUsed
            break
        }
    }

    return output
}
