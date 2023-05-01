package exercise

import (
	"errors"
	"regexp"
)

var placeholderRegex = regexp.MustCompile(`{(\d+)}`)

type CreateMultipleChoiceExerciseCommand struct {
	Prompt        string
	Options       []string
	CorrectAnswer string
}

func NewCreateMultipleChoiceExerciseCommand(
	prompt string,
	options []string,
	correctAnswer string,
) CreateMultipleChoiceExerciseCommand {
	return CreateMultipleChoiceExerciseCommand{
		Prompt:        prompt,
		Options:       options,
		CorrectAnswer: correctAnswer,
	}
}

type CreateFillInTheBlankExerciseCommand struct {
	Prompt        string // e.g. "This is a {0} truck."
	CorrectAnswer string // e.g. "fire"
}

func NewCreateFillInTheBlankExerciseCommand(prompt, correctAnswer string) (*CreateFillInTheBlankExerciseCommand, error) {
	placeholders := placeholderRegex.FindAllStringSubmatch(prompt, -1)
	if len(placeholders) == 0 {
		return nil, errors.New("no placeholder found in prompt")
	}
	if len(placeholders) > 1 {
		return nil, errors.New("more than one placeholder found in prompt")
	}
	if placeholders[0][1] != "0" {
		return nil, errors.New("placeholder {0} not found in prompt")
	}

	return &CreateFillInTheBlankExerciseCommand{
		Prompt:        prompt,
		CorrectAnswer: correctAnswer,
	}, nil
}
