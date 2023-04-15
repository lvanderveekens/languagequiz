package api

import (
	"log"
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

func (h *ExerciseHandler) CreateExercise(c *gin.Context) {
	_, err := h.exerciseStorage.CreateExercise()
	if err != nil {
		log.Fatalf("error: failed to create exercise: %s", err)
	}

	// TODO: map to JSON and return

	c.JSON(http.StatusCreated, gin.H{})
}
