package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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
		err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&req)
		if err != nil {
			return NewError(http.StatusBadRequest, fmt.Sprintf("failed to decode request body as struct: %s", err.Error()))
		}

		exercise, err := h.exerciseStorage.CreateMultipleChoiceExercise(req.toCommand())
		if err != nil {
			return fmt.Errorf("failed to create exercise: %w", err)
		}

		dto = mapMultipleChoiceExerciseToDto(*exercise)
	case exercise.TypeCompleteTheSentence:
		var req CreateCompleteTheSentenceExerciseRequest
		err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&req)
		if err != nil {
			return NewError(http.StatusBadRequest, fmt.Sprintf("failed to decode request body as struct: %s", err.Error()))
		}

		exercise, err := h.exerciseStorage.CreateCompleteTheSentenceExercise(req.toCommand())
		if err != nil {
			return fmt.Errorf("failed to create exercise: %w", err)
		}

		dto = mapCompleteTheSentenceExerciseToDto(*exercise)
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

func mapMultipleChoiceExerciseToDto(e exercise.MultipleChoiceExercise) MultipleChoiceExercise {
	return newMultipleChoiceExercise(
		NewExercise(
			e.ID,
			e.CreatedAt.Format(time.RFC3339),
			e.UpdatedAt.Format(time.RFC3339),
			exercise.TypeMultipleChoice,
		),
		e.Question,
		e.Options,
		e.CorrectOption,
	)
}

func mapCompleteTheSentenceExerciseToDto(e exercise.CompleteTheSentenceExercise) CompleteTheSentenceExercise {
	return NewCompleteTheSentenceExercise(
		NewExercise(
			e.ID,
			e.CreatedAt.Format(time.RFC3339),
			e.UpdatedAt.Format(time.RFC3339),
			exercise.TypeCompleteTheSentence,
		),
		e.BeforeGap,
		e.Gap,
		e.AfterGap,
	)
}

func mapExerciseToDto(e any) (any, error) {
	switch e := e.(type) {
	case exercise.MultipleChoiceExercise:
		return mapMultipleChoiceExerciseToDto(e), nil
	case exercise.CompleteTheSentenceExercise:
		return mapCompleteTheSentenceExerciseToDto(e), nil
	default:
		return nil, fmt.Errorf("unknown exercise type: %T", e)
	}
}

type Exercise struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func NewExercise(id, createdAt, updatedAt, exerciseType string) Exercise {
	return Exercise{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Type:      exerciseType,
	}
}

type MultipleChoiceExercise struct {
	Exercise
	Question      string   `json:"question"`
	Options       []string `json:"options"`
	CorrectOption string   `json:"correctOption"`
}

func newMultipleChoiceExercise(
	exercise Exercise,
	question string,
	options []string,
	correctOption string,
) MultipleChoiceExercise {
	return MultipleChoiceExercise{
		Exercise:      exercise,
		Question:      question,
		Options:       options,
		CorrectOption: correctOption,
	}
}

type CompleteTheSentenceExercise struct {
	Exercise
	BeforeGap string `json:"beforeGap"`
	Gap       string `json:"gap"`
	AfterGap  string `json:"afterGap"`
}

func NewCompleteTheSentenceExercise(
	exercise Exercise,
	beforeGap, gap, afterGap string,
) CompleteTheSentenceExercise {
	return CompleteTheSentenceExercise{
		Exercise:  exercise,
		BeforeGap: beforeGap,
		Gap:       gap,
		AfterGap:  afterGap,
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

type CreateCompleteTheSentenceExerciseRequest struct {
	CreateExerciseRequest
	BeforeGap string `json:"beforeGap"`
	Gap       string `json:"gap"`
	AfterGap  string `json:"afterGap"`
}

func (r *CreateCompleteTheSentenceExerciseRequest) toCommand() exercise.CreateCompleteTheSentenceExerciseCommand {
	return exercise.NewCreateCompleteTheSentenceExerciseCommand(
		r.BeforeGap,
		r.Gap,
		r.AfterGap,
	)
}
