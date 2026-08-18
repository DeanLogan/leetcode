package main

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

func renderReadme(q *question) string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "# %s. %s\n\n", q.QuestionID, q.Title)
	builder.WriteString(renderMarkdownFromHTML(q.Content))
	builder.WriteString("\n\n## Submission Screenshot\n\n")
	fmt.Fprintf(&builder, "![Image](./%s.png)\n", q.TitleSlug)

	return builder.String()
}

func renderSolutionStub(title string) string {
	return fmt.Sprintf(`package main

func main() {
	// TODO: implement %s
}
`, title)
}

func renderMarkdownFromHTML(src string) string {
	content := html.UnescapeString(src)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\u00a0", " ")

	content = regexp.MustCompile(`(?is)<pre[^>]*>(.*?)</pre>`).ReplaceAllStringFunc(content, func(match string) string {
		groups := regexp.MustCompile(`(?is)<pre[^>]*>(.*?)</pre>`).FindStringSubmatch(match)
		if len(groups) < 2 {
			return match
		}
		inner := stripHTML(groups[1])
		inner = strings.Trim(inner, "\n\r\t ")
		if inner == "" {
			return ""
		}
		return "\n```\n" + inner + "\n```\n\n"
	})

	content = regexp.MustCompile(`(?is)<ol[^>]*>(.*?)</ol>`).ReplaceAllStringFunc(content, func(match string) string {
		groups := regexp.MustCompile(`(?is)<ol[^>]*>(.*?)</ol>`).FindStringSubmatch(match)
		if len(groups) < 2 {
			return match
		}
		return renderListBlock(groups[1], true)
	})

	content = regexp.MustCompile(`(?is)<ul[^>]*>(.*?)</ul>`).ReplaceAllStringFunc(content, func(match string) string {
		groups := regexp.MustCompile(`(?is)<ul[^>]*>(.*?)</ul>`).FindStringSubmatch(match)
		if len(groups) < 2 {
			return match
		}
		return renderListBlock(groups[1], false)
	})

	content = regexp.MustCompile(`(?is)<img[^>]*src="([^"]+)"[^>]*>`).ReplaceAllString(content, "![Image]($1)")
	content = regexp.MustCompile(`(?is)<a[^>]*href="([^"]+)"[^>]*>(.*?)</a>`).ReplaceAllString(content, "[$2]($1)")
	content = regexp.MustCompile(`(?is)<strong[^>]*>(.*?)</strong>`).ReplaceAllString(content, "**$1**")
	content = regexp.MustCompile(`(?is)<em[^>]*>(.*?)</em>`).ReplaceAllString(content, "*$1*")
	content = regexp.MustCompile(`(?is)<code[^>]*>(.*?)</code>`).ReplaceAllString(content, "`$1`")
	content = regexp.MustCompile(`(?is)<br\s*/?>`).ReplaceAllString(content, "\n")
	content = regexp.MustCompile(`(?is)</p>`).ReplaceAllString(content, "\n\n")
	content = regexp.MustCompile(`(?is)<p[^>]*>`).ReplaceAllString(content, "")
	content = regexp.MustCompile(`(?is)</?div[^>]*>`).ReplaceAllString(content, "")
	content = regexp.MustCompile(`(?is)</?span[^>]*>`).ReplaceAllString(content, "")
	content = regexp.MustCompile(`(?is)</?sup[^>]*>`).ReplaceAllString(content, "")
	content = regexp.MustCompile(`(?is)</?sub[^>]*>`).ReplaceAllString(content, "")
	content = regexp.MustCompile(`(?is)</?tbody[^>]*>`).ReplaceAllString(content, "")
	content = regexp.MustCompile(`(?is)</?thead[^>]*>`).ReplaceAllString(content, "")
	content = regexp.MustCompile(`(?is)</?tr[^>]*>`).ReplaceAllString(content, "")
	content = regexp.MustCompile(`(?is)</?td[^>]*>`).ReplaceAllString(content, "")
	content = regexp.MustCompile(`(?is)</?th[^>]*>`).ReplaceAllString(content, "")
	content = regexp.MustCompile(`(?is)</?table[^>]*>`).ReplaceAllString(content, "")
	content = stripHTML(content)

	return normalizeMarkdown(content)
}

func normalizeMarkdown(content string) string {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\u00a0", " ")
	content = regexp.MustCompile(`\n{3,}`).ReplaceAllString(content, "\n")
	lines := strings.Split(content, "\n")
	exampleRe := regexp.MustCompile(`^\*\*Example \d+:\*\*$`)
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t")
		if exampleRe.MatchString(lines[i]) {
			lines[i] += "  "
		}
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func stripHTML(content string) string {
	content = regexp.MustCompile(`(?is)<li[^>]*>`).ReplaceAllString(content, "- ")
	content = regexp.MustCompile(`(?is)</li>`).ReplaceAllString(content, "\n")
	content = regexp.MustCompile(`(?is)</?ol[^>]*>`).ReplaceAllString(content, "\n")
	content = regexp.MustCompile(`(?is)</?ul[^>]*>`).ReplaceAllString(content, "\n")
	content = regexp.MustCompile(`(?is)</?[a-zA-Z!][^>]*>`).ReplaceAllString(content, "")
	return content
}

func renderListBlock(inner string, ordered bool) string {
	liRegex := regexp.MustCompile(`(?is)<li[^>]*>(.*?)</li>`)
	matches := liRegex.FindAllStringSubmatch(inner, -1)
	if len(matches) == 0 {
		return ""
	}

	var builder strings.Builder
	for index, match := range matches {
		item := ""
		if len(match) > 1 {
			item = renderInlineMarkdown(match[1])
		}
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}

		prefix := "- "
		if ordered {
			prefix = fmt.Sprintf("%d. ", index+1)
		}
		builder.WriteString(prefix + item + "\n")
	}

	if builder.Len() == 0 {
		return ""
	}

	builder.WriteString("\n")
	return builder.String()
}

func renderInlineMarkdown(content string) string {
	content = html.UnescapeString(content)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\u00a0", " ")

	content = regexp.MustCompile(`(?is)<a[^>]*href="([^"]+)"[^>]*>(.*?)</a>`).ReplaceAllString(content, "[$2]($1)")
	content = regexp.MustCompile(`(?is)<strong[^>]*>(.*?)</strong>`).ReplaceAllString(content, "**$1**")
	content = regexp.MustCompile(`(?is)<em[^>]*>(.*?)</em>`).ReplaceAllString(content, "*$1*")
	content = regexp.MustCompile(`(?is)<code[^>]*>(.*?)</code>`).ReplaceAllString(content, "`$1`")
	content = regexp.MustCompile(`(?is)<br\s*/?>`).ReplaceAllString(content, "\n")
	content = stripHTML(content)
	content = strings.Join(strings.Fields(content), " ")

	return strings.TrimSpace(content)
}
