package aiservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/srt"
)

type OpenAiClient struct {
	Client *openai.Client
	Model  openai.ChatModel
}

type OpenAiCompatibleClient struct {
	*OpenAiClient
}

var openAiModels = map[string]OpenAiModelWithFeatures{
	"GPT-4o Mini":   {Model: openai.ChatModelGPT4oMini, Features: ModelFeatures{StructuredOutputs: true}},
	"GPT-4o":        {Model: openai.ChatModelGPT4o, Features: ModelFeatures{StructuredOutputs: true}},
	"o1-mini":       {Model: openai.ChatModelO1Mini, Features: ModelFeatures{StructuredOutputs: false}},
	"o1":            {Model: openai.ChatModelO1, Features: ModelFeatures{StructuredOutputs: false}},
	"o1-2024-12-17": {Model: openai.ChatModelO1_2024_12_17, Features: ModelFeatures{StructuredOutputs: true}},
	"o3-mini":       {Model: openai.ChatModelO3Mini, Features: ModelFeatures{StructuredOutputs: true}},
}

type OpenAiModelWithFeatures struct {
	Model    openai.ChatModel
	Features ModelFeatures
}

func OpenAiModels(features ...ModelFeature) []string {
	models := []string{}
	for name, model := range openAiModels {
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

type OpenAiSegments struct {
	Segments []OpenAiSegment `json:"segments" jsonschema_description:"The sections of the article"`
}

type OpenAiSegment struct {
	Title string `json:"title" jsonschema_description:"The title of the section"`
	Start int    `json:"start" jsonschema_description:"The start sectionID"`
}

func GenerateSchema[T any]() interface{} {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

var OpenAiSegmentsSchema = GenerateSchema[OpenAiSegments]()

type OpenAiClientParams struct {
	Key   string
	Model string
}

func NewOpenAiClient(params OpenAiClientParams) *OpenAiClient {
	return &OpenAiClient{
		Client: openai.NewClient(option.WithAPIKey(params.Key)),
		Model:  openAiModels[params.Model].Model,
	}
}

func (oa *OpenAiClient) String() string {
	return fmt.Sprintf("OpenAi with model: %s", oa.Model)
}

func (oa *OpenAiClient) BasicGenerate(prompt, dialogue string) (string, error) {
	chat, err := oa.Client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Model: openai.F(oa.Model),
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(fmt.Sprintf("%s\n%s", prompt, dialogue)),
		}),
	})
	if err != nil {
		return "", err
	}
	return chat.Choices[0].Message.Content, nil
}

func (oa *OpenAiClient) ArticleSegmentsPhase(prompt string, srt *srt.SRT) (*SegmentsPhase, error) {

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("article"),
		Description: openai.F("Base article schema"),
		Schema:      openai.F(OpenAiSegmentsSchema),
		Strict:      openai.Bool(true),
	}

	userMessage := fmt.Sprintf("%s\n%s", prompt, srt.String())

	chat, err := oa.Client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(userMessage),
		}),
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
			openai.ResponseFormatJSONSchemaParam{
				Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
				JSONSchema: openai.F(schemaParam),
			},
		),
		Model: openai.F(oa.Model),
	})
	if err != nil {
		return nil, err
	}

	openAiSegments := SegmentsPhase{}
	_ = json.Unmarshal([]byte(chat.Choices[0].Message.Content), &openAiSegments)

	return &openAiSegments, nil
}

func (oa *OpenAiClient) SegmentContentPhase(prompt, dialogue string) (string, error) {
	userMessage := fmt.Sprintf("%s\n%s", prompt, dialogue)
	chat, err := oa.Client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Model: openai.F(oa.Model),
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(userMessage),
		}),
	})
	if err != nil {
		return "", err
	}
	return chat.Choices[0].Message.Content, nil
}

type OpenAiRefinePhase struct {
	Segments []OpenAiRefineSegment `json:"segments" jsonschema_description:"The sections of the article"`
}

type OpenAiRefineSegment struct {
	Title   string `json:"title" jsonschema_description:"The title of the section"`
	Content string `json:"content" jsonschema_description:"Very descriptive and detailed content of the section"`
}

var OpenAiRefineSchema = GenerateSchema[OpenAiRefinePhase]()

func (oa *OpenAiClient) RefinePhase(prompt, article string) (*RefinePhase, error) {
	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("refined_article"),
		Description: openai.F("Refined article schema"),
		Schema:      openai.F(OpenAiRefineSchema),
		Strict:      openai.Bool(true),
	}
	userMessage := fmt.Sprintf("%s\n%s", prompt, article)
	chat, err := oa.Client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Model: openai.F(oa.Model),
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
			openai.ResponseFormatJSONSchemaParam{
				Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
				JSONSchema: openai.F(schemaParam),
			},
		),
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(userMessage),
		}),
	})
	if err != nil {
		return nil, err
	}
	openAiRefine := RefinePhase{}
	_ = json.Unmarshal([]byte(chat.Choices[0].Message.Content), &openAiRefine)
	return &openAiRefine, nil
}

func (oa *OpenAiClient) GenericPrompt(prompt string) (string, error) {
	chat, err := oa.Client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Model: openai.F(oa.Model),
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
	})
	if err != nil {
		return "", err
	}
	return chat.Choices[0].Message.Content, nil
}
