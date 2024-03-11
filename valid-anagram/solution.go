package main 

import (
	"fmt"
	"reflect"
)

func main(){
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}

func isAnagram(s string, t string) bool {
    if reflect.DeepEqual(convertStrToMap(s), convertStrToMap(t)) {
        return true
    }
    return false
}

func convertStrToMap(str string) map[rune]int {
    strMap := make(map[rune]int)
    for _, char := range str {
        strMap[char]++
    }
    return strMap
}