package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lvanderveekens/language-resources/exercise"
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
	exercise, err := h.exerciseStorage.Create()
	if err != nil {
		return fmt.Errorf("failed to create exercise: %w", err)
	}

	dto := h.toDto(*exercise)
	c.JSON(http.StatusCreated, dto)
	return nil
}

func (h *ExerciseHandler) GetExercises(c *gin.Context) error {
	exercises, err := h.exerciseStorage.Find()
	if err != nil {
		return fmt.Errorf("failed to find exercises: %w", err)
	}

	dtos := make([]Exercise, 0)
	for _, exercise := range exercises {
		dtos = append(dtos, h.toDto(exercise))
	}

	c.JSON(http.StatusOK, dtos)
	return nil
}

func (h *ExerciseHandler) toDto(e exercise.Exercise) Exercise {
	return NewExercise(e.ID, e.CreatedAt.Format(time.RFC3339), e.UpdatedAt.Format(time.RFC3339))
}

type Exercise struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func NewExercise(id, createdAt, updatedAt string) Exercise {
	return Exercise{ID: id, CreatedAt: createdAt, UpdatedAt: updatedAt}
}
