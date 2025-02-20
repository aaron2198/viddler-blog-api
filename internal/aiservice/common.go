package aiservice

import (
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/srt"
)

type ModelFeatures map[ModelFeature]bool

type Client interface {
	String() string
	GenericPrompt(prompt string) (string, error)
	BasicGenerate(prompt, dialogue string) (string, error)
	ArticleSegmentsPhase(propmpt string, srt *srt.SRT) (*SegmentsPhase, error)
	SegmentContentPhase(prompt, dialogue string) (string, error)
	RefinePhase(prompt, article string) (*RefinePhase, error)
}

type TemplatePhaseMap map[string]ClientTemplate
type BuiltPhaseMap map[string]Client

type ClientTemplate struct {
	Client     string `json:"client"`
	Model      string `json:"model"`
	UserPrompt string `json:"userPrompt"`
}

var AvailableClients = []string{"google", "openai", "ollama", "anthropic"}

type ModelFeature string

const (
	PsuedoStructuredOutputs ModelFeature = "psuedo_structured_outputs"
	StructuredOutputs       ModelFeature = "structured_outputs"
	DirectVideoInput        ModelFeature = "direct_video_input"
)

func ModelsForClient(client string, features ...ModelFeature) []string {
	switch client {
	case "google":
		return VertexAiModels(features...)
	case "openai":
		return OpenAiModels(features...)
	case "ollama":
		return OllamaModels(features...)
	case "anthropic":
		return AnthropicModels(features...)
	}
	return []string{}
}

func ModelOptions(features ...ModelFeature) map[string][]string {
	modelOptions := make(map[string][]string)
	for _, client := range AvailableClients {
		modelOptions[client] = ModelsForClient(client, features...)
	}
	return modelOptions
}
