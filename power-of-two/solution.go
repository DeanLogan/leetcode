package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(3))
}

func isPowerOfTwo(n int) bool {
    if n % 2 == 0 && n > 0 {
		return isPowerOfTwo(n/2)
	}
	return n == 1
}