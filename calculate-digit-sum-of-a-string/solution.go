package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(digitSum("11111222223", 3))
	fmt.Println(digitSum("00000000", 3))
}

func digitSum(s string, k int) string {
    for len(s) > k {
        newS := ""
        for i := 0; i < len(s); i += k {
            sum := 0
			group := s[i:min(i+k, len(s))]
            for _, char := range group {
                sum += int(char - '0')
            }
            newS += strconv.Itoa(sum)
        }
        s = newS
    }
    return s
}