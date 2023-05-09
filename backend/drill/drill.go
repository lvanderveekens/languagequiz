package drill

import (
	"languagequiz/drill/exercise"
	"time"
)

// Drill: Name
//   Section A: Name
// 		1. Exercise
// 		2. Exercise
// 		3. Exercise
//   Section B: Name
// 		1. Exercise
// 		2. Exercise
// 		3. Exercise

type Drill struct {
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
