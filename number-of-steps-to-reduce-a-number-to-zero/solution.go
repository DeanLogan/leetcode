package main

import (
	"fmt"
)

func main(){
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
	fmt.Println(numberOfSteps(123))
}

func numberOfSteps(num int) int {
	steps := 0
    for num != 0 {
		if num % 2 == 0 {
			num = num / 2
		} else {
			num --
		}
		steps++
	}
	return steps
}