package exercise

import (
	"errors"
	"regexp"
	"strconv"
)

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
	BeforeGap string
	Gap       string
	AfterGap  string
}

func NewCreateCompleteTheSentenceExerciseCommand(
	beforeGap, gap, afterGap string,
) CreateCompleteTheSentenceExerciseCommand {
	return CreateCompleteTheSentenceExerciseCommand{
		BeforeGap: beforeGap,
		Gap:       gap,
		AfterGap:  afterGap,
	}
}

type CreateCompleteTheTextExerciseCommand struct {
	Text   string   // "This is a family {0}. Hi, how are {1} doing?"
	Blanks []string // ["member", "you"]
}

func NewCreateCompleteTheTextExerciseCommand(
	text string,
	blanks []string,
) (*CreateCompleteTheTextExerciseCommand, error) {
	re := regexp.MustCompile(`{(\d+)}`)
	placeholders := re.FindAllStringSubmatch(text, -1)
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
