package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
	fmt.Println(sortPeople([]string{"Alice","Bob","Bob"}, []int{155,185,150}))
}

type Person struct {
	name string
	height int
}

func sortPeople(names []string, heights []int) []string {
    namesAndHeights := make([]Person, len(names))
    for i := range names {
        namesAndHeights[i] = Person{names[i], heights[i]}
    }

    quickSort(namesAndHeights, 0, len(namesAndHeights)-1)

    for i, person := range namesAndHeights {
        names[i] = person.name
    }

    return names
}

func quickSort(people []Person, low, high int) {
    if low < high {
        pi := partition(people, low, high)
        quickSort(people, low, pi-1)
        quickSort(people, pi+1, high)
    }
}

func partition(people []Person, low, high int) int {
    pivot := people[high].height
    i := low - 1
    for j := low; j < high; j++ {
        if people[j].height > pivot {
            i++
            people[i], people[j] = people[j], people[i]
        }
    }
    people[i+1], people[high] = people[high], people[i+1]
    return i + 1
}