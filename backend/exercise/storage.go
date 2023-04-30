package exercise

type Storage interface {
	CreateMultipleChoiceExercise(e CreateMultipleChoiceExerciseCommand) (*MultipleChoiceExercise, error)
	CreateCompleteTheSentenceExercise(e CreateCompleteTheSentenceExerciseCommand) (*CompleteTheSentenceExercise, error)
	CreateCompleteTheTextExercise(e CreateCompleteTheTextExerciseCommand) (*CompleteTheTextExercise, error)
	Find() ([]any, error)
}
