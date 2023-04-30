package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvanderveekens/testmaker/exercise"
)

type ExerciseHandler struct {
	exerciseStorage exercise.Storage
}

func NewExerciseHandler(exerciseStorage exercise.Storage) *ExerciseHandler {
	return &ExerciseHandler{
		exerciseStorage: exerciseStorage,
	}
}

func (h *ExerciseHandler) CreateExercise(c *gin.Context) error {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %w", err)
	}

	var req CreateExerciseRequest
	err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&req)
	if err != nil {
		return fmt.Errorf("failed to decode request body as map: %w", err)
	}

	var dto any
	switch req.Type {
	case exercise.TypeMultipleChoice:
		var req CreateMultipleChoiceExerciseRequest
		if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&req); err != nil {
			return NewError(http.StatusBadRequest, fmt.Sprintf("failed to decode request body as struct: %s", err.Error()))
		}
		if err := req.Validate(); err != nil {
			return NewError(http.StatusBadRequest, err.Error())
		}

		exercise, err := h.exerciseStorage.CreateMultipleChoiceExercise(req.toCommand())
		if err != nil {
			return fmt.Errorf("failed to create exercise: %w", err)
		}

		dto = newMultipleChoiceExerciseDto(*exercise)
	case exercise.TypeCompleteTheSentence:
		var req CreateCompleteTheSentenceExerciseRequest
		if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&req); err != nil {
			return NewError(http.StatusBadRequest, fmt.Sprintf("failed to decode request body as struct: %s", err.Error()))
		}
		if err := req.Validate(); err != nil {
			return NewError(http.StatusBadRequest, err.Error())
		}

		cmd, err := req.toCommand()
		if err != nil {
			return NewError(http.StatusBadRequest, err.Error())
		}

		exercise, err := h.exerciseStorage.CreateCompleteTheSentenceExercise(*cmd)
		if err != nil {
			return fmt.Errorf("failed to create exercise: %w", err)
		}

		dto = newCompleteTheSentenceExerciseDto(*exercise)
	case exercise.TypeCompleteTheText:
		var req CreateCompleteTheTextExerciseRequest
		if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&req); err != nil {
			return NewError(http.StatusBadRequest, fmt.Sprintf("failed to decode request body as struct: %s", err.Error()))
		}
		if err := req.Validate(); err != nil {
			return NewError(http.StatusBadRequest, err.Error())
		}

		cmd, err := req.toCommand()
		if err != nil {
			return NewError(http.StatusBadRequest, err.Error())
		}

		exercise, err := h.exerciseStorage.CreateCompleteTheTextExercise(*cmd)
		if err != nil {
			return fmt.Errorf("failed to create exercise: %w", err)
		}

		dto = newCompleteTheTextExerciseDto(*exercise)
	default:
		return NewError(http.StatusBadRequest, fmt.Sprintf("Unsupported exercise type: %v", req.Type))
	}

	c.JSON(http.StatusCreated, dto)
	return nil
}

func (h *ExerciseHandler) GetExercises(c *gin.Context) error {
	exercises, err := h.exerciseStorage.Find()
	if err != nil {
		return fmt.Errorf("failed to find exercises: %w", err)
	}

	dtos := make([]any, 0)
	for _, exercise := range exercises {
		dto, err := mapExerciseToDto(exercise)
		if err != nil {
			return fmt.Errorf("failed to map exercise to dto: %w", err)
		}
		dtos = append(dtos, dto)
	}

	c.JSON(http.StatusOK, dtos)
	return nil
}

func mapExerciseToDto(e any) (any, error) {
	switch e := e.(type) {
	case exercise.MultipleChoiceExercise:
		return newMultipleChoiceExerciseDto(e), nil
	case exercise.CompleteTheSentenceExercise:
		return newCompleteTheSentenceExerciseDto(e), nil
	case exercise.CompleteTheTextExercise:
		return newCompleteTheTextExerciseDto(e), nil
	default:
		return nil, fmt.Errorf("unknown exercise type: %T", e)
	}
}

