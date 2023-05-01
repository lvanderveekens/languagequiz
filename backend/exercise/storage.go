package exercise

type Storage interface {
	CreateMultipleChoiceExercise(e CreateMultipleChoiceExerciseCommand) (*MultipleChoiceExercise, error)
	CreateFillInTheBlankExercise(e CreateFillInTheBlankExerciseCommand) (*FillInTheBlankExercise, error)
	Find() ([]any, error)
}
