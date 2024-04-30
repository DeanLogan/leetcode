package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}

func isValid(s string) bool {
	brackets := map[byte]byte{
        ')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]byte, 0)

	for _, char := range []byte(s) {
		oppositeBracket, inMap := brackets[char]
		if !inMap {
			stack = append(stack, char)
			continue
		}
        stackLen := len(stack)
		if stackLen == 0 {
			return false
		}
		if stack[stackLen - 1] != oppositeBracket {
			return false
		}
		stack = stack[:stackLen - 1]
	}

	return len(stack) == 0
}
