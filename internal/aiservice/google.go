package aiservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/vertexai/genai"
	"github.com/googleapis/gax-go/v2/apierror"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/srt"
)

var vertexAiModels = map[string]VertexAiModelWithFeatures{
	"Gemini 2.0 Flash":     {Model: "gemini-2.0-flash-exp", Features: ModelFeatures{StructuredOutputs: true, DirectVideoInput: true}},
	"Gemini 1.5 Pro":       {Model: "gemini-1.5-pro-exp", Features: ModelFeatures{StructuredOutputs: true, DirectVideoInput: true}},
	"Gemini 1.5 Flash":     {Model: "gemini-1.5-flash-exp", Features: ModelFeatures{StructuredOutputs: true, DirectVideoInput: true}},
	"Gemini 2.0 Pro 02-05": {Model: "gemini-2.0-pro-exp-02-05", Features: ModelFeatures{StructuredOutputs: true, DirectVideoInput: true}},
}

type VertexAiModelWithFeatures struct {
	Model    string
	Features ModelFeatures
}

func VertexAiModels(features ...ModelFeature) []string {
	models := []string{}
	for name, model := range vertexAiModels {
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

type VertexAiClientParams struct {
	Key      string
	Project  string
	Location string
	Model    string
}

type VertexAiClient struct {
	Client *genai.Client
	Model  *genai.GenerativeModel
}

func NewVertexAiClient(params VertexAiClientParams) *VertexAiClient {
	client, err := genai.NewClient(context.Background(), params.Project, params.Location)
	if err != nil {
		fmt.Println("failed to create vertex ai client", err)
		return nil
	}
	return &VertexAiClient{
		Client: client,
		Model:  client.GenerativeModel(vertexAiModels[params.Model].Model),
	}
}

func (vc *VertexAiClient) String() string {
	return fmt.Sprintf("Vertex AI with model: %s", vc.Model.Name())
}

func (vc *VertexAiClient) BasicGenerate(prompt, dialogue string) (string, error) {
	resp, err := vc.Model.GenerateContent(context.Background(), genai.Text(fmt.Sprintf("%s\n%s", prompt, dialogue)))
	if err != nil {
		return "", err
	}
	return string(resp.Candidates[0].Content.Parts[0].(genai.Text)), nil
}

func (vc *VertexAiClient) WithModelOption(opt func(*genai.GenerativeModel) *genai.GenerativeModel) *VertexAiClient {
	vc.Model = opt(vc.Model)
	return vc
}

func WithSegmentsSchema(model *genai.GenerativeModel) *genai.GenerativeModel {
	model.ResponseSchema = vertexAiSegmentsSchema
	model.ResponseMIMEType = "application/json"
	return model
}

func WithRefineSchema(model *genai.GenerativeModel) *genai.GenerativeModel {
	model.ResponseSchema = vertexAiRefinedSegmentsSchema
	model.ResponseMIMEType = "application/json"
	return model
}

func (vc *VertexAiClient) WithModelSchema(schema *genai.Schema) *VertexAiClient {
	vc.Model.ResponseSchema = schema
	return vc
}

var vertexAiSegmentsSchema = &genai.Schema{
	Type: genai.TypeArray,
	Items: &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"title": {Type: genai.TypeString, Description: "The title of the section"},
			"start": {Type: genai.TypeInteger, Description: "The start srt sectionID from the transcript"},
		},
	},
}

func (vc *VertexAiClient) ArticleSegmentsPhase(prompt string, srt *srt.SRT) (*SegmentsPhase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	vc.WithModelOption(WithSegmentsSchema)
	resp, err := vc.throttleGenerateContent(ctx, genai.Text(fmt.Sprintf("%s\n%s", prompt, srt.String())))
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

func (vc *VertexAiClient) SegmentContentPhase(prompt, dialogue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	resp, err := vc.throttleGenerateContent(ctx, genai.Text(fmt.Sprintf("%s\n%s", prompt, dialogue)))
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

var vertexAiRefinedSegmentsSchema = &genai.Schema{
	Type: genai.TypeArray,
	Items: &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"title":   {Type: genai.TypeString, Description: "The title of the section"},
			"content": {Type: genai.TypeString, Description: "The article content for this section"},
		},
	},
}

func (vc *VertexAiClient) RefinePhase(prompt, article string) (*RefinePhase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	vc.WithModelOption(WithRefineSchema)
	resp, err := vc.throttleGenerateContent(ctx, genai.Text(fmt.Sprintf("%s\n%s", prompt, article)))
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

func (vc *VertexAiClient) throttleGenerateContent(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
	if t, ok := ctx.Deadline(); ok {
		if time.Until(t)*time.Millisecond < 0 {
			return nil, fmt.Errorf("context deadline exceeded")
		}
	} else {
		return nil, fmt.Errorf("context deadline not set")
	}
	resp, err := vc.Model.GenerateContent(ctx, parts...)
	if err != nil {
		var apierr *apierror.APIError
		if ok := errors.As(err, &apierr); ok {
			if apierr.HTTPCode() == 429 {
				fmt.Println("rate limit exceeded, sleeping")
				time.Sleep(500 * time.Millisecond)
				return vc.throttleGenerateContent(ctx, parts...)
			}
		}
		return nil, err
	}
	return resp, nil
}

var basePrompt = `
    You are a top-tier content marketing and SEO expert. 
    You will watch or analyze the video at {video_url} to gather the main points.

    Your task:
    1. Turn the video content into a well-structured, SEO-optimized blog post.
    2. Include a catchy title, an engaging introduction, subheadings with relevant keywords,
       bullet-point tips, a clear conclusion, and a compelling call to action.
    3. Provide a recommended meta title and meta description, along with a list of relevant target keywords.
    4. The main focus of the blog post is: '{topic_description}'.

    Please format your response as follows:

    - Title of the Blog
    - Introduction (2-3 short paragraphs)
    - Main Sections (with subheadings that incorporate your chosen keywords)
      * Explanation or bullet-point tips under each subheading
    - Conclusion
    - Call to Action
    - Recommended Meta Title
    - Recommended Meta Description
    - Relevant Target Keywords
`

func (vc *VertexAiClient) SpecialVideoPrompt(url, prompt string) (string, error) {
	if prompt == "" {
		prompt = basePrompt
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	resp, err := vc.throttleGenerateContent(ctx, genai.Text(prompt), genai.FileData{
		FileURI:  url,
		MIMEType: "video/webm",
	})
	if err != nil {
		return "", err
	}
	return string(resp.Candidates[0].Content.Parts[0].(genai.Text)), nil
}

func (vc *VertexAiClient) GenericPrompt(prompt string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()
	resp, err := vc.throttleGenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}
	return string(resp.Candidates[0].Content.Parts[0].(genai.Text)), nil
}
