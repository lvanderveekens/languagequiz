package exercise

import (
	"time"
)

type Exercise interface {
	CheckAnswer(answer any) bool
	GetAnswer() any
	GetType() string
}

type exerciseBase struct {
	ID        string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newExerciseBase(id, exerciseType string, createdAt, updatedAt time.Time) exerciseBase {
	return exerciseBase{
		ID:        id,
		Type:      exerciseType,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type MultipleChoiceExercise struct {
	exerciseBase
	Question string
	Choices  []string
	Answer   string
}

func NewMultipleChoiceExercise(
	id string,
	createdAt, updatedAt time.Time,
	question string,
	choices []string,
	answer string,
) MultipleChoiceExercise {
	return MultipleChoiceExercise{
		exerciseBase: newExerciseBase(id, TypeMultipleChoice, createdAt, updatedAt),
		Question:     question,
		Choices:      choices,
		Answer:       answer,
	}
}

func (e *MultipleChoiceExercise) CheckAnswer(answer any) bool {
	if answer, ok := answer.(string); ok {
		return e.Answer == answer
	}
	return false
}

func (e *MultipleChoiceExercise) GetAnswer() any {
	return e.Answer
}

func (e *MultipleChoiceExercise) GetType() string {
	return e.Type
}

type FillInTheBlankExercise struct {
	exerciseBase
	Question string // e.g. "This is a ______ truck."
	Answer   string // e.g. "fire"
}

func NewFillInTheBlankExercise(
	id string,
	createdAt, updatedAt time.Time,
	question, answer string,
) FillInTheBlankExercise {
	return FillInTheBlankExercise{
		exerciseBase: newExerciseBase(id, TypeFillInTheBlank, createdAt, updatedAt),
		Question:     question,
		Answer:       answer,
	}
}

func (e *FillInTheBlankExercise) CheckAnswer(answer any) bool {
	if answer, ok := answer.(string); ok {
		return e.Answer == answer
	}
	return false
}

func (e *FillInTheBlankExercise) GetAnswer() any {
	return e.Answer
}

func (e *FillInTheBlankExercise) GetType() string {
	return e.Type
}

type SentenceCorrectionExercise struct {
	exerciseBase
	Sentence          string
	CorrectedSentence string
}

func NewSentenceCorrectionExercise(
	id string,
	createdAt, updatedAt time.Time,
	sentence, correctedSentence string,
) SentenceCorrectionExercise {
	return SentenceCorrectionExercise{
		exerciseBase:      newExerciseBase(id, TypeSentenceCorrection, createdAt, updatedAt),
		Sentence:          sentence,
		CorrectedSentence: correctedSentence,
	}
}

func (e *SentenceCorrectionExercise) CheckAnswer(answer any) bool {
	if answer, ok := answer.(string); ok {
		return e.CorrectedSentence == answer
	}
	return false
}

func (e *SentenceCorrectionExercise) GetAnswer() any {
	return e.CorrectedSentence
}

func (e *SentenceCorrectionExercise) GetType() string {
	return e.Type
}
