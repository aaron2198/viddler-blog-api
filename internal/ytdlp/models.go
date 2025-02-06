package ytdlp

type VideoMetadata struct {
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	FullTitle      string   `json:"fulltitle"`
	DisplayID      string   `json:"display_id"`
	DurationString string   `json:"duration_string"`
	Categories     []string `json:"categories"`
	Tags           []string `json:"tags"`
	Description    string   `json:"description"`

	// Channel and uploader information
	Channel              string `json:"channel"`
	ChannelFollowerCount int    `json:"channel_follower_count"`
	ChannelIsVerified    bool   `json:"channel_is_verified"`
	Uploader             string `json:"uploader"`
	UploaderID           string `json:"uploader_id"`
	UploaderURL          string `json:"uploader_url"`

	// Engagement metrics
	LikeCount int `json:"like_count"`

	// Upload and availability information
	UploadDate   string `json:"upload_date"`
	Availability string `json:"availability"`

	// URL information
	WebpageURLBasename string `json:"webpage_url_basename"`
	WebpageURLDomain   string `json:"webpage_url_domain"`

	// Extractor information
	Extractor    string `json:"extractor"`
	ExtractorKey string `json:"extractor_key"`

	// Live status
	IsLive  bool `json:"is_live"`
	WasLive bool `json:"was_live"`

	// Technical details
	Epoch          int64    `json:"epoch"`
	Format         string   `json:"format"`
	FormatID       string   `json:"format_id"`
	Ext            string   `json:"ext"`
	Protocol       string   `json:"protocol"`
	Language       string   `json:"language"`
	FormatNote     string   `json:"format_note"`
	FilesizeApprox int64    `json:"filesize_approx"`
	TBR            float64  `json:"tbr"`
	Width          int      `json:"width"`
	Height         int      `json:"height"`
	Resolution     string   `json:"resolution"`
	FPS            int      `json:"fps"`
	DynamicRange   string   `json:"dynamic_range"`
	Vcodec         string   `json:"vcodec"`
	VBR            float64  `json:"vbr"`
	AspectRatio    float64  `json:"aspect_ratio"`
	Acodec         string   `json:"acodec"`
	ABR            float64  `json:"abr"`
	ASR            int      `json:"asr"`
	AudioChannels  int      `json:"audio_channels"`
	Type           string   `json:"_type"`
	Formats        []Format `json:"formats"`

	Chapters []Chapter `json:"chapters"`
}

type Chapter struct {
	StartTime float64 `json:"start_time"`
	EndTime   float64 `json:"end_time"`
	Title     string  `json:"title"`
}

type Format struct {
	// Basic format information
	FormatID   string  `json:"format_id"`
	FormatNote string  `json:"format_note"`
	Ext        string  `json:"ext"`
	Protocol   string  `json:"protocol"`
	Acodec     string  `json:"acodec"`
	Vcodec     string  `json:"vcodec"`
	URL        string  `json:"url"`
	Width      int     `json:"width,omitempty"`
	Height     int     `json:"height,omitempty"`
	FPS        float64 `json:"fps,omitempty"`

	// Storyboard specific
	Rows      int        `json:"rows,omitempty"`
	Columns   int        `json:"columns,omitempty"`
	Fragments []Fragment `json:"fragments,omitempty"`

	// Resolution and quality
	Resolution   string  `json:"resolution"`
	AspectRatio  float64 `json:"aspect_ratio"`
	Quality      float64 `json:"quality,omitempty"`
	DynamicRange string  `json:"dynamic_range,omitempty"`

	// Audio specific
	ASR           int     `json:"asr,omitempty"`
	AudioChannels int     `json:"audio_channels,omitempty"`
	AudioExt      string  `json:"audio_ext"`
	ABR           float64 `json:"abr"`

	// Video specific
	VideoExt string  `json:"video_ext"`
	VBR      float64 `json:"vbr"`

	// File information
	Filesize       int64 `json:"filesize,omitempty"`
	FilesizeApprox int64 `json:"filesize_approx,omitempty"`

	// Technical details
	TBR       float64 `json:"tbr,omitempty"`
	Container string  `json:"container,omitempty"`
	HasDRM    bool    `json:"has_drm"`

	// Language
	Language           string `json:"language,omitempty"`
	LanguagePreference int    `json:"language_preference,omitempty"`

	// Preferences and source
	SourcePreference int `json:"source_preference,omitempty"`

	// Headers and options
	HTTPHeaders       HTTPHeaders       `json:"http_headers"`
	DownloaderOptions DownloaderOptions `json:"downloader_options,omitempty"`

	// Format description
	Format string `json:"format"`
}

type Fragment struct {
	URL      string  `json:"url"`
	Duration float64 `json:"duration"`
}

type HTTPHeaders struct {
	UserAgent      string `json:"User-Agent"`
	Accept         string `json:"Accept"`
	AcceptLanguage string `json:"Accept-Language"`
	SecFetchMode   string `json:"Sec-Fetch-Mode"`
}

type DownloaderOptions struct {
	HTTPChunkSize int `json:"http_chunk_size"`
}
