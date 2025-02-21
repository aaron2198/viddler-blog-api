package generator

import "fmt"

type OutputType string

const (
	OutputTypeArticle      OutputType = "article"
	OutputTypeRecipe       OutputType = "recipe"
	OutputTypeInstructions OutputType = "instructions"
	OutputTypeSummary      OutputType = "summary"
	OutputTypeQuiz         OutputType = "quiz"
)

type PromptForType map[OutputType]string

var prompts = PromptForType{
	OutputTypeArticle:      "Use the following transcript to create an article, consider including SEO keywords, a call to action and a well thought out conclusion.",
	OutputTypeRecipe:       "Use the following transcript to create a recipe, first include a summary of ingredients, then steps to prepare.",
	OutputTypeInstructions: "Use the following transcript to create a set of instructions, consider including a list of tools and supplies needed, then prvide a list of steps to complete the task.",
	OutputTypeSummary:      "Use the following transcript to create a summary of the video.",
	OutputTypeQuiz:         "Use the following transcript to create a quiz with at least 10 questions, some are multiple choice and others are open ended. Provide a summary on the suggested correct answers.",
}

var PostPrompt = "Format the output with markdown."
var PrePrompt = ""

func PromptFor(outputType OutputType) string {
	return fmt.Sprintf("%s %s %s", PrePrompt, prompts[outputType], PostPrompt)
}

func WithSentence(prompt, sentence string) string {
	return fmt.Sprintf("%s %s", prompt, sentence)
}

func ContentTypes() []OutputType {
	r := []OutputType{}
	for k := range prompts {
		r = append(r, k)
	}
	return r
}
