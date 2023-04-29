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

type CreateCompleteTheSentenceExerciseCommand struct {
	BeforeGap string
	Gap       string
	AfterGap  string
}

func NewCreateCompleteTheSentenceExerciseCommand(
	beforeGap, gap, afterGap string,
) CreateCompleteTheSentenceExerciseCommand {
	return CreateCompleteTheSentenceExerciseCommand{
		BeforeGap: beforeGap,
		Gap:       gap,
		AfterGap:  afterGap,
	}
}
