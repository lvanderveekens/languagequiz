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

type Section struct {
	Name      string
	Exercises []exercise.Exercise
}
