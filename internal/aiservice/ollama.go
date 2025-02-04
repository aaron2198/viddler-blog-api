package aiservice

import (
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OllamaClient struct {
	*OpenAiCompatibleClient
}

var ollamaModels = map[string]OllamaModelWithFeatures{
	"Dolphin Llama3": {Model: "dolphin-llama3", Features: ModelFeatures{StructuredOutputs: true}},
	"DeepSeek R1":    {Model: "deepseek-r1-lgctx", Features: ModelFeatures{StructuredOutputs: true}},
}

type OllamaModelWithFeatures struct {
	Model    string
	Features ModelFeatures
}

func OllamaModels(features ...ModelFeature) []string {
	models := []string{}
	for name, model := range ollamaModels {
		if len(features) == 0 {
			models = append(models, name)
		} else {
			for _, feature := range features {
				if model.Features[feature] {
					models = append(models, name)
				}
			}
		}
	}
	return models
}

type OllamaClientParams struct {
	Key     string
	BaseURL string
	Model   string
}

func NewOllamaClient(params OllamaClientParams) *OllamaClient {
	c := openai.NewClient(option.WithAPIKey(params.Key), option.WithBaseURL(params.BaseURL))
	return &OllamaClient{
		OpenAiCompatibleClient: &OpenAiCompatibleClient{
			OpenAiClient: &OpenAiClient{
				Client: c,
				Model:  ollamaModels[params.Model].Model,
			},
		},
	}
}
func (oa *OllamaClient) String() string {
	return fmt.Sprintf("Ollama with model: %s", oa.Model)
}
