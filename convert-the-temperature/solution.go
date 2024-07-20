package main

import (
	"fmt"
)

func main(){
	fmt.Println(convertTemperature(36.50))
	fmt.Println(convertTemperature(122.11))
}

func convertTemperature(celsius float64) []float64 {
    return []float64{celsius + 273.15, (celsius * 1.8) + 32.00}
}