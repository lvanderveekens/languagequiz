package exercise

type Storage interface {
	CreateMultipleChoiceExercise(e MultipleChoiceExercise) (*MultipleChoiceExercise, error)
	Find() ([]any, error)
}
