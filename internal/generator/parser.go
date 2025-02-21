package generator

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func MarkdownToHTML(input string) string {
	// Create markdown parser with common extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)

	// Parse markdown into AST
	ast := p.Parse([]byte(input))

	// Create HTML renderer with common flags
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	// Render to HTML
	var buf bytes.Buffer
	buf.Write(markdown.Render(ast, renderer))

	return buf.String()
}

func GenericListParser(input string) []string {
	// Remove any leading/trailing whitespace
	input = strings.TrimSpace(input)

	// If empty return empty slice
	if input == "" {
		return []string{}
	}

	var items []string

	// Check if comma separated
	if strings.Contains(input, ",") {
		// Split by comma and clean each item
		items = strings.Split(input, ",")
		for i, item := range items {
			items[i] = strings.TrimSpace(item)
		}
		return items
	}

	// Check if newline separated
	if strings.Contains(input, "\n") {
		// Split by newline and clean each item
		items = strings.Split(input, "\n")
		for i, item := range items {
			items[i] = strings.TrimSpace(item)
		}
		return items
	}

	// Check if numbered list (e.g. "1. Item")
	if matched, _ := regexp.MatchString(`^\d+\.`, input); matched {
		// Split by newline and remove numbers
		lines := strings.Split(input, "\n")
		for _, line := range lines {
			// Remove number prefix and clean
			item := regexp.MustCompile(`^\d+\.\s*`).ReplaceAllString(line, "")
			item = strings.TrimSpace(item)
			if item != "" {
				items = append(items, item)
			}
		}
		return items
	}

	// Check if numbered list without newlines (e.g. "1. Item 2. Item")
	if matched, _ := regexp.MatchString(`\d+\.`, input); matched {
		// Split by number prefix
		parts := regexp.MustCompile(`\d+\.\s*`).Split(input, -1)
		for _, part := range parts {
			item := strings.TrimSpace(part)
			if item != "" {
				items = append(items, item)
			}
		}
		return items
	}

	// Default to space separated
	items = strings.Fields(input)
	return items
}
