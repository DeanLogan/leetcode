package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("10101", "1011"))
}

func addBinary(a string, b string) string {
    var result []byte
    carry := 0

    i, j := len(a)-1, len(b)-1

    for i >= 0 || j >= 0 || carry > 0 {
        sum := carry

        if i >= 0 {
            sum += int(a[i] - '0')
            i--
        }

        if j >= 0 {
            sum += int(b[j] - '0')
            j--
        }

        result = append(result, byte(sum%2)+'0')
        carry = sum / 2
    }

    for k, n := 0, len(result); k < n/2; k++ {
        result[k], result[n-1-k] = result[n-1-k], result[k]
    }

    return string(result)
}