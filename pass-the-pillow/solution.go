package main

import (
	"fmt"
)

func main() {
	fmt.Println(passThePillow(4, 5))
	fmt.Println(passThePillow(3, 2))
}

func passThePillow(n int, time int) int {
    cycleLength := 2 * (n - 1)
    effectiveTime := time % cycleLength

    if effectiveTime < n {
        return effectiveTime + 1
    } else {
        return 2*n - effectiveTime - 1
    }
}