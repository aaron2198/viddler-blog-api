package aiservice

import (
	"fmt"
	"sort"
	"strings"
)

type SegmentsPhase struct {
	Segments []*Segment
}

func (s *Segment) StartNumber() int {
	return s.Start
}

type HasStartNumber interface {
	StartNumber() int
}

func ReorderSegments[T HasStartNumber](segments []T) []T {
	sort.Slice(segments, func(i, j int) bool {
		return segments[i].StartNumber() < segments[j].StartNumber()
	})
	return segments
}

type Segment struct {
	Title string `json:"title"`
	Start int    `json:"start"`
}

type SegmentContentPhase struct {
	Segments []SegmentWithContent
}

func (scp SegmentContentPhase) String() string {
	var buf strings.Builder
	for _, segment := range scp.Segments {
		buf.WriteString(fmt.Sprintf("## %s\n\n", segment.Title))
		buf.WriteString(segment.Content)
		buf.WriteString("\n\n")
	}
	return buf.String()
}

type SegmentWithContent struct {
	*Segment
	Content string
}

type RefinePhase struct {
	Segments []RefinedSegment
}

type RefinedSegment struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PhaseBasedGeneration struct {
	Segments *SegmentsPhase
	Content  *SegmentContentPhase
	Refine   *RefinePhase
}
