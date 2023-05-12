package quiz

import (
	"languagequiz/quiz/exercise"
	"time"
)

// Quiz: Name
//   Section A: Name
// 		1. Exercise
// 		2. Exercise
// 		3. Exercise
//   Section B: Name
// 		1. Exercise
// 		2. Exercise
// 		3. Exercise

type Quiz struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Sections  []Section
}

func (q *Quiz) GetExercises() []exercise.Exercise {
	exercises := make([]exercise.Exercise, 0)
	for _, s := range q.Sections {
		exercises = append(exercises, s.Exercises...)
	}
	return exercises
}

func New(id string, createdAt, updatedAt time.Time, name string, sections []Section) Quiz {
	return Quiz{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      name,
		Sections:  sections,
	}
}

type Section struct {
	Name      string
	Exercises []exercise.Exercise
}

func NewSection(name string, exercises []exercise.Exercise) Section {
	return Section{
		Name:      name,
		Exercises: exercises,
	}
}
