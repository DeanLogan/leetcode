package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(numPrimeArrangements(5))
	fmt.Println(numPrimeArrangements(100))
}

func numPrimeArrangements(n int) int {
    primeCount := 0
    for i := 1; i <= n; i++ {
        if isPrime(i) {
            primeCount++
        }
    }
    nonPrimeCount := n - primeCount
    return (factorial(primeCount) * factorial(nonPrimeCount)) % 1000000007
}

func isPrime(x int) bool {
    if x < 2 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
        if x%i == 0 {
            return false
        }
    }
    return true
}

func factorial(x int) int {
    result := 1
    for i := 1; i <= x; i++ {
        result = (result * i) % (1000000007)
    }
    return result
}