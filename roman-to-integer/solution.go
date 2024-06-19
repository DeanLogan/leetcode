package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
    symbolMap := map[byte]int {
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i:=len(s)-1; i>=0; i-- {
		if s[i] == 'I' || i==0 {
			result += symbolMap[s[i]]
		} else {
			if symbolMap[s[i-1]] < symbolMap[s[i]] {
				result += (symbolMap[s[i]] - symbolMap[s[i-1]])
				i--
			} else {
				result += symbolMap[s[i]]
			}
		}
	}
	return result
}