package api

import (
	"fmt"
	"net/http"

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

	// FIXME: doesn't log timezone
	dto := NewExercise(e.ID.String(), e.CreatedAt.Format("2006-01-02T15:04:05Z07:00"))
	c.JSON(http.StatusCreated, dto)
	return nil
}

type Exercise struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

func NewExercise(id, createdAt string) Exercise {
	return Exercise{ID: id, CreatedAt: createdAt}
}
