package quiz

type CreateQuizCommand struct {
	Name     string
	Sections []CreateQuizSectionCommand
}

func NewCreateQuizCommand(
	name string,
	sections []CreateQuizSectionCommand,
) CreateQuizCommand {
	return CreateQuizCommand{
		Name:     name,
		Sections: sections,
	}
}

type CreateQuizSectionCommand struct {
	Name      string
	Exercises []any
}

func NewCreateSectionCommand(
	name string,
	exercises []any,
) CreateQuizSectionCommand {
	return CreateQuizSectionCommand{
		Name:      name,
		Exercises: exercises,
	}
}
