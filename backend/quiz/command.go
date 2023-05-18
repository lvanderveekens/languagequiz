package quiz

import (
	"fmt"
	"languagequiz/quiz/exercise"

	"golang.org/x/text/language"
)

type CreateQuizCommand struct {
	Name        string
	LanguageTag language.Tag
	Sections    []CreateSectionCommand
}

func NewCreateQuizCommand(
	name string,
	languageTag language.Tag,
	sections []CreateSectionCommand,
) CreateQuizCommand {
	return CreateQuizCommand{
		Name:        name,
		LanguageTag: languageTag,
		Sections:    sections,
	}
}

type CreateSectionCommand struct {
	Name      string
	Exercises []exercise.CreateExerciseCommand
}

func NewCreateSectionCommand(
	name string,
	exercises []exercise.CreateExerciseCommand,
) (*CreateSectionCommand, error) {
	types := make([]string, 0)
	for _, exercise := range exercises {
		types = append(types, exercise.Type())
	}

	uniqueTypes := removeDuplicates(types)
	if len(uniqueTypes) > 1 {
		return nil, fmt.Errorf("section cannot have more than one exercise type: %v", uniqueTypes)
	}

	return &CreateSectionCommand{
		Name:      name,
		Exercises: exercises,
	}, nil
}

func removeDuplicates[T string](items []T) []T {
	keys := make(map[T]bool)
	list := make([]T, 0)
	for _, item := range items {
		if _, value := keys[item]; !value {
			keys[item] = true
			list = append(list, item)
		}
	}
	return list
}
