package exercise

type Storage interface {
	CreateMultipleChoiceExercise(e CreateMultipleChoiceExerciseCommand) (*MultipleChoiceExercise, error)
	Find() ([]any, error)
}
