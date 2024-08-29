package main

import (
	"fmt"
)

func main(){
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
    charCount := [128]int{}
    for i := range magazine {
        charCount[magazine[i]]++
    }
    for i := range ransomNote {
        charCount[ransomNote[i]]--
        if charCount[ransomNote[i]] < 0 {
            return false
        }
    }
    return true
}