package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeSubfolders([]string{"/a","/a/b","/c/d","/c/d/e","/c/f"}))
	fmt.Println(removeSubfolders([]string{"/a","/a/b/c","/a/b/d"}))
	fmt.Println(removeSubfolders([]string{"/a/b/c","/a/b/ca","/a/b/d"}))
}

func removeSubfolders(folder []string) []string {
	fmt.Println("---")

	var result []string
	folderMap := map[string]bool{}
	for _, dir := range folder {
		folderMap[dir] = true
	}

	fmt.Println(folderMap)

	for _, dir := range folder {
		checkDir := dir
		var found bool
		for checkDir != "" {
			checkDir = removeAfterLastSlash(checkDir)
			fmt.Println(">"+checkDir+"<")
			if folderMap[checkDir] {
				found = true
				break
			}
		}

		if !found{
			result = append(result, dir)
		}
	}

	fmt.Println("---")
	return result
}

func removeAfterLastSlash(input string) string {
    lastSlashIndex := strings.LastIndex(input, "/")
    if lastSlashIndex != -1 {
        return input[:lastSlashIndex]
    }
    return input
}
