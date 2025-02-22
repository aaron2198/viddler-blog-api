package generator

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/aiservice"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/bucket"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/srt"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/ytdlp"
)

type GenerateMode string

const (
	BasicGenerate      GenerateMode = "basic"
	VideoBasedGenerate GenerateMode = "video"
	PhaseBasedGenerate GenerateMode = "phase"
)

type PromptStep string

const (
	BasicPrompt        PromptStep = "basic"
	VideoPrompt        PromptStep = "video"
	PhaseSegmentPrompt PromptStep = "segment"
	PhaseContentPrompt PromptStep = "content"
	PhaseRefinePrompt  PromptStep = "refine"
)

type PromptRequirementsMap map[PromptStep][]aiservice.ModelFeature

var PromptRequirements = PromptRequirementsMap{
	BasicPrompt:        []aiservice.ModelFeature{},
	VideoPrompt:        []aiservice.ModelFeature{aiservice.DirectVideoInput},
	PhaseSegmentPrompt: []aiservice.ModelFeature{aiservice.PsuedoStructuredOutputs, aiservice.StructuredOutputs},
	PhaseContentPrompt: []aiservice.ModelFeature{},
	PhaseRefinePrompt:  []aiservice.ModelFeature{aiservice.PsuedoStructuredOutputs, aiservice.StructuredOutputs},
}

var PhaseOrder = []PromptStep{PhaseSegmentPrompt, PhaseContentPrompt, PhaseRefinePrompt}
var AvailablePhases = map[PromptStep]string{
	PhaseSegmentPrompt: "Break the video into segments that separate ideas.",
	PhaseContentPrompt: "Use an excerpt of the transcript to create article content.",
	PhaseRefinePrompt:  "Refine the article content to make it more coherent and engaging.",
}

type UserProvidedOptions struct {
	Initialized          bool
	VideoUrl             string
	UserStylePrompt      string
	Client               string
	Model                string
	ChaptersAsSections   bool
	Mode                 GenerateMode
	EmbedVideo           bool
	IncludeDescription   bool
	IncludeTags          bool
	OutputType           OutputType
	SEOTerms             []string
	SelectedPhaseOptions aiservice.TemplatePhaseMap
}

// TODO: change name to Generator, generates all types of content
type ArticleGenerator struct {
	Config               *ArticleGeneratorConfig
	BucketStore          *bucket.BucketStore
	Video                *Video
	Phases               aiservice.BuiltPhaseMap
	Progress             chan any
	Complete             chan struct{}
	Options              *UserProvidedOptions
	Result               *ArticleResult
	PhaseBasedGeneration *aiservice.PhaseBasedGeneration
}

type Video struct {
	Id       string
	Url      string
	Metadata ytdlp.VideoMetadata
	SRT      *srt.SRT
}

type ArticleGeneratorConfig struct {
	OpenAiAPIKey        string
	VertexAPIKey        string
	VertexProject       string
	VertexLocation      string
	OllamaAPIKey        string
	AnthropicAPIKey     string
	OllamaEndpoint      string
	SubtitlesPath       string
	ThumbnailsPath      string
	MetadataPath        string
	GeneratedImagesPath string
}

var DefaultPhases = aiservice.TemplatePhaseMap{
	"segments": aiservice.ClientTemplate{Client: "google", Model: "Gemini 2.0 Flash", UserPrompt: "Use this transcript to identify sections for an article, read the entire transcript and determine at most 10 sections to refine the content into. Each section has a 'sectionID: number', use this number to indicate the beginning of a section and sections should be in order. Each section should have an accurate title. Respond with the provided JSON schema. transcript: "},
	"content":  aiservice.ClientTemplate{Client: "google", Model: "Gemini 2.0 Flash", UserPrompt: "Using the provided dialigue, create paragraph summarizing it in the form of an article. This is one piece of a larger article and will be combined with other sections to form a full article so do not mention lack of context from previous sections. Avoid terms like 'this section' or 'this section is about': "},
	"refine":   aiservice.ClientTemplate{Client: "openai", Model: "GPT-4o Mini", UserPrompt: "The users prompt is a article. Please refine the article to be more concise and coherent. Each section should have an accurate title and updated content. If a section is not relavent to the article it can be removed or combined with another section. Respond with the provided JSON schema. article: "},
}

