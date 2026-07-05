package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	urlFlag := flag.String("url", "", "LeetCode problem URL")
	rootFlag := flag.String("root", "", "Repository root override")
	flag.Parse()

	input := strings.TrimSpace(*urlFlag)
	if input == "" && flag.NArg() > 0 {
		input = strings.TrimSpace(flag.Arg(0))
	}
	if input == "" {
		fmt.Fprintln(os.Stderr, "usage: go run ./tools/leetcode-template -url https://leetcode.com/problems/two-sum/")
		os.Exit(1)
	}

	slug, err := parseSlug(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse url:", err)
		os.Exit(1)
	}

	questionData, err := fetchQuestion(slug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "fetch leetcode data:", err)
		os.Exit(1)
	}

	root := strings.TrimSpace(*rootFlag)
	if root == "" {
		root, err = findRepoRoot()
		if err != nil {
			fmt.Fprintln(os.Stderr, "locate repo root:", err)
			os.Exit(1)
		}
	}

	folderPath := filepath.Join(root, questionData.TitleSlug)
	if _, err := os.Stat(folderPath); err == nil {
		fmt.Fprintln(os.Stderr, "target folder already exists:", folderPath)
		os.Exit(1)
	} else if !errors.Is(err, os.ErrNotExist) {
		fmt.Fprintln(os.Stderr, "check target folder:", err)
		os.Exit(1)
	}

	if err := os.MkdirAll(folderPath, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, "create folder:", err)
		os.Exit(1)
	}

	if err := writeFile(filepath.Join(folderPath, "solution.go"), renderSolutionStub(questionData.Title)); err != nil {
		fmt.Fprintln(os.Stderr, "write solution.go:", err)
		os.Exit(1)
	}

	if err := writeFile(filepath.Join(folderPath, "README.md"), renderReadme(questionData)); err != nil {
		fmt.Fprintln(os.Stderr, "write README.md:", err)
		os.Exit(1)
	}

	fmt.Println("created:", folderPath)
}
