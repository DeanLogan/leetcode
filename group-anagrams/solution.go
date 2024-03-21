package main 

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))
	fmt.Println(groupAnagrams([]string{}))
	fmt.Println(groupAnagrams([]string{"a"}))
}

func groupAnagrams(strs []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, s := range strs {
		sorted := sortString(s)
		anagramGroups[sorted] = append(anagramGroups[sorted], s)
	}

	groupedAnagrams := [][]string{}
	for _, group := range anagramGroups {
		groupedAnagrams = append(groupedAnagrams, group)
	}

	return groupedAnagrams
}

func sortString(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}