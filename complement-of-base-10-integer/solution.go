package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(bitwiseComplement(5))
	fmt.Println(bitwiseComplement(7))
	fmt.Println(bitwiseComplement(10))
}

func bitwiseComplement(n int) int {
    binaryStr := strconv.FormatInt(int64(n), 2)
	newBinaryStr := ""
	for i := range binaryStr {
		if binaryStr[i] == '1' {
			newBinaryStr += "0"
		} else {
			newBinaryStr += "1"
		}
	}
	result, err := strconv.ParseInt(newBinaryStr, 2, 64)
    if err != nil {
        fmt.Println("Error converting binary string to int:", err)
        return 0
    }
    return int(result) 
}
