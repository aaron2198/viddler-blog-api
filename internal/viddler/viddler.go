package viddler

import (
	"context"

	"github.com/joho/godotenv"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/aiservice"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/bucket"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/config"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/db"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/generator"
)

type Viddler struct {
	Config      *config.Config
	BucketStore *bucket.BucketStore
	DB          *db.DB
}

func New() (*Viddler, error) {
	godotenv.Load()
	viddler := &Viddler{
		Config: config.New(),
	}
	var err error
	viddler.BucketStore, err = bucket.New(viddler.Config.BucketStore)
	if err != nil {
		return nil, err
	}
	viddler.DB, err = db.Init(context.Background(), viddler.Config.DB)
	if err != nil {
		return nil, err
	}
	return viddler, nil
}

func (v *Viddler) StoreArticle(ctx context.Context, options *generator.UserProvidedOptions, article *generator.ArticleResult) (int, error) {
	tx, err := v.DB.Client.Tx(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	dbArticle, err := tx.Article.Create().
		SetTitle(article.Title).
		SetThumbnail(article.Thumbnail).
		SetVideoURL(options.VideoUrl).
		SetVideoID(article.VideoId).
		SetDescription(article.Description).
		SetUploader(article.Uploader).
		SetUploaderURL(article.UploaderUrl).
		SetHTML(article.HTML).Save(ctx)
	if err != nil {
		return 0, err
	}
	dbUserOptions, err := tx.UserOptions.Create().
		SetVideoURL(options.VideoUrl).
		SetUserStylePrompt(options.UserStylePrompt).
		SetClient(options.Client).
		SetModel(options.Model).
		SetMode(string(options.Mode)).
		SetChaptersAsSections(options.ChaptersAsSections).
		SetEmbedVideo(options.EmbedVideo).
		SetIncludeDescription(options.IncludeDescription).
		SetIncludeTags(options.IncludeTags).
		SetArticle(dbArticle).
		Save(ctx)
	if err != nil {
		return dbArticle.ID, err
	}

	if options.Mode == generator.PhaseBasedGenerate {
		for name, opts := range options.SelectedPhaseOptions {
			_, err := tx.PhaseOptions.Create().
				SetPhaseName(name).
				SetClient(opts.Client).
				SetModel(opts.Model).
				SetPrompt(opts.UserPrompt).
				SetUserOptions(dbUserOptions).
				Save(ctx)
			if err != nil {
				return dbArticle.ID, err
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return dbArticle.ID, err
	}
	return dbArticle.ID, nil
}

func (v *Viddler) GetArticle(ctx context.Context, id int) (*generator.ArticleResult, error) {
	dbArticle, err := v.DB.Client.Article.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	userOptions, err := dbArticle.QueryUserOptions().Only(ctx)
	if err != nil {
		return nil, err
	}
	article := &generator.ArticleResult{
		VideoUrl:    dbArticle.VideoURL,
		VideoId:     dbArticle.VideoID,
		Uploader:    dbArticle.Uploader,
		UploaderUrl: dbArticle.UploaderURL,
		Description: dbArticle.Description,
		Title:       dbArticle.Title,
		Thumbnail:   dbArticle.Thumbnail,
		Sections:    []generator.ArticleSection{},
		Images:      []generator.ArticleImage{},
		HTML:        dbArticle.HTML,
		Options: &generator.UserProvidedOptions{
			VideoUrl:           userOptions.VideoURL,
			UserStylePrompt:    userOptions.UserStylePrompt,
			Client:             userOptions.Client,
			Model:              userOptions.Model,
			ChaptersAsSections: userOptions.ChaptersAsSections,
			EmbedVideo:         userOptions.EmbedVideo,
			IncludeDescription: userOptions.IncludeDescription,
			IncludeTags:        userOptions.IncludeTags,
			Mode:               generator.GenerateMode(userOptions.Mode),
		},
	}
	phaseOptions, err := userOptions.QueryPhaseOptions().All(ctx)
	if err != nil {
		return nil, err
	}
	article.Options.SelectedPhaseOptions = make(aiservice.TemplatePhaseMap)
	for _, phaseOption := range phaseOptions {
		article.Options.SelectedPhaseOptions[phaseOption.PhaseName] =
			aiservice.ClientTemplate{
				Client:     phaseOption.Client,
				Model:      phaseOption.Model,
				UserPrompt: phaseOption.Prompt,
			}
	}
	return article, nil
}
