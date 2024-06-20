package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAcronym([]string{"alice","bob","charlie"}, "abc"))
	fmt.Println(isAcronym([]string{"an","apple","charlie"}, "a"))
	fmt.Println(isAcronym([]string{"never","gonna","give","up","on","you"}, "ngguoy"))
}

func isAcronym(words []string, s string) bool {
	var newAcronym string
	for i := 0; i < len(words); i++ {
		newAcronym += string(words[i][0])
	}
	return newAcronym == s
}