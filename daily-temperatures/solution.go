package main

import (
	"fmt"
)

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}

func dailyTemperatures(temperatures []int) []int {
    result := make([]int, len(temperatures))
    
    for i := len(temperatures) - 1; i >= 0; i-- {
        j := i + 1  
        for j < len(temperatures) && temperatures[j] <= temperatures[i] {
            if result[j] <= 0 {
                break
            }
            j += result[j]
        }
        if j < len(temperatures) && temperatures[j] > temperatures[i] {
            result[i] = j - i
        }
    }
    return result
}