type ExerciseDto struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func newExerciseDto(e exercise.Exercise, exerciseType string) ExerciseDto {
	return ExerciseDto{
		ID:   e.ID,
		Type: exerciseType,
	}
}

type MultipleChoiceExerciseDto struct {
	ExerciseDto
	Question      string   `json:"question"`
	Options       []string `json:"options"`
	CorrectOption string   `json:"correctOption"`
}

func newMultipleChoiceExerciseDto(e exercise.MultipleChoiceExercise) MultipleChoiceExerciseDto {
	return MultipleChoiceExerciseDto{
		ExerciseDto:   newExerciseDto(e.Exercise, exercise.TypeMultipleChoice),
		Question:      e.Question,
		Options:       e.Options,
		CorrectOption: e.CorrectOption,
	}
}

type CompleteTheSentenceExerciseDto struct {
	ExerciseDto
	Sentence string `json:"sentence"`
	Blank    string `json:"blank"`
}

func newCompleteTheSentenceExerciseDto(
	e exercise.CompleteTheSentenceExercise,
) CompleteTheSentenceExerciseDto {
	return CompleteTheSentenceExerciseDto{
		ExerciseDto: newExerciseDto(e.Exercise, exercise.TypeCompleteTheSentence),
		Sentence:    e.Sentence,
		Blank:       e.Blank,
	}
}

type CompleteTheTextExerciseDto struct {
	ExerciseDto
	Text   string   `json:"text"`
	Blanks []string `json:"blanks"`
}

func newCompleteTheTextExerciseDto(e exercise.CompleteTheTextExercise) CompleteTheTextExerciseDto {
	return CompleteTheTextExerciseDto{
		ExerciseDto: newExerciseDto(e.Exercise, exercise.TypeCompleteTheText),
		Text:        e.Text,
		Blanks:      e.Blanks,
	}
}

type CreateExerciseRequest struct {
	Type string `json:"type"`
}

type CreateMultipleChoiceExerciseRequest struct {
	CreateExerciseRequest
	Question      string   `json:"question"`
	Options       []string `json:"options"`
	CorrectOption string   `json:"correctOption"`
}

func (r *CreateMultipleChoiceExerciseRequest) toCommand() exercise.CreateMultipleChoiceExerciseCommand {
	return exercise.NewCreateMultipleChoiceExerciseCommand(
		r.Question,
		r.Options,
		r.CorrectOption,
	)
}

func (r *CreateMultipleChoiceExerciseRequest) Validate() error {
	if r.Question == "" {
		return errors.New("required field is missing: question")
	}
	if r.Options == nil {
		return errors.New("required field is missing: options")
	}
	if r.CorrectOption == "" {
		return errors.New("required field is missing: correctOption")
	}
	return nil
}

type CreateCompleteTheSentenceExerciseRequest struct {
	CreateExerciseRequest
	Sentence string `json:"sentence"`
	Blank    string `json:"blank"`
}

func (r *CreateCompleteTheSentenceExerciseRequest) toCommand() (*exercise.CreateCompleteTheSentenceExerciseCommand, error) {
	return exercise.NewCreateCompleteTheSentenceExerciseCommand(
		r.Sentence,
		r.Blank,
	)
}

func (r *CreateCompleteTheSentenceExerciseRequest) Validate() error {
	if r.Sentence == "" {
		return errors.New("required field is missing: sentence")
	}
	if r.Blank == "" {
		return errors.New("required field is missing: blank")
	}
	return nil
}

type CreateCompleteTheTextExerciseRequest struct {
	CreateExerciseRequest
	Text   string   `json:"text"`
	Blanks []string `json:"blanks"`
}

func (r *CreateCompleteTheTextExerciseRequest) toCommand() (*exercise.CreateCompleteTheTextExerciseCommand, error) {
	return exercise.NewCreateCompleteTheTextExerciseCommand(r.Text, r.Blanks)
}

func (r *CreateCompleteTheTextExerciseRequest) Validate() error {
	if r.Text == "" {
		return errors.New("required field is missing: text")
	}
	if r.Blanks == nil {
		return errors.New("required field is missing: blanks")
	}
	return nil
}
