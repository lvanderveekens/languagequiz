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
	BeforeGap string
	Gap       string
	AfterGap  string
}

func NewCompleteTheSentenceExercise(exercise Exercise, beforeGap, gap, afterGap string) CompleteTheSentenceExercise {
	return CompleteTheSentenceExercise{
		Exercise:  exercise,
		BeforeGap: beforeGap,
		Gap:       gap,
		AfterGap:  afterGap,
	}
}
