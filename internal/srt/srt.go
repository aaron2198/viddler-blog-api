package srt

import (
	"fmt"
	"strconv"
	"strings"
)

type SRT struct {
	Items []Item
}

// This simple string is a distilled format of an SRT with less complex formatting for models to interpret
func (srt *SRT) String() string {
	var builder strings.Builder
	for _, item := range srt.Items {
		builder.WriteString(fmt.Sprintf("SectionID: %d\n%s\n\n", item.Number, strings.Join(item.Text, "\n")))
	}
	return builder.String()
}

// This is the complete SRT file format to spec
func (srt *SRT) AsFile() string {
	var builder strings.Builder
	for _, item := range srt.Items {
		// Section number
		builder.WriteString(fmt.Sprintf("%d\n", item.Number))

		// Timestamp range in format HH:MM:SS,mmm --> HH:MM:SS,mmm
		startHours := int(item.StartSecond) / 3600
		startMinutes := (int(item.StartSecond) % 3600) / 60
		startSeconds := int(item.StartSecond) % 60
		startMillis := int((item.StartSecond - float64(int(item.StartSecond))) * 1000)

		endHours := int(item.EndSecond) / 3600
		endMinutes := (int(item.EndSecond) % 3600) / 60
		endSeconds := int(item.EndSecond) % 60
		endMillis := int((item.EndSecond - float64(int(item.EndSecond))) * 1000)

		builder.WriteString(fmt.Sprintf("%02d:%02d:%02d,%03d --> %02d:%02d:%02d,%03d\n",
			startHours, startMinutes, startSeconds, startMillis,
			endHours, endMinutes, endSeconds, endMillis))

		// Subtitle text
		builder.WriteString(strings.Join(item.Text, "\n") + "\n\n")
	}
	return builder.String()
}

type Item struct {
	Number      int
	StartSecond float64
	EndSecond   float64
	Text        []string
}

func Parse(content string) *SRT {
	srt := &SRT{
		Items: make([]Item, 0),
	}

	lines := strings.Split(content, "\n")
	var currentItem Item
	var textBuffer []string
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			if currentItem.Number != 0 {
				currentItem.Text = textBuffer
				srt.Items = append(srt.Items, currentItem)
				currentItem = Item{}
				textBuffer = nil
			}
			continue
		}

		// Try to parse as number to identify start of new item
		if num, err := strconv.Atoi(line); err == nil {
			if currentItem.Number != 0 {
				currentItem.Text = textBuffer
				srt.Items = append(srt.Items, currentItem)
				textBuffer = nil
			}
			currentItem = Item{Number: num}
			continue
		}

		// Check if line contains timestamp range
		if strings.Contains(line, " --> ") {
			start, end := convertSrtRangeToSeconds(line)
			currentItem.StartSecond = start
			currentItem.EndSecond = end
			textBuffer = make([]string, 0)
			continue
		}

		// If we have a current item and we're past the timestamp,
		// this must be text content
		if currentItem.Number != 0 && textBuffer != nil {
			textBuffer = append(textBuffer, line)
		}
	}

	// Add the final item if it exists
	if currentItem.Number != 0 {
		currentItem.Text = textBuffer
		srt.Items = append(srt.Items, currentItem)
	}

	return srt
}

func (srt *SRT) Cleanse() {
	srt.DeduplicateSubs()
	srt.RemoveEmptySubs()
}

func (srt *SRT) DeduplicateSubs() {
	for i := 0; i < len(srt.Items)-2; i++ { // -2 to allow checking next next item
		currentItem := &srt.Items[i]
		nextItem := &srt.Items[i+1]
		nextNextItem := &srt.Items[i+2]

		// Check each line in current item
		for _, currentLine := range currentItem.Text {
			// Check if any line in next item contains this line
			for _, nextLine := range nextItem.Text {
				if strings.TrimSpace(currentLine) == strings.TrimSpace(nextLine) {
					// Remove the duplicate line from next item
					nextItem.Text = removeString(nextItem.Text, nextLine)
				}
			}
			// Also check next next item
			for _, nextNextLine := range nextNextItem.Text {
				if strings.TrimSpace(currentLine) == strings.TrimSpace(nextNextLine) {
					// Remove the duplicate line from next next item
					nextNextItem.Text = removeString(nextNextItem.Text, nextNextLine)
				}
			}
		}
	}

	// Handle the last pair separately since we can't check next next
	if len(srt.Items) >= 2 {
		lastIdx := len(srt.Items) - 1
		secondLastItem := &srt.Items[lastIdx-1]
		lastItem := &srt.Items[lastIdx]

		for _, secondLastLine := range secondLastItem.Text {
			for _, lastLine := range lastItem.Text {
				if strings.TrimSpace(secondLastLine) == strings.TrimSpace(lastLine) {
					lastItem.Text = removeString(lastItem.Text, lastLine)
				}
			}
		}
	}
}

func (srt *SRT) RemoveEmptySubs() {
	newItems := make([]Item, 0)
	for i := 0; i < len(srt.Items); i++ {
		if len(srt.Items[i].Text) != 0 {
			newItems = append(newItems, srt.Items[i])
			newItems[len(newItems)-1].Number = len(newItems)
		}
	}
	srt.Items = newItems
}

func removeString(slice []string, s string) []string {
	result := make([]string, 0)
	for _, item := range slice {
		if item != s {
			result = append(result, item)
		}
	}
	return result
}

func convertSrtTimestampToSeconds(timestamp string) float64 {
	// Format: HH:MM:SS,mmm
	parts := strings.Split(timestamp, ":")
	if len(parts) != 3 {
		return 0
	}

	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])

	// Split seconds and milliseconds
	secondParts := strings.Split(parts[2], ",")
	if len(secondParts) != 2 {
		return 0
	}

	seconds, _ := strconv.Atoi(secondParts[0])
	milliseconds, _ := strconv.Atoi(secondParts[1])

	totalSeconds := float64(hours*3600 + minutes*60 + seconds)
	totalSeconds += float64(milliseconds) / 1000.0

	return totalSeconds
}

func convertSrtRangeToSeconds(timeRange string) (float64, float64) {
	// Split on arrow
	parts := strings.Split(timeRange, " --> ")
	if len(parts) != 2 {
		return 0, 0
	}

	startSeconds := convertSrtTimestampToSeconds(parts[0])
	endSeconds := convertSrtTimestampToSeconds(parts[1])

	return startSeconds, endSeconds
}

func (srt *SRT) ChunkOfDialogue(start int, end int) string {
	var dialogueBuilder strings.Builder
	for i, item := range srt.Items[start:end] {
		dialogueBuilder.WriteString(strings.Join(item.Text, "\n"))
		if i < len(srt.Items[start:end])-1 {
			dialogueBuilder.WriteString("\n")
		}
	}
	return dialogueBuilder.String()
}

func (srt *SRT) MapTimesToSections(times []float64) []int {
	sections := make([]int, 0)
	//first is always 0
	sections = append(sections, 0)
	remainingTimes := times[1:]
	for _, time := range remainingTimes {
		found := false
		for j, item := range srt.Items {
			if time >= item.StartSecond && time < srt.Items[j+1].StartSecond {
				sections = append(sections, item.Number)
				found = true
				break
			}
		}
		if !found {
			sections = append(sections, -1) // Add sentinel value if no section found
		}
	}
	return sections
}
