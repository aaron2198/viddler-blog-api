package generator

import (
	"bytes"

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
