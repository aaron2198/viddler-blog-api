package viddler

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/bucket"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/config"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/generator"
)

type Viddler struct {
	Config      *config.Config
	BucketStore *bucket.BucketStore
}

func New() (*Viddler, error) {
	godotenv.Load()
	viddler := &Viddler{
		Config: config.New(),
	}
	var err error
	viddler.BucketStore, err = bucket.New(&viddler.Config.BucketStore)
	if err != nil {
		return nil, err
	}
	return viddler, err
}

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
