package main

import (
	"fmt"
)

func main() {
	fmt.Println(buddyStrings("ab", "ba"))
	fmt.Println(buddyStrings("ab", "ab"))
	fmt.Println(buddyStrings("aa", "aa"))
}

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == B {
		m := make(map[byte]bool)
		for i := 0; i < len(A); i++ {
			if m[A[i]] {
				return true
			}
			m[A[i]] = true
		}
		return false
	}
	var diff []int
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			diff = append(diff, i)
		}
	}
	return len(diff) == 2 && A[diff[0]] == B[diff[1]] && A[diff[1]] == B[diff[0]]
}