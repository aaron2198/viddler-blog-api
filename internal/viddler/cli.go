package viddler

import (
	"fmt"
	"log"

	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/generator"
)

func (viddler *Viddler) Cli(url string) {
	options := generator.UserProvidedOptions{
		Client:   "openai",
		Model:    "GPT-4o",
		VideoUrl: url,
	}

	params := generator.ArticleGeneratorParams{
		Config:      viddler.Config.ArticleGenerator,
		BucketStore: viddler.BucketStore,
		Options:     &options,
	}

	article, err := generator.New(&params).GenerateArticle()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(article)
}
