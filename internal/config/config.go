package config

import (
	"os"

	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/bucket"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/generator"
)

type Config struct {
	Domain              string
	BucketStore         bucket.BucketStoreConfig
	ArticleGenerator    *generator.ArticleGeneratorConfig
	GeneratedImagesPath string
}

func New() *Config {
	return &Config{
		Domain: os.Getenv("DOMAIN"),
		BucketStore: bucket.BucketStoreConfig{
			Type:      os.Getenv("BUCKET_STORE_TYPE"),
			Endpoint:  os.Getenv("BUCKET_STORE_ENDPOINT"),
			AccessKey: os.Getenv("BUCKET_STORE_ACCESS_KEY"),
			SecretKey: os.Getenv("BUCKET_STORE_SECRET_KEY"),
			Bucket:    os.Getenv("BUCKET_STORE_BUCKET"),
		},
		GeneratedImagesPath: orDefaultString("GENERATED_IMAGES_PATH", "generated-images/"),
		ArticleGenerator: &generator.ArticleGeneratorConfig{
			SubtitlesPath:       orDefaultString("SUBTITLES_PATH", "subtitles/"),
			ThumbnailsPath:      orDefaultString("THUMBNAILS_PATH", "thumbnails/"),
			MetadataPath:        orDefaultString("METADATA_PATH", "metadata/"),
			GeneratedImagesPath: orDefaultString("GENERATED_IMAGES_PATH", "generated-images/"),
			OpenAiAPIKey:        os.Getenv("OPENAI_API_KEY"),
			VertexAPIKey:        os.Getenv("VERTEX_API_KEY"),
			VertexProject:       os.Getenv("VERTEX_PROJECT"),
			VertexLocation:      os.Getenv("VERTEX_LOCATION"),
			OllamaAPIKey:        os.Getenv("OLLAMA_API_KEY"),
			OllamaEndpoint:      os.Getenv("OLLAMA_ENDPOINT"),
			AnthropicAPIKey:     os.Getenv("ANTHROPIC_API_KEY"),
		},
	}
}

func orDefaultString(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
