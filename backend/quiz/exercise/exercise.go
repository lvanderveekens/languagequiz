package exercise

import (
	"time"
)

type Exercise interface {
	CheckAnswer(answer any) bool
	Answer() any
	Feedback() *string
}

type exerciseBase struct {
	ID        string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	feedback  *string
}

func (b *exerciseBase) Feedback() *string {
	return b.feedback
}

func newExerciseBase(id, _type string, createdAt, updatedAt time.Time, feedback *string) exerciseBase {
	return exerciseBase{
		ID:        id,
		Type:      _type,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		feedback:  feedback,
	}
}

type MultipleChoiceExercise struct {
	exerciseBase
	Question string
	Choices  []string
	answer   string
}

func NewMultipleChoiceExercise(
	id string,
	createdAt, updatedAt time.Time,
	feedback *string,
	question string,
	choices []string,
	answer string,
) MultipleChoiceExercise {
	return MultipleChoiceExercise{
		exerciseBase: newExerciseBase(id, TypeMultipleChoice, createdAt, updatedAt, feedback),
		Question:     question,
		Choices:      choices,
		answer:       answer,
	}
}

func (e *MultipleChoiceExercise) CheckAnswer(answer any) bool {
	if answer, ok := answer.(string); ok {
		return e.answer == answer
	}
	return false
}

func (e *MultipleChoiceExercise) Answer() any {
	return e.answer
}

type FillInTheBlankExercise struct {
	exerciseBase
	Question string // e.g. "This is a ______ truck."
	answer   string // e.g. "fire"
}

func NewFillInTheBlankExercise(
	id string,
	createdAt, updatedAt time.Time,
	feedback *string,
	question, answer string,
) FillInTheBlankExercise {
	return FillInTheBlankExercise{
		exerciseBase: newExerciseBase(id, TypeFillInTheBlank, createdAt, updatedAt, feedback),
		Question:     question,
		answer:       answer,
	}
}

func (e *FillInTheBlankExercise) CheckAnswer(answer any) bool {
	if answer, ok := answer.(string); ok {
		return normalizeAnswer(e.answer) == normalizeAnswer(answer)
	}
	return false
}

func (e *FillInTheBlankExercise) Answer() any {
	return e.answer
}

type SentenceCorrectionExercise struct {
	exerciseBase
	Sentence          string
	CorrectedSentence string
}

func NewSentenceCorrectionExercise(
	id string,
	createdAt, updatedAt time.Time,
	feedback *string,
	sentence, correctedSentence string,
) SentenceCorrectionExercise {
	return SentenceCorrectionExercise{
		exerciseBase:      newExerciseBase(id, TypeSentenceCorrection, createdAt, updatedAt, feedback),
		Sentence:          sentence,
		CorrectedSentence: correctedSentence,
	}
}

func (e *SentenceCorrectionExercise) CheckAnswer(answer any) bool {
	if answer, ok := answer.(string); ok {
		return normalizeAnswer(e.CorrectedSentence) == normalizeAnswer(answer)
	}
	return false
}

func (e *SentenceCorrectionExercise) Answer() any {
	return e.CorrectedSentence
}
