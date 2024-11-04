package main

import (
	"fmt"
)

func main() {
	fmt.Println(divisorSubstrings(240, 2))
	fmt.Println(divisorSubstrings(430043, 2))
}

func divisorSubstrings(num int, k int) int {
    count := 0
    divisor := 1
    for i := 0; i < k; i++ {
        divisor *= 10
    }

    for n := num; n >= divisor/10; n /= 10 {
        subInt := n % divisor
        if subInt != 0 && num%subInt == 0 {
            count++
        }
    }

    return count
}