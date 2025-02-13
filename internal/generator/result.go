package generator

import (
	"bytes"
	"fmt"
	"html/template"
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
	HTML        string
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

func (ar *ArticleResult) html() error {
	tmpl, err := template.ParseFiles("templates/article.html")
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}
	style := "basic_article"
	if ar.Options.Mode == PhaseBasedGenerate {
		style = "phase_article"
	}
	buf := bytes.Buffer{}

	err = tmpl.ExecuteTemplate(&buf, style, ar)
	if err != nil {
		return fmt.Errorf("failed to generate article html: %w", err)
	}
	ar.HTML = buf.String()
	return nil
}

func (ar *ArticleResult) RawBody() template.HTML {
	return template.HTML(ar.Body)
}
