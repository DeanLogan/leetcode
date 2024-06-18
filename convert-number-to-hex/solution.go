package main

import (
	"fmt"
)

func main(){
	fmt.Println(toHex(26))
	fmt.Println(toHex(-1))
}

func toHex(num int) string {
    if num == 0 {
        return "0"
    }
    if num < 0 {
        num += 1 << 32
    }
    hexChars := "0123456789abcdef"
    res := ""
    for num > 0 {
        res = string(hexChars[num%16]) + res
        num /= 16
    }
    return res
}