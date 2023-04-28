package exercise

type CreateMultipleChoiceExerciseCommand struct {
	Question      string
	Options       []string
	CorrectOption string
}

func NewCreateMultipleChoiceExerciseCommand(
	question string,
	options []string,
	correctOption string,
) CreateMultipleChoiceExerciseCommand {
	return CreateMultipleChoiceExerciseCommand{
		Question:      question,
		Options:       options,
		CorrectOption: correctOption,
	}
}
