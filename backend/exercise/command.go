package exercise

import (
	"errors"
	"regexp"
	"strconv"
)

var placeholderRegex = regexp.MustCompile(`{(\d+)}`)

type CreateMultipleChoiceExerciseCommand struct {
	Question      string
	Options       []string
	CorrectOption string
}

func NewCreateMultipleChoiceExerciseCommand(
	question string,
	options []string,
	correctOption string,
) CreateMultipleChoiceExerciseCommand {
	return CreateMultipleChoiceExerciseCommand{
		Question:      question,
		Options:       options,
		CorrectOption: correctOption,
	}
}

type CreateCompleteTheSentenceExerciseCommand struct {
	Sentence string // e.g. "This is a {0} truck."
	Blank    string // e.g. "fire"
}

func NewCreateCompleteTheSentenceExerciseCommand(sentence, blank string) (*CreateCompleteTheSentenceExerciseCommand, error) {
	placeholders := placeholderRegex.FindAllStringSubmatch(sentence, -1)
	if len(placeholders) == 0 {
		return nil, errors.New("no placeholder found in sentence")
	}
	if len(placeholders) > 1 {
		return nil, errors.New("more than one placeholder found in sentence")
	}
	if placeholders[0][1] != "0" {
		return nil, errors.New("placeholder {0} not found")
	}

	return &CreateCompleteTheSentenceExerciseCommand{
		Sentence: sentence,
		Blank:    blank,
	}, nil
}

type CreateCompleteTheTextExerciseCommand struct {
	Text   string   // e.g. "This is a family {0}. Hi, how are {1} doing?"
	Blanks []string // e.g. ["member", "you"]
}

func NewCreateCompleteTheTextExerciseCommand(
	text string,
	blanks []string,
) (*CreateCompleteTheTextExerciseCommand, error) {
	placeholders := placeholderRegex.FindAllStringSubmatch(text, -1)
	if len(placeholders) != len(blanks) {
		return nil, errors.New("number of placeholders does not match number of blanks")
	}
	if len(placeholders) == 0 {
		return nil, errors.New("no placeholders found in text")
	}
	if placeholders[0][1] != "0" {
		return nil, errors.New("placeholders are not starting from 0")
	}
	for i := 1; i < len(placeholders); i++ {
		prev, _ := strconv.Atoi(placeholders[i-1][1])
		curr, _ := strconv.Atoi(placeholders[i][1])
		if curr != prev+1 {
			return nil, errors.New("placeholders are not incrementing")
		}
	}

	return &CreateCompleteTheTextExerciseCommand{
		Text:   text,
		Blanks: blanks,
	}, nil
}
