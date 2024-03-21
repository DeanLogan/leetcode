package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(topKFrequent([]int{1,1,1,2,2,3}, 2))
	fmt.Println(topKFrequent([]int{}, 1))
}

func topKFrequent(nums []int, k int) []int {
    frequency := make(map[int]int)
    for _, num := range nums {
        frequency[num]++
    }

    mapToSlice := [][]int{}
    for key, value := range frequency {
        mapToSlice = append(mapToSlice, []int{key,value})
    }

    sort.Slice(mapToSlice, func(i, j int) bool {
        return mapToSlice[i][1] > mapToSlice[j][1]
    })

    mostFrequentNums := make([]int, k)

    for i := range mostFrequentNums {
        mostFrequentNums[i] = mapToSlice[i][0]
    }

    return mostFrequentNums
}