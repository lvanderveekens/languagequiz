package exercise

import (
	"time"
)

type ExerciseBase struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewExerciseBase(id string, createdAt, updatedAt time.Time) ExerciseBase {
	return ExerciseBase{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type MultipleChoiceExercise struct {
	ExerciseBase
	Prompt        string
	Options       []string
	CorrectAnswer string
}

func NewMultipleChoiceExercise(
	exerciseBase ExerciseBase,
	prompt string,
	options []string,
	correctAnswer string,
) MultipleChoiceExercise {
	return MultipleChoiceExercise{
		ExerciseBase:  exerciseBase,
		Prompt:        prompt,
		Options:       options,
		CorrectAnswer: correctAnswer,
	}
}

func (e *MultipleChoiceExercise) CheckAnswer(answer string) bool {
	return e.CorrectAnswer == answer
}

type FillInTheBlankExercise struct {
	ExerciseBase
	Prompt        string // e.g. "This is a {0} truck."
	CorrectAnswer string // e.g. "fire"
}

func NewFillInTheBlankExercise(exercise ExerciseBase, prompt, correctAnswer string) FillInTheBlankExercise {
	return FillInTheBlankExercise{
		ExerciseBase:  exercise,
		Prompt:        prompt,
		CorrectAnswer: correctAnswer,
	}
}

func (e *FillInTheBlankExercise) CheckAnswer(answer string) bool {
	return e.CorrectAnswer == answer
}
