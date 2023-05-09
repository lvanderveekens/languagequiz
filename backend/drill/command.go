package drill

type CreateDrillCommand struct {
	Name     string
	Sections []CreateSectionCommand
}

func NewCreateDrillCommand(
	name string,
	sections []CreateSectionCommand,
) CreateDrillCommand {
	return CreateDrillCommand{
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
