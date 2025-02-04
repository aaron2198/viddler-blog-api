package ytdlp

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ParseVideoID(videoURL string) (string, error) {
	videoID := ""
	// Handle different URL formats
	switch {
	case len(videoURL) == 11:
		// Direct video ID
		videoID = videoURL
	case len(videoURL) > 11:
		// Full URL formats:
		// https://www.youtube.com/watch?v=VIDEO_ID
		// https://youtu.be/VIDEO_ID
		// https://youtube.com/shorts/VIDEO_ID
		if idx := strings.Index(videoURL, "watch?v="); idx != -1 {
			videoID = videoURL[idx+8:]
			if ampIdx := strings.Index(videoID, "&"); ampIdx != -1 {
				videoID = videoID[:ampIdx]
			}
		} else if idx := strings.Index(videoURL, "youtu.be/"); idx != -1 {
			videoID = videoURL[idx+9:]
			if slashIdx := strings.Index(videoID, "/"); slashIdx != -1 {
				videoID = videoID[:slashIdx]
			}
		} else if idx := strings.Index(videoURL, "shorts/"); idx != -1 {
			videoID = videoURL[idx+7:]
			if slashIdx := strings.Index(videoID, "/"); slashIdx != -1 {
				videoID = videoID[:slashIdx]
			}
			if qIdx := strings.Index(videoID, "?"); qIdx != -1 {
				videoID = videoID[:qIdx]
			}
		}
	}

	if videoID == "" {
		return "", fmt.Errorf("could not parse video ID from URL")
	}
	return videoID, nil
}

func DownloadSubtitles(videoID string, outputDir string) (string, error) {
	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create output directory: %w", err)
	}

	// Construct the output template
	outputTemplate := filepath.Join(outputDir, videoID+".%(ext)s")

	// Prepare yt-dlp command with arguments
	cmd := exec.Command("yt-dlp",
		"--skip-download",  // Don't download the video
		"--write-sub",      // Write subtitle file
		"--write-auto-sub", // Write automatically generated subtitle file
		"--sub-lang", "en", // Download English subtitles
		"--convert-subs", "srt", // Convert subtitles to SRT format
		"-o", outputTemplate, // Output template
		"https://www.youtube.com/watch?v="+videoID,
	)

	// Execute the command
	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("failed to download subtitles: %w\nOutput: %s", err, output)
	}

	// Return the expected subtitle file path
	return filepath.Join(outputDir, videoID+".en.srt"), nil
}

func DownloadThumbnail(videoID string, outputDir string) (string, error) {
	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create output directory: %w", err)
	}

	// Prepare yt-dlp command with arguments
	cmd := exec.Command("yt-dlp",
		"--skip-download",             // Don't download the video
		"--write-thumbnail",           // Write thumbnail file
		"--convert-thumbnails", "jpg", // Convert thumbnail to jpg
		"-o", filepath.Join(outputDir, videoID), // Output template
		"https://www.youtube.com/watch?v="+videoID,
	)

	// Execute the command
	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("failed to download thumbnail: %w\nOutput: %s", err, output)
	}

	return filepath.Join(outputDir, videoID+".jpg"), nil
}

func DownloadVideoMetadata(videoID string, outputDir string) (string, error) {
	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create output directory: %w", err)
	}

	outputPath := filepath.Join(outputDir, videoID+".info.json")

	// Prepare yt-dlp command with arguments
	cmd := exec.Command("yt-dlp",
		"--skip-download",                       // Don't download the video
		"--write-info-json",                     // Write video metadata to JSON file
		"-o", filepath.Join(outputDir, videoID), // Output template
		"https://www.youtube.com/watch?v="+videoID,
	)

	// Execute the command
	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("failed to download metadata: %w\nOutput: %s", err, output)
	}

	return outputPath, nil
}

func GetChapterStartTimes(chapters []Chapter) []float64 {
	flattened := make([]float64, 0)
	for _, chapter := range chapters {
		flattened = append(flattened, chapter.StartTime)
	}
	return flattened
}
