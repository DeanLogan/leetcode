package main 

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}

func reverse(x int) int {
    const intMax = 1<<31 - 1 
    const intMin = -(1 << 31)
    const maxDiv10 = intMax / 10
    const minDiv10 = intMin / 10

    result := 0
    for x != 0 {
        digit := x % 10
        x /= 10

        if result > maxDiv10 || (result == maxDiv10 && digit > 7) {
            return 0
        }
        if result < minDiv10 || (result == minDiv10 && digit < -8) {
            return 0
        }

        result = result*10 + digit
    }
    return result
}