package main 

import (
    "regexp"
    "fmt"
    "strings"
)

// to improve speed in this problem I decided to write my own regex function.
// The original implementation was using the regex package from the standard library and can be seen in the funcion isPalindromeFirstAttempt 

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
    fmt.Println(isPalindromeFirstAttempt("racecar"))
}

func isAlphanumeric(c byte) bool {
    return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    strLen := len(s)
    left, right := 0, strLen-1
    
    for left < right {
        for left < strLen && !isAlphanumeric(s[left]) {
            left++
        }
        for right >= 0 && !isAlphanumeric(s[right]) {
            right--
        }
        if left >= strLen || right < 0 {
            break
        }
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}


func isPalindromeFirstAttempt(s string) bool {
    regex := regexp.MustCompile("[^a-zA-Z0-9]+")
    s = regex.ReplaceAllString(s, "")
    s = strings.ToLower(s)
    strLen := len(s)
    for i:=0; i<strLen/2; i++ {
        if s[i] != s[strLen-i-1] {
            return false
        }
    }
    return true
}

