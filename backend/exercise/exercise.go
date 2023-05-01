package exercise

import (
	"reflect"
	"time"
)

type SingleAnswerExercise interface {
	CheckAnswer(answer string) bool
}

type MultipleAnswersExercise interface {
	CheckAnswers(answers []string) bool
}

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
	Question      string
	Options       []string
	CorrectOption string
}

func NewMultipleChoiceExercise(
	exerciseBase ExerciseBase,
	question string,
	options []string,
	correctOption string,
) MultipleChoiceExercise {
	return MultipleChoiceExercise{
		ExerciseBase:  exerciseBase,
		Question:      question,
		Options:       options,
		CorrectOption: correctOption,
	}
}

func (e *MultipleChoiceExercise) CheckAnswer(answer string) bool {
	return e.CorrectOption == answer
}

type CompleteTheSentenceExercise struct {
	ExerciseBase
	Sentence string // e.g. "This is a {0} truck."
	Blank    string // e.g. "fire"
}

func NewCompleteTheSentenceExercise(exercise ExerciseBase, sentence, blank string) CompleteTheSentenceExercise {
	return CompleteTheSentenceExercise{
		ExerciseBase: exercise,
		Sentence:     sentence,
		Blank:        blank,
	}
}

func (e *CompleteTheSentenceExercise) CheckAnswer(answer string) bool {
	return e.Blank == answer
}

type CompleteTheTextExercise struct {
	ExerciseBase
	Text   string   // e.g. "This is a family {0}. Hi, how are {1} doing?"
	Blanks []string // e.g. ["member", "you"]
}

func NewCompleteTheTextExercise(exercise ExerciseBase, text string, blanks []string) CompleteTheTextExercise {
	return CompleteTheTextExercise{
		ExerciseBase: exercise,
		Text:         text,
		Blanks:       blanks,
	}
}

func (e *CompleteTheTextExercise) CheckAnswers(answers []string) bool {
	return reflect.DeepEqual(e.Blanks, answers)
}
