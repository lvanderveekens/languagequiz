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

func (h *ExerciseHandler) SubmitAnswers(c *gin.Context) error {
	// TODO: validate request
	var req SubmitAnswersRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		return fmt.Errorf("failed to decode request body: %w", err)
	}

	results := make([]ExerciseResult, 0)
	for _, submission := range req.Submissions {
		e, err := h.exerciseStorage.FindByID(submission.ExerciseID)
		if err != nil {
			return fmt.Errorf("failed to find exercise by id: %w", err)
		}

		correct := e.CheckAnswer(submission.Answer)
		result := NewExerciseResult(submission.ExerciseID, e.GetType(), correct, e.GetAnswer())
		results = append(results, result)
	}

	c.JSON(http.StatusOK, NewSubmitAnswersResponse(results))
	return nil
}

func mapExerciseToDto(e exercise.Exercise) (any, error) {
	switch e := e.(type) {
	case *exercise.MultipleChoiceExercise:
		return newMultipleChoiceExerciseDto(*e), nil
	case *exercise.FillInTheBlankExercise:
		return newFillInTheBlankExerciseDto(*e), nil
	default:
		return nil, fmt.Errorf("unknown exercise type: %T", e)
	}
}

type ExerciseDto struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func newExerciseDto(id string, exerciseType string) ExerciseDto {
	return ExerciseDto{
		ID:   id,
		Type: exerciseType,
	}
}

type MultipleChoiceExerciseDto struct {
	ExerciseDto
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}

func newMultipleChoiceExerciseDto(e exercise.MultipleChoiceExercise) MultipleChoiceExerciseDto {
	return MultipleChoiceExerciseDto{
		ExerciseDto: newExerciseDto(e.ID, exercise.TypeMultipleChoice),
		Question:    e.Question,
		Options:     e.Options,
		Answer:      e.Answer,
	}
}

type FillInTheBlankExerciseDto struct {
	ExerciseDto
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func newFillInTheBlankExerciseDto(e exercise.FillInTheBlankExercise) FillInTheBlankExerciseDto {
	return FillInTheBlankExerciseDto{
		ExerciseDto: newExerciseDto(e.ID, exercise.TypeFillInTheBlank),
		Question:    e.Question,
		Answer:      e.Answer,
	}
}

type CreateExerciseRequest struct {
	Type string `json:"type"`
}

type CreateMultipleChoiceExerciseRequest struct {
	CreateExerciseRequest
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}

func (r *CreateMultipleChoiceExerciseRequest) toCommand() exercise.CreateMultipleChoiceExerciseCommand {
	return exercise.NewCreateMultipleChoiceExerciseCommand(
		r.Question,
		r.Options,
		r.Answer,
	)
}

func (r *CreateMultipleChoiceExerciseRequest) Validate() error {
	if r.Question == "" {
		return errors.New("required field is missing: question")
	}
	if r.Options == nil {
		return errors.New("required field is missing: options")
	}
	if r.Answer == "" {
		return errors.New("required field is missing: answer")
	}
	return nil
}

type CreateFillInTheBlankExerciseRequest struct {
	CreateExerciseRequest
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (r *CreateFillInTheBlankExerciseRequest) toCommand() (*exercise.CreateFillInTheBlankExerciseCommand, error) {
	return exercise.NewCreateFillInTheBlankExerciseCommand(
		r.Question,
		r.Answer,
	)
}

func (r *CreateFillInTheBlankExerciseRequest) Validate() error {
	if r.Question == "" {
		return errors.New("required field is missing: question")
	}
	if r.Answer == "" {
		return errors.New("required field is missing: answer")
	}
	return nil
}

type SubmitAnswersRequest struct {
	Submissions []ExerciseSubmission `json:"submissions"`
}

type ExerciseSubmission struct {
	ExerciseID string `json:"exerciseId"`
	Answer     any    `json:"answer"`
}

type SubmitAnswersResponse struct {
	Results []ExerciseResult `json:"results"`
}

func NewSubmitAnswersResponse(results []ExerciseResult) SubmitAnswersResponse {
	return SubmitAnswersResponse{
		Results: results,
	}
}

type ExerciseResult struct {
	ExerciseID   string `json:"exerciseId"`
	ExerciseType string `json:"exerciseType"`
	Correct      bool   `json:"correct"`
	Answer       any    `json:"answer"`
}

func NewExerciseResult(exerciseID string, exerciseType string, correct bool, answer any) ExerciseResult {
	return ExerciseResult{
		ExerciseID:   exerciseID,
		ExerciseType: exerciseType,
		Correct:      correct,
		Answer:       answer,
	}
}
