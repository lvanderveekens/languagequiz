package exercise

type Storage interface {
	CreateMultipleChoiceExercise(e CreateMultipleChoiceExerciseCommand) (*MultipleChoiceExercise, error)
	CreateFillInTheBlankExercise(e CreateFillInTheBlankExerciseCommand) (*FillInTheBlankExercise, error)
	CreateSentenceCorrectionExercise(e CreateSentenceCorrectionExerciseCommand) (*SentenceCorrectionExercise, error)

	Find() ([]Exercise, error)
	FindByID(id string) (Exercise, error)
}
