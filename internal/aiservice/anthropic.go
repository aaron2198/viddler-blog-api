package aiservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/srt"
)

type AnthropicClient struct {
	Client *anthropic.Client
	Model  string
}

var anthropicModels = map[string]AnthropicModelWithFeatures{
	"Claude 3.5 Sonnet 2024-10-22": {Model: anthropic.ModelClaude3_5Sonnet20241022, Features: ModelFeatures{PsuedoStructuredOutputs: true}},
	"Claude 3.5 Haiku":             {Model: anthropic.ModelClaude3_5HaikuLatest, Features: ModelFeatures{PsuedoStructuredOutputs: true}},
}

type AnthropicModelWithFeatures struct {
	Model    anthropic.Model
	Features ModelFeatures
}

type AnthropicClientParams struct {
	Key   string
	Model anthropic.Model
}

func NewAnthropicClient(params AnthropicClientParams) *AnthropicClient {
	return &AnthropicClient{
		Client: anthropic.NewClient(option.WithAPIKey(params.Key)),
		Model:  anthropicModels[params.Model].Model,
	}
}

func (ac *AnthropicClient) String() string {
	return fmt.Sprintf("Anthropic with model: %s", ac.Model)
}

func (ac *AnthropicClient) BasicGenerate(prompt, dialogue string) (string, error) {
	message, err := ac.Client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(ac.Model),
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(fmt.Sprintf("%s\n%s", prompt, dialogue))),
		}),
	})
	if err != nil {
		return "", err
	}
	return message.Content[0].Text, nil
}

func AnthropicModels(features ...ModelFeature) []string {
	models := []string{}
	for name, model := range anthropicModels {
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

var psuedoSegmentsSchema = "output in JSON format with the key 'segments' as a list of objects with the keys 'title' and 'start'."

func (ac *AnthropicClient) ArticleSegmentsPhase(prompt string, srt *srt.SRT) (*SegmentsPhase, error) {
	userPrompt := fmt.Sprintf("%s %s\n%s", psuedoSegmentsSchema, prompt, srt.String())
	message, err := ac.Client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(ac.Model),
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(userPrompt)),
		}),
	})
	if err != nil {
		return nil, err
	}
	segments := &SegmentsPhase{}
	err = json.Unmarshal([]byte(message.Content[0].Text), segments)
	if err != nil {
		return nil, err
	}
	return segments, nil
}

func (ac *AnthropicClient) SegmentContentPhase(prompt, dialogue string) (string, error) {
	userPrompt := fmt.Sprintf("%s\n%s", prompt, dialogue)
	message, err := ac.Client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(ac.Model),
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(userPrompt)),
		}),
	})
	if err != nil {
		return "", err
	}
	return message.Content[0].Text, nil
}

var psuedoRefineSchema = "output in JSON format with the key 'segments' as a list of objects with the keys 'title' and 'content'."

func (ac *AnthropicClient) RefinePhase(prompt, article string) (*RefinePhase, error) {
	userPrompt := fmt.Sprintf("%s %s\n%s", psuedoRefineSchema, prompt, article)
	message, err := ac.Client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(ac.Model),
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(userPrompt)),
		}),
	})
	if err != nil {
		return nil, err
	}
	refine := &RefinePhase{}
	err = json.Unmarshal([]byte(message.Content[0].Text), refine)
	if err != nil {
		return nil, err
	}
	return refine, nil
}

func (ac *AnthropicClient) GenericPrompt(prompt string) (string, error) {
	message, err := ac.Client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(ac.Model),
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		}),
	})
	if err != nil {
		return "", err
	}
	return message.Content[0].Text, nil
}
