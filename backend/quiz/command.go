package quiz

import "golang.org/x/text/language"

type CreateQuizCommand struct {
	Name        string
	LanguageTag language.Tag
	Sections    []CreateQuizSectionCommand
}

func NewCreateQuizCommand(
	name string,
	languageTag language.Tag,
	sections []CreateQuizSectionCommand,
) CreateQuizCommand {
	return CreateQuizCommand{
		Name:        name,
		LanguageTag: languageTag,
		Sections:    sections,
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
