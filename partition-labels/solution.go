package main

import (
	"fmt"
)

func main(){
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
}

func partitionLabels(s string) []int {
    last := make([]int, 26)
    for i, char := range s {
        last[char-'a'] = i
    }

    var partitions []int
    start, end := 0, 0

    for i, char := range s {
        end = max(end, last[char-'a'])
        if i == end {
            partitions = append(partitions, end-start+1)
            start = i + 1
        }
    }

    return partitions
}