type ArticleGeneratorParams struct {
	Config      *ArticleGeneratorConfig
	BucketStore *bucket.BucketStore
	Options     *UserProvidedOptions
}

func New(params *ArticleGeneratorParams) *ArticleGenerator {
	ag := &ArticleGenerator{
		Config:               params.Config,
		BucketStore:          params.BucketStore,
		Options:              params.Options,
		Progress:             make(chan any, 1),
		Complete:             make(chan struct{}, 1),
		Phases:               make(aiservice.BuiltPhaseMap),
		PhaseBasedGeneration: &aiservice.PhaseBasedGeneration{},
		Result: &ArticleResult{
			Options: params.Options,
		},
	}
	ag.Video = &Video{
		Url: params.Options.VideoUrl,
	}
	ag.WithCustomPhases(DefaultPhases)
	if len(ag.Options.SelectedPhaseOptions) > 0 {
		ag.WithCustomPhases(ag.Options.SelectedPhaseOptions)
	}
	return ag
}

func (ag *ArticleGenerator) WithCustomPhases(template aiservice.TemplatePhaseMap) *ArticleGenerator {
	for _, phase := range PhaseOrder {
		client := template[string(phase)].Client
		model := template[string(phase)].Model
		aiclient := ag.GetClient(client, model)
		if aiclient != nil {
			ag.Phases[string(phase)] = aiclient
		}
	}
	return ag
}

func (ag *ArticleGenerator) GetClient(client, model string) aiservice.Client {
	switch client {
	case "google":
		return aiservice.NewVertexAiClient(aiservice.VertexAiClientParams{
			Key:      ag.Config.VertexAPIKey,
			Project:  ag.Config.VertexProject,
			Location: ag.Config.VertexLocation,
			Model:    model,
		})
	case "openai":
		return aiservice.NewOpenAiClient(aiservice.OpenAiClientParams{
			Key:   ag.Config.OpenAiAPIKey,
			Model: model,
		})
	case "ollama":
		return aiservice.NewOllamaClient(aiservice.OllamaClientParams{
			Key:     ag.Config.OllamaAPIKey,
			BaseURL: ag.Config.OllamaEndpoint,
			Model:   model,
		})
	case "anthropic":
		return aiservice.NewAnthropicClient(aiservice.AnthropicClientParams{
			Key:   ag.Config.AnthropicAPIKey,
			Model: model,
		})
	default:
		return aiservice.NewOpenAiClient(aiservice.OpenAiClientParams{
			Key:   ag.Config.OpenAiAPIKey,
			Model: "GPT-4o Mini",
		})
	}
}

type QueueData struct {
	SEOTerms        []string
	StyleSuggestion string
	Metadata        *ArticleResult
}

func (ag *ArticleGenerator) QueueArticle() (*QueueData, error) {
	if err := ag.Setup(); err != nil {
		return nil, err
	}
	qd := &QueueData{
		Metadata: ag.Result,
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(ag *ArticleGenerator, qd *QueueData) {
		fmt.Println(ag.Options.Client, ag.Options.Model)
		prompt := fmt.Sprintf("Generate a comma separated list of 10 article SEO terms for the following video: titled '%s', with the description '%s', and the following dialogue '%s'.", ag.Result.Title, ag.Result.Description, ag.Video.SRT.String())
		terms, _ := ag.GetClient(ag.Options.Client, ag.Options.Model).GenericPrompt(prompt)
		qd.SEOTerms = GenericListParser(terms)
		wg.Done()
	}(ag, qd)
	go func(ag *ArticleGenerator, qd *QueueData) {
		fmt.Println(ag.Options.Client, ag.Options.Model)
		prompt := fmt.Sprintf("If you were to write and article for this video, what would be a good style to write it in? Respond exlusively with a one sentence description of a suggested style. Title '%s', with the description '%s', and the following dialogue '%s'.", ag.Result.Title, ag.Result.Description, ag.Video.SRT.String())
		style, _ := ag.GetClient(ag.Options.Client, ag.Options.Model).GenericPrompt(prompt)
		qd.StyleSuggestion = style
		wg.Done()
	}(ag, qd)
	wg.Wait()
	return qd, nil
}

func (ag *ArticleGenerator) GenerateArticle() (*ArticleResult, error) {
	if err := ag.Setup(); err != nil {
		ag.Complete <- struct{}{}
		return nil, err
	}

	switch ag.Options.Mode {
	case BasicGenerate:
		content, err := ag.ExecuteBasicGeneration()
		if err != nil {
			ag.Complete <- struct{}{}
			return nil, fmt.Errorf("failed to generate AI content: %w", err)
		}
		ag.Result.Body = content
	case VideoBasedGenerate:
		content, err := ag.ExecuteVideoBasedGeneration()
		if err != nil {
			ag.Complete <- struct{}{}
			return nil, fmt.Errorf("failed to generate AI content: %w", err)
		}
		ag.Result.Body = content
	case PhaseBasedGenerate:
		if err := ag.ExecutePhaseGeneration(); err != nil {
			ag.Complete <- struct{}{}
			return nil, fmt.Errorf("failed to generate AI content: %w", err)
		}
	}
	if err := ag.Result.html(); err != nil {
		ag.Complete <- struct{}{}
		return nil, fmt.Errorf("failed to generate HTML: %w", err)
	}
	ag.Complete <- struct{}{}
	return ag.Result, nil
}

func (ag *ArticleGenerator) StoreTranscript(ctx context.Context) (key string, err error) {
	key = ag.Config.SubtitlesPath + ag.Video.Id + ".en.srt"
	if exists, err := ag.BucketStore.ObjectExists(ctx, key); err != nil || !exists {
		subtitlePath, err := ytdlp.DownloadSubtitles(ag.Video.Id, "/tmp/subtitles")
		if err != nil {
			return "", fmt.Errorf("failed to download subtitles: %s", err)
		}
		err = ag.BucketStore.PutObjectFromFs(ctx, subtitlePath, key)
		if err != nil {
			return "", fmt.Errorf("failed to upload subtitles: %s", err)
		}
	}
	return key, nil
}

func (ag *ArticleGenerator) StoreThumbnail(ctx context.Context) (key string, err error) {
	key = ag.Config.ThumbnailsPath + ag.Video.Id + ".jpg"
	if exists, err := ag.BucketStore.ObjectExists(ctx, key); err != nil || !exists {
		thumbnailPath, err := ytdlp.DownloadThumbnail(ag.Video.Id, "/tmp/thumbnails")
		if err != nil {
			return "", fmt.Errorf("failed to download thumbnail: %s", err)
		}
		err = ag.BucketStore.PutObjectFromFs(ctx, thumbnailPath, key)
		if err != nil {
			return "", fmt.Errorf("failed to upload thumbnail: %s", err)
		}
	}
	return key, nil
}

func (ag *ArticleGenerator) StoreMetadata(ctx context.Context) (key string, err error) {
	key = ag.Config.MetadataPath + ag.Video.Id + ".info.json"
	if exists, err := ag.BucketStore.ObjectExists(ctx, key); err != nil || !exists {
		metadataPath, err := ytdlp.DownloadVideoMetadata(ag.Video.Id, "/tmp/metadata")
		if err != nil {
			return "", fmt.Errorf("failed to download metadata: %s", err)
		}
		err = ag.BucketStore.PutObjectFromFs(ctx, metadataPath, key)
		if err != nil {
			return "", fmt.Errorf("failed to upload metadata: %s", err)
		}
	}
	return key, nil
}

func (ag *ArticleGenerator) Setup() error {
	var err error
	ag.Video.Id, err = ytdlp.ParseVideoID(ag.Options.VideoUrl)
	if err != nil {
		return fmt.Errorf("failed to parse video ID: %s", err)
	}

	ctx := context.Background()
	//Get Transcript
	transcriptKey, err := ag.StoreTranscript(ctx)
	if err != nil {
		return fmt.Errorf("failed to store transcript: %s", err)
	}

	//Get Thumbnail
	thumbnailKey, err := ag.StoreThumbnail(ctx)
	if err != nil {
		return fmt.Errorf("failed to store thumbnail: %s", err)
	}
	ag.Result.Thumbnail, err = ag.BucketStore.PublicUrl(thumbnailKey)
	if err != nil {
		return fmt.Errorf("failed to get public url for thumbnail: %s", err)
	}

	//Get Video Metadata
	metadataKey, err := ag.StoreMetadata(ctx)
	if err != nil {
		return fmt.Errorf("failed to store metadata: %s", err)
	}
	metadataString, err := ag.BucketStore.GetFileAsString(ctx, metadataKey)
	if err != nil {
		return fmt.Errorf("failed to get metadata: %s", err)
	}

	//Store Metadata
	var metadata ytdlp.VideoMetadata
	json.Unmarshal([]byte(metadataString), &metadata)
	ag.Video.Metadata = metadata
	ag.Result.Title = metadata.FullTitle
	ag.Result.VideoId = ag.Video.Id
	ag.Result.VideoUrl = ag.Options.VideoUrl
	ag.Result.Uploader = metadata.Uploader
	ag.Result.UploaderUrl = metadata.UploaderURL
	ag.Result.Description = metadata.Description
	ag.Result.Tags = metadata.Tags
	ag.Result.Categories = metadata.Categories

	//Parse Transcript
	srtContent, err := ag.BucketStore.GetFileAsString(ctx, transcriptKey)
	if err != nil {
		return fmt.Errorf("failed to get subtitles: %s", err)
	}
	ag.Video.SRT = srt.Parse(srtContent)
	//Clean up duplicates and empty entries
	ag.Video.SRT.Cleanse()

	//Set Progress Channel
	ag.ProgressPrinter()
	return nil
}

func (ag *ArticleGenerator) ExecuteBasicGeneration() (string, error) {
	client := ag.GetClient(ag.Options.Client, ag.Options.Model)
	prompt := WithSentence(PromptFor(ag.Options.OutputType), ag.Options.UserStylePrompt)
	if len(ag.Options.SEOTerms) > 0 {
		prompt = WithSentence(prompt, fmt.Sprintf("SEO Terms to include in the article: %s", strings.Join(ag.Options.SEOTerms, ", ")))
	}
	fmt.Println(prompt)
	content, err := client.BasicGenerate(prompt, ag.Video.SRT.String())
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %s", err)
	}
	return MarkdownToHTML(content), nil
}

