package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvanderveekens/testparrot/exercise"
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
	case exercise.TypeFillInTheBlank:
		var req CreateFillInTheBlankExerciseRequest
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

		exercise, err := h.exerciseStorage.CreateFillInTheBlankExercise(*cmd)
		if err != nil {
			return fmt.Errorf("failed to create exercise: %w", err)
		}

		dto = newFillInTheBlankExerciseDto(*exercise)
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
	case exercise.FillInTheBlankExercise:
		return newFillInTheBlankExerciseDto(e), nil
	default:
		return nil, fmt.Errorf("unknown exercise type: %T", e)
	}
}

type ExerciseDto struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func newExerciseDto(e exercise.ExerciseBase, exerciseType string) ExerciseDto {
	return ExerciseDto{
		ID:   e.ID,
		Type: exerciseType,
	}
}

type MultipleChoiceExerciseDto struct {
	ExerciseDto
	Prompt        string   `json:"prompt"`
	Options       []string `json:"options"`
	CorrectAnswer string   `json:"correctAnswer"`
}

func newMultipleChoiceExerciseDto(e exercise.MultipleChoiceExercise) MultipleChoiceExerciseDto {
	return MultipleChoiceExerciseDto{
		ExerciseDto:   newExerciseDto(e.ExerciseBase, exercise.TypeMultipleChoice),
		Prompt:        e.Prompt,
		Options:       e.Options,
		CorrectAnswer: e.CorrectAnswer,
	}
}

type FillInTheBlankExerciseDto struct {
	ExerciseDto
	Prompt        string `json:"prompt"`
	CorrectAnswer string `json:"correctAnswer"`
}

func newFillInTheBlankExerciseDto(e exercise.FillInTheBlankExercise) FillInTheBlankExerciseDto {
	return FillInTheBlankExerciseDto{
		ExerciseDto:   newExerciseDto(e.ExerciseBase, exercise.TypeFillInTheBlank),
		Prompt:        e.Prompt,
		CorrectAnswer: e.CorrectAnswer,
	}
}

type CreateExerciseRequest struct {
	Type string `json:"type"`
}

type CreateMultipleChoiceExerciseRequest struct {
	CreateExerciseRequest
	Prompt        string   `json:"prompt"`
	Options       []string `json:"options"`
	CorrectAnswer string   `json:"correctAnswer"`
}

func (r *CreateMultipleChoiceExerciseRequest) toCommand() exercise.CreateMultipleChoiceExerciseCommand {
	return exercise.NewCreateMultipleChoiceExerciseCommand(
		r.Prompt,
		r.Options,
		r.CorrectAnswer,
	)
}

func (r *CreateMultipleChoiceExerciseRequest) Validate() error {
	if r.Prompt == "" {
		return errors.New("required field is missing: prompt")
	}
	if r.Options == nil {
		return errors.New("required field is missing: options")
	}
	if r.CorrectAnswer == "" {
		return errors.New("required field is missing: correctAnswer")
	}
	return nil
}

type CreateFillInTheBlankExerciseRequest struct {
	CreateExerciseRequest
	Prompt        string `json:"prompt"`
	CorrectAnswer string `json:"correctAnswer"`
}

func (r *CreateFillInTheBlankExerciseRequest) toCommand() (*exercise.CreateFillInTheBlankExerciseCommand, error) {
	return exercise.NewCreateFillInTheBlankExerciseCommand(
		r.Prompt,
		r.CorrectAnswer,
	)
}

func (r *CreateFillInTheBlankExerciseRequest) Validate() error {
	if r.Prompt == "" {
		return errors.New("required field is missing: prompt")
	}
	if r.CorrectAnswer == "" {
		return errors.New("required field is missing: correctAnswer")
	}
	return nil
}
