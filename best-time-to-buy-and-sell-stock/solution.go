package main

import (
	"fmt"
)

func main(){
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}

func maxProfit(prices []int) int {
    min := 100000
    profit := 0
    for _, price := range prices {
        if price > min {
            tempProfit := price - min
            if tempProfit > profit {
                profit = tempProfit
            }
        } else {
            min = price
        }
    }
    return profit
}