func (ag *ArticleGenerator) ExecuteVideoBasedGeneration() (string, error) {
	model := "Gemini 2.0 Flash"
	if ag.Options.Model != "" {
		model = ag.Options.Model
	}
	client := aiservice.NewVertexAiClient(aiservice.VertexAiClientParams{
		Key:      "",
		Project:  ag.Config.VertexProject,
		Location: ag.Config.VertexLocation,
		Model:    model,
	})
	if client == nil {
		return "", fmt.Errorf("failed to get client: %s", "Vertex AI")
	}
	prompt := WithSentence(PromptFor(ag.Options.OutputType), ag.Options.UserStylePrompt)
	if len(ag.Options.SEOTerms) > 0 {
		prompt = WithSentence(prompt, fmt.Sprintf("SEO Terms to include in the article: %s", strings.Join(ag.Options.SEOTerms, ", ")))
	}
	resp, err := client.SpecialVideoPrompt(ag.Options.VideoUrl, prompt)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %s", err)
	}
	return MarkdownToHTML(resp), nil
}

func (ag *ArticleGenerator) ExecutePhaseGeneration() error {
	segmentsPhase := &aiservice.SegmentsPhase{}
	var err error
	if ag.Options.ChaptersAsSections && len(ag.Video.Metadata.Chapters) > 0 {
		starts := ytdlp.GetChapterStartTimes(ag.Video.Metadata.Chapters)
		sectionIds := ag.Video.SRT.MapTimesToSections(starts)
		for i, chapter := range ag.Video.Metadata.Chapters {
			segmentsPhase.Segments = append(segmentsPhase.Segments, &aiservice.Segment{
				Start: sectionIds[i],
				Title: chapter.Title,
			})
		}
	} else {
		fmt.Println("Segmenting using: ", ag.Phases["segments"].String())
		segmentsPhase, err = ag.SegmentPhase()
		if err != nil {
			return err
		}
	}
	segmentsPhase.Segments = aiservice.ReorderSegments(segmentsPhase.Segments)
	ag.Progress <- segmentsPhase
	ag.PhaseBasedGeneration.Segments = segmentsPhase

	fmt.Println("Generating content using: ", ag.Phases["content"].String())
	segmentContentPhase, err := ag.SegmentContentPhase(segmentsPhase)
	if err != nil {
		return err
	}
	ag.Progress <- segmentContentPhase
	segmentContentPhase.Segments = aiservice.ReorderSegments(segmentContentPhase.Segments)
	ag.PhaseBasedGeneration.Content = segmentContentPhase

	fmt.Println("Refining content using: ", ag.Phases["refine"].String())
	refinedPhase, err := ag.RefineSegmentsPhase(segmentContentPhase)
	if err != nil {
		return err
	}
	ag.Progress <- refinedPhase
	ag.PhaseBasedGeneration.Refine = refinedPhase

	ag.Result.Sections = []ArticleSection{}
	for _, refinedSegment := range refinedPhase.Segments {
		ag.Result.Sections = append(ag.Result.Sections, ArticleSection(refinedSegment))
	}

	return nil
}

