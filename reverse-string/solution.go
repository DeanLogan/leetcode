package main

import (
	"fmt"
)

func main(){
	eg1 := []byte{'h','e','l','l','o'}
	reverseString(eg1)
	fmt.Println(eg1)
	eg2 := []byte{'h','a','n','n','a','h'}
	reverseString(eg2)
	fmt.Println(eg2)
}

func reverseString(s []byte) {
    left, right := 0, len(s)-1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
}