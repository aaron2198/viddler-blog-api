package generator

import (
	"fmt"
	"html/template"
	"io"
)

type ArticleResult struct {
	Options     *UserProvidedOptions
	VideoUrl    string
	VideoId     string
	Uploader    string
	UploaderUrl string
	Description string
	Tags        []string
	Categories  []string
	Title       string
	Thumbnail   string
	Sections    []ArticleSection
	Images      []ArticleImage
	Body        string
}

type ArticleImage struct {
	URL          string
	Caption      string
	SectionIndex int
}

type ArticleSection struct {
	Title   string
	Content string
}

func (ar *ArticleResult) HTML(style string, w io.Writer) error {
	tmpl, err := template.ParseFiles("templates/article.html")
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}
	return tmpl.ExecuteTemplate(w, style, ar)
}

func (ar *ArticleResult) RawBody() template.HTML {
	return template.HTML(ar.Body)
}