func (ag *ArticleGenerator) SegmentPhase() (*aiservice.SegmentsPhase, error) {
	prompt := ag.Options.SelectedPhaseOptions["segments"].UserPrompt
	segments, err := ag.Phases["segments"].ArticleSegmentsPhase(prompt, ag.Video.SRT)
	if err != nil {
		return nil, err
	}
	return segments, nil
}

func (ag *ArticleGenerator) SegmentContentPhase(segmentsPhase *aiservice.SegmentsPhase) (*aiservice.SegmentContentPhase, error) {

	wg := sync.WaitGroup{}
	wg.Add(len(segmentsPhase.Segments))

	//use start only instead of start and end
	segmentContentPhase := aiservice.SegmentContentPhase{}
	for i := 0; i < len(segmentsPhase.Segments); i++ {
		segment := segmentsPhase.Segments[i]
		next := &aiservice.Segment{}
		if i == len(segmentsPhase.Segments)-1 {
			next = &aiservice.Segment{
				Start: len(ag.Video.SRT.Items),
			}
		} else {
			next = segmentsPhase.Segments[i+1]
		}
		go func(segment, next *aiservice.Segment) {
			dialogue := ag.Video.SRT.ChunkOfDialogue(segment.Start, next.Start-1)
			prompt := ag.Options.SelectedPhaseOptions["content"].UserPrompt
			content, err := ag.Phases["content"].SegmentContentPhase(prompt, dialogue)
			if err != nil {
				fmt.Printf("error generating section %d: %s\n", i, err)
				wg.Done()
				return
			}
			swc := aiservice.SegmentWithContent{
				Segment: segment,
				Content: content,
			}
			segmentContentPhase.Segments = append(segmentContentPhase.Segments, swc)
			ag.Progress <- swc
			wg.Done()
		}(segment, next)
	}

	wg.Wait()
	return &segmentContentPhase, nil
}

func (ag *ArticleGenerator) RefineSegmentsPhase(SegmentsWithContent *aiservice.SegmentContentPhase) (*aiservice.RefinePhase, error) {
	prompt := ag.Options.SelectedPhaseOptions["refine"].UserPrompt
	return ag.Phases["refine"].RefinePhase(prompt, SegmentsWithContent.String())
}

func (ag *ArticleGenerator) ProgressPrinter() {
	go func() {
		for {
			select {
			case status := <-ag.Progress:
				raw, _ := json.MarshalIndent(status, "", "\t")
				fmt.Println(string(raw))
			case <-ag.Complete:
				return
			}
		}
	}()
}
