package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculateTax([][]int{{3,50}, {7,10}, {12,25}}, 10))
	fmt.Println(calculateTax([][]int{{1,0}, {4,25}, {5,50}}, 2))
	fmt.Println(calculateTax([][]int{{2,50}}, 0))
}

func calculateTax(brackets [][]int, income int) float64 {
    taxed := 0.0
    previousUpper := 0

    for _, bracket := range brackets {
        upper, percent := bracket[0], bracket[1]
        if income <= previousUpper {
            break
        }
        taxableIncome := min(income, upper) - previousUpper
        taxed += float64(taxableIncome) * float64(percent) / 100.0
        previousUpper = upper
    }

    return taxed
}