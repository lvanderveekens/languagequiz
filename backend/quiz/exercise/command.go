package exercise

import (
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/exp/slices"
)

var blankRegex = regexp.MustCompile(`______`)

type CreateMultipleChoiceExerciseCommand struct {
	Question string
	Choices  []string
	Answer   string
	Feedback *string
}

func NewCreateMultipleChoiceExerciseCommand(
	question string,
	choices []string,
	answer string,
	feedback *string,
) (*CreateMultipleChoiceExerciseCommand, error) {
	if len(choices) != 4 {
		return nil, fmt.Errorf("expected 4 choices, found: %d", len(choices))
	}
	if !slices.Contains(choices, answer) {
		return nil, fmt.Errorf("choices do not contain answer")
	}

	return &CreateMultipleChoiceExerciseCommand{
		Question: question,
		Choices:  choices,
		Answer:   answer,
		Feedback: feedback,
	}, nil
}

type CreateFillInTheBlankExerciseCommand struct {
	Question string // e.g. "This is a ______ truck."
	Answer   string // e.g. "fire"
	Feedback *string
}

func NewCreateFillInTheBlankExerciseCommand(
	question, answer string,
	feedback *string,
) (*CreateFillInTheBlankExerciseCommand, error) {
	blanks := blankRegex.FindAllStringSubmatch(question, -1)
	if len(blanks) == 0 {
		return nil, errors.New("no blank found in question")
	}
	if len(blanks) > 1 {
		return nil, errors.New("more than one blank found in question")
	}

	return &CreateFillInTheBlankExerciseCommand{
		Question: question,
		Answer:   answer,
		Feedback: feedback,
	}, nil
}

type CreateSentenceCorrectionExerciseCommand struct {
	Sentence          string
	CorrectedSentence string
	Feedback          *string
}

func NewCreateSentenceCorrectionExerciseCommand(
	sentence, correctedSentence string,
	feedback *string,
) (*CreateSentenceCorrectionExerciseCommand, error) {
	if sentence == correctedSentence {
		return nil, errors.New("sentence and correctedSentence are the same")
	}

	return &CreateSentenceCorrectionExerciseCommand{
		Sentence:          sentence,
		CorrectedSentence: correctedSentence,
		Feedback:          feedback,
	}, nil
}
