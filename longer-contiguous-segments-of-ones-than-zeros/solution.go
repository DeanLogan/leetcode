package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkZeroOnes("1101"))
	fmt.Println(checkZeroOnes("111000"))
	fmt.Println(checkZeroOnes("110100010"))
}

func checkZeroOnes(s string) bool {
    zeroCount := 0
    oneCount := 0
    currentCount := 1

    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            currentCount++
        } else {
            if s[i-1] == '0' {
                if currentCount > zeroCount {
                    zeroCount = currentCount
                }
            } else {
                if currentCount > oneCount {
                    oneCount = currentCount
                }
            }
            currentCount = 1
        }
    }

    if s[len(s)-1] == '0' {
        if currentCount > zeroCount {
            zeroCount = currentCount
        }
    } else {
        if currentCount > oneCount {
            oneCount = currentCount
        }
    }

    return oneCount > zeroCount
}