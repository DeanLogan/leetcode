package main

import (
	"fmt"
)

func main(){
	fmt.Println(countVowelSubstrings("aeiouu"))
	fmt.Println(countVowelSubstrings("unicornarihan"))
	fmt.Println(countVowelSubstrings("cuaieuouac"))
}

func countVowelSubstrings(word string) int {
    count := 0
    n := len(word)

    for i := 0; i < n; i++ {
        if !isVowel(word[i]) {
            continue
        }
        vowelSet := make(map[byte]bool)
        for j := i; j < n; j++ {
            if !isVowel(word[j]) {
                break
            }
            vowelSet[word[j]] = true
            if len(vowelSet) == 5 {
                count++
            }
        }
    }

    return count
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
