package exercise

import (
	"time"
)

type Exercise struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(id string, createdAt, updatedAt time.Time) Exercise {
	return Exercise{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type MultipleChoiceExercise struct {
	Exercise
	Question      string
	Options       []string
	CorrectOption string
}

func NewMultipleChoiceExercise(
	exercise Exercise,
	question string,
	options []string,
	correctOption string,
) MultipleChoiceExercise {
	return MultipleChoiceExercise{
		Exercise:      exercise,
		Question:      question,
		Options:       options,
		CorrectOption: correctOption,
	}
}

type CompleteTheSentenceExercise struct {
	Exercise
	Sentence string // e.g. "This is a {0} truck."
	Blank    string // e.g. "fire"
}

func NewCompleteTheSentenceExercise(exercise Exercise, sentence, blank string) CompleteTheSentenceExercise {
	return CompleteTheSentenceExercise{
		Exercise: exercise,
		Sentence: sentence,
		Blank:    blank,
	}
}

type CompleteTheTextExercise struct {
	Exercise
	Text   string   // e.g. "This is a family {0}. Hi, how are {1} doing?"
	Blanks []string // e.g. ["member", "you"]
}

func NewCompleteTheTextExercise(exercise Exercise, text string, blanks []string) CompleteTheTextExercise {
	return CompleteTheTextExercise{
		Exercise: exercise,
		Text:     text,
		Blanks:   blanks,
	}
}
