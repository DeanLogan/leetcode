package main

import (
	"fmt"
)

func main(){
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(checkIfPangram("leetcode"))
}

func checkIfPangram(sentence string) bool {
    alphabetArr := [26]bool{}

    for _, char := range sentence {
        alphabetArr[char - 'a'] = true
    }

    for _, found := range alphabetArr {
        if !found {
            return false
        }
    }

    return true
}