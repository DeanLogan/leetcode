package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumWealth([][]int{{1,2,3}, {3,2,1}}))
	fmt.Println("--")
	fmt.Println(maximumWealth([][]int{{1,5}, {7,3}, {3,5}}))
}

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
    for _, person := range accounts {
		wealth := 0
		for _, account := range person {
			wealth += account
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}
	return maxWealth
}