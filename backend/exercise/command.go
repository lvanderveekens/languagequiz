package exercise

import (
	"errors"
	"regexp"
)

var placeholderRegex = regexp.MustCompile(`{(\d+)}`)

type CreateMultipleChoiceExerciseCommand struct {
	Question string
	Options  []string
	Answer   string
}

func NewCreateMultipleChoiceExerciseCommand(
	question string,
	options []string,
	answer string,
) CreateMultipleChoiceExerciseCommand {
	return CreateMultipleChoiceExerciseCommand{
		Question: question,
		Options:  options,
		Answer:   answer,
	}
}

type CreateFillInTheBlankExerciseCommand struct {
	Question string // e.g. "This is a {0} truck."
	Answer   string // e.g. "fire"
}

func NewCreateFillInTheBlankExerciseCommand(question, answer string) (*CreateFillInTheBlankExerciseCommand, error) {
	placeholders := placeholderRegex.FindAllStringSubmatch(question, -1)
	if len(placeholders) == 0 {
		return nil, errors.New("no placeholder found in question")
	}
	if len(placeholders) > 1 {
		return nil, errors.New("more than one placeholder found in question")
	}
	if placeholders[0][1] != "0" {
		return nil, errors.New("placeholder {0} not found in question")
	}

	return &CreateFillInTheBlankExerciseCommand{
		Question: question,
		Answer:   answer,
	}, nil
}
