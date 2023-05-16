package api

import (
	"errors"
	"fmt"

	"languagequiz/quiz/exercise"
)

func mapExerciseToDTO(e exercise.Exercise) (any, error) {
	switch e := e.(type) {
	case *exercise.MultipleChoiceExercise:
		return mapMultipleChoiceExerciseToDTO(*e), nil
	case *exercise.FillInTheBlankExercise:
		return mapFillInTheBlankExerciseToDTO(*e), nil
	case *exercise.SentenceCorrectionExercise:
		return mapSentenceCorrectionExerciseToDTO(*e), nil
	default:
		return nil, fmt.Errorf("unknown exercise type: %T", e)
	}
}

type exerciseDTOBase struct {
	Type string `json:"type"`
}

func newExerciseDTOBase(exerciseType string) exerciseDTOBase {
	return exerciseDTOBase{
		Type: exerciseType,
	}
}

type multipleChoiceExerciseDTO struct {
	exerciseDTOBase
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
}

func newMultipleChoiceExerciseDTO(id, question string, choices []string) multipleChoiceExerciseDTO {
	return multipleChoiceExerciseDTO{
		exerciseDTOBase: newExerciseDTOBase(exercise.TypeMultipleChoice),
		Question:        question,
		Choices:         choices,
	}
}

func mapMultipleChoiceExerciseToDTO(e exercise.MultipleChoiceExercise) multipleChoiceExerciseDTO {
	return newMultipleChoiceExerciseDTO(e.ID, e.Question, e.Choices)
}

type fillInTheBlankExerciseDTO struct {
	exerciseDTOBase
	Question string `json:"question"`
}

func newFillInTheBlankExerciseDTO(id, question string) fillInTheBlankExerciseDTO {
	return fillInTheBlankExerciseDTO{
		exerciseDTOBase: newExerciseDTOBase(exercise.TypeFillInTheBlank),
		Question:        question,
	}
}

func mapFillInTheBlankExerciseToDTO(e exercise.FillInTheBlankExercise) fillInTheBlankExerciseDTO {
	return newFillInTheBlankExerciseDTO(e.ID, e.Question)
}

type sentenceCorrectionExerciseDTO struct {
	exerciseDTOBase
	Sentence string `json:"sentence"`
}

func newSentenceCorrectionExerciseDTO(id, sentence string) sentenceCorrectionExerciseDTO {
	return sentenceCorrectionExerciseDTO{
		exerciseDTOBase: newExerciseDTOBase(exercise.TypeSentenceCorrection),
		Sentence:        sentence,
	}
}

func mapSentenceCorrectionExerciseToDTO(e exercise.SentenceCorrectionExercise) sentenceCorrectionExerciseDTO {
	return newSentenceCorrectionExerciseDTO(e.ID, e.Sentence)
}

type createExerciseRequestBase struct {
	Type     string  `json:"type"`
	Feedback *string `json:"feedback"`
}

type createMultipleChoiceExerciseRequest struct {
	createExerciseRequestBase
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
	Answer   string   `json:"answer"`
}

func (r *createMultipleChoiceExerciseRequest) toCommand() (*exercise.CreateMultipleChoiceExerciseCommand, error) {
	return exercise.NewCreateMultipleChoiceExerciseCommand(
		r.Question,
		r.Choices,
		r.Answer,
		r.Feedback,
	)
}

func (r *createMultipleChoiceExerciseRequest) Validate() error {
	if r.Question == "" {
		return errors.New("required field is missing: question")
	}
	if r.Choices == nil {
		return errors.New("required field is missing: choices")
	}
	if r.Answer == "" {
		return errors.New("required field is missing: answer")
	}
	return nil
}

type createFillInTheBlankExerciseRequest struct {
	createExerciseRequestBase
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (r *createFillInTheBlankExerciseRequest) toCommand() (*exercise.CreateFillInTheBlankExerciseCommand, error) {
	return exercise.NewCreateFillInTheBlankExerciseCommand(
		r.Question,
		r.Answer,
		r.Feedback,
	)
}

func (r *createFillInTheBlankExerciseRequest) Validate() error {
	if r.Question == "" {
		return errors.New("required field is missing: question")
	}
	if r.Answer == "" {
		return errors.New("required field is missing: answer")
	}
	return nil
}

type createSentenceCorrectionExerciseRequest struct {
	createExerciseRequestBase
	Sentence          string `json:"sentence"`
	CorrectedSentence string `json:"correctedSentence"`
}

func (r *createSentenceCorrectionExerciseRequest) toCommand() (*exercise.CreateSentenceCorrectionExerciseCommand, error) {
	return exercise.NewCreateSentenceCorrectionExerciseCommand(
		r.Sentence,
		r.CorrectedSentence,
		r.Feedback,
	)
}

func (r *createSentenceCorrectionExerciseRequest) Validate() error {
	if r.Sentence == "" {
		return errors.New("required field is missing: sentence")
	}
	if r.CorrectedSentence == "" {
		return errors.New("required field is missing: correctedSentence")
	}
	return nil
}
