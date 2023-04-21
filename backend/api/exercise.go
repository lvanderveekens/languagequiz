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
	e, err := h.exerciseStorage.CreateExercise()
	if err != nil {
		return fmt.Errorf("failed to create exercise: %w", err)
	}

	dto := NewExercise(e.ID, e.CreatedAt.Format(time.RFC3339), e.UpdatedAt.Format(time.RFC3339))
	c.JSON(http.StatusCreated, dto)
	return nil
}

type Exercise struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func NewExercise(id, createdAt, updatedAt string) Exercise {
	return Exercise{ID: id, CreatedAt: createdAt, UpdatedAt: updatedAt}
}
