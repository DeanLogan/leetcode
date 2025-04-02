package main 

import (
	"fmt"
)

func main() {
	fmt.Println(decrypt([]int{5,7,1,4}, 3))
	fmt.Println(decrypt([]int{1,2,3,4}, 0))
	fmt.Println(decrypt([]int{12,5,6,13}, -2))
}

func decrypt(code []int, k int) []int {
    n := len(code) 
    decryptedCode := make([]int, n) 

    if k == 0 {
        return decryptedCode
    }

    var windowSum int
    for i := 0; i < abs(k); i++ {
        windowSum += code[i]
    }

    for i := abs(k); i < n+abs(k); i++ {
        nextElement := code[i % n]
        prevElement := code[i - abs(k)]
        pos := i

        if k > 0 {
            pos = i - k - 1 
        }

        decryptedCode[(pos + n) % n] = windowSum
        windowSum += nextElement - prevElement
    }

    return decryptedCode
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}