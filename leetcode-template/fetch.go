package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func parseSlug(raw string) (string, error) {
	if !strings.Contains(raw, "://") {
		return sanitizeSlug(raw), nil
	}

	parsed, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	parts := strings.Split(strings.Trim(parsed.Path, "/"), "/")
	for i, part := range parts {
		if part == "problems" && i+1 < len(parts) {
			return sanitizeSlug(parts[i+1]), nil
		}
	}

	return "", fmt.Errorf("could not find /problems/<slug>/ in %q", raw)
}

func sanitizeSlug(slug string) string {
	slug = strings.TrimSpace(strings.ToLower(slug))
	slug = strings.Trim(slug, "/")
	return slug
}

func fetchQuestion(slug string) (*question, error) {
	query := `query questionData($titleSlug: String!) {
		question(titleSlug: $titleSlug) {
			questionId
			title
			titleSlug
			content
			difficulty
			exampleTestcases
			hints
			topicTags {
				name
				slug
			}
		}
	}`

	payload := map[string]any{
		"query": query,
		"variables": map[string]string{
			"titleSlug": slug,
		},
		"operationName": "questionData",
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout:   20 * time.Second,
		Transport: &http.Transport{Proxy: nil},
	}
	request, err := http.NewRequest(http.MethodPost, "https://leetcode.com/graphql", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	request.Header.Set("Referer", fmt.Sprintf("https://leetcode.com/problems/%s/", slug))

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("leetcode responded with %s: %s", response.Status, strings.TrimSpace(string(body)))
	}

	var decoded questionResponse
	if err := json.Unmarshal(body, &decoded); err != nil {
		return nil, err
	}
	if len(decoded.Errors) > 0 {
		return nil, errors.New(decoded.Errors[0].Message)
	}
	if decoded.Data.Question == nil {
		return nil, errors.New("question not found")
	}

	return decoded.Data.Question, nil
}
