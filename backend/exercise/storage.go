package exercise

type Storage interface {
	CreateMultipleChoiceExercise(e CreateMultipleChoiceExerciseCommand) (*MultipleChoiceExercise, error)
	CreateCompleteTheSentenceExercise(e CreateCompleteTheSentenceExerciseCommand) (*CompleteTheSentenceExercise, error)
	Find() ([]any, error)
}
