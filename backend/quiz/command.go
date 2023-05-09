package quiz

type CreateQuizCommand struct {
	Name     string
	Sections []CreateSectionCommand
}

func NewCreateQuizCommand(
	name string,
	sections []CreateSectionCommand,
) CreateQuizCommand {
	return CreateQuizCommand{
		Name:     name,
		Sections: sections,
	}
}

type CreateSectionCommand struct {
	Name      string
	Exercises []any
}

func NewCreateSectionCommand(
	name string,
	exercises []any,
) CreateSectionCommand {
	return CreateSectionCommand{
		Name:      name,
		Exercises: exercises,
	}
}
