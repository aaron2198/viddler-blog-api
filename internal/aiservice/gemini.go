package aiservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/googleapis/gax-go/v2/apierror"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/srt"
	"google.golang.org/api/option"
)

type GeminiClient struct {
	Client *genai.Client
	Model  *genai.GenerativeModel
}

var geminiModels = map[string]GeminiModelWithFeatures{
	"Gemini 2.0 Flash": {Model: "gemini-2.0-flash-exp", Features: ModelFeatures{StructuredOutputs: true}},
	"Gemini 1.5 Pro":   {Model: "gemini-1.5-pro-exp", Features: ModelFeatures{StructuredOutputs: true}},
	"Gemini 1.5 Flash": {Model: "gemini-1.5-flash-exp", Features: ModelFeatures{StructuredOutputs: true}},
}

type GeminiModelWithFeatures struct {
	Model    string
	Features ModelFeatures
}

func GeminiModels(features ...ModelFeature) []string {
	models := []string{}
	for name, model := range geminiModels {
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

type GeminiClientParams struct {
	Key   string
	Model string
}

func NewGeminiClient(params GeminiClientParams) *GeminiClient {
	gc, _ := genai.NewClient(
		context.Background(),
		option.WithAPIKey(params.Key),
	)
	return &GeminiClient{
		Client: gc,
		Model:  gc.GenerativeModel(geminiModels[params.Model].Model),
	}
}

func (gc *GeminiClient) String() string {
	info, _ := gc.Model.Info(context.Background())
	return fmt.Sprintf("Gemini with model: %s", info.DisplayName)
}

func (gc *GeminiClient) BasicGenerate(prompt, dialogue string) (string, error) {
	resp, err := gc.Model.GenerateContent(context.Background(), genai.Text(fmt.Sprintf("%s\n%s", prompt, dialogue)))
	if err != nil {
		return "", err
	}
	return string(resp.Candidates[0].Content.Parts[0].(genai.Text)), nil
}

func (gc *GeminiClient) WithModelOption(opt func(*genai.GenerativeModel) *genai.GenerativeModel) *GeminiClient {
	gc.Model = opt(gc.Model)
	return gc
}

func WithSegmentsSchema(model *genai.GenerativeModel) *genai.GenerativeModel {
	model.ResponseSchema = geminiSegmentsSchema
	model.ResponseMIMEType = "application/json"
	return model
}

func WithRefineSchema(model *genai.GenerativeModel) *genai.GenerativeModel {
	model.ResponseSchema = geminiRefinedSegmentsSchema
	model.ResponseMIMEType = "application/json"
	return model
}

func (gc *GeminiClient) WithModelSchema(schema *genai.Schema) *GeminiClient {
	gc.Model.ResponseSchema = schema
	return gc
}

var geminiSegmentsSchema = &genai.Schema{
	Type: genai.TypeArray,
	Items: &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"title": {Type: genai.TypeString, Description: "The title of the section"},
			"start": {Type: genai.TypeInteger, Description: "The start srt sectionID from the transcript"},
		},
	},
}

func (gc *GeminiClient) ArticleSegmentsPhase(prompt string, srt *srt.SRT) (*SegmentsPhase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	gc.WithModelOption(WithSegmentsSchema)
	resp, err := gc.throttleGenerateContent(ctx, genai.Text(fmt.Sprintf("%s\n%s", prompt, srt.String())))
	if err != nil {
		return nil, err
	}
	segments := []*Segment{}
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			err := json.Unmarshal([]byte(txt), &segments)
			if err != nil {
				return nil, err
			}
		}
	}
	return &SegmentsPhase{Segments: segments}, nil
}

func (gc *GeminiClient) SegmentContentPhase(prompt, dialogue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	resp, err := gc.throttleGenerateContent(ctx, genai.Text(fmt.Sprintf("%s\n%s", prompt, dialogue)))
	if err != nil {
		return "", err
	}
	content := ""
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			content = string(txt)
		}
	}
	return content, nil
}

var geminiRefinedSegmentsSchema = &genai.Schema{
	Type: genai.TypeArray,
	Items: &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"title":   {Type: genai.TypeString, Description: "The title of the section"},
			"content": {Type: genai.TypeString, Description: "The article content for this section"},
		},
	},
}

func (gc *GeminiClient) RefinePhase(prompt, article string) (*RefinePhase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	gc.WithModelOption(WithRefineSchema)
	resp, err := gc.throttleGenerateContent(ctx, genai.Text(fmt.Sprintf("%s\n%s", prompt, article)))
	if err != nil {
		return nil, err
	}
	segments := []RefinedSegment{}
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			err = json.Unmarshal([]byte(txt), &segments)
			if err != nil {
				return nil, err
			}
		}
	}
	return &RefinePhase{Segments: segments}, nil
}

func (gc *GeminiClient) throttleGenerateContent(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
	if t, ok := ctx.Deadline(); ok {
		if time.Until(t)*time.Millisecond < 0 {
			return nil, fmt.Errorf("context deadline exceeded")
		}
	} else {
		return nil, fmt.Errorf("context deadline not set")
	}
	resp, err := gc.Model.GenerateContent(ctx, parts...)
	if err != nil {
		var apierr *apierror.APIError
		if ok := errors.As(err, &apierr); ok {
			if apierr.HTTPCode() == 429 {
				fmt.Println("rate limit exceeded, sleeping")
				time.Sleep(500 * time.Millisecond)
				return gc.throttleGenerateContent(ctx, parts...)
			}
		}
		return nil, err
	}
	return resp, nil
}
