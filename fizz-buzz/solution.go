package main

import (
	"fmt"
)

func main() {
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
    res := make([]string, n) 
    for i := 1; i <= n; i++ {
        if i%3 == 0 && i%5 == 0 {
            res[i-1] = "FizzBuzz"
        } else if i%3 == 0 {
            res[i-1] = "Fizz"
        } else if i%5 == 0 {
            res[i-1] = "Buzz"
        } else {
            res[i-1] = fmt.Sprintf("%d", i)
        }
    }
    return res
}