package main

import (
	"fmt"
)

func main(){
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
    strInt := fmt.Sprintf("%d", x)
	for i := 0; i < len(strInt)/2; i++ {
		if strInt[i] != strInt[len(strInt)-i-1] {
			return false
		}
	}
	return true
}