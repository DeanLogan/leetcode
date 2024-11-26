package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(maximum69Number(9669))
	fmt.Println(maximum69Number(9996))
	fmt.Println(maximum69Number(9999))
}

func maximum69Number(num int) int {
    s := strings.Split(strconv.Itoa(num), "")
    for i := 0; i < len(s); i++ {
        if s[i] == "6" {
            s[i] = "9"
            break
        }
    }
    ans, _ := strconv.Atoi(strings.Join(s, ""))
    return ans
}