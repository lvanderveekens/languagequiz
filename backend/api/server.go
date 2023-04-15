package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	handlers *Handlers
}

func NewServer(handlers *Handlers) *Server {
	return &Server{
		handlers: handlers,
	}
}

type Handlers struct {
	exercise *ExerciseHandler
}

func NewHandlers(exerciseHandler *ExerciseHandler) *Handlers {
	return &Handlers{
		exercise: exerciseHandler,
	}
}

func (s *Server) Start(port int) error {
	r := gin.Default()

	r.POST("/v1/exercises", s.handlers.exercise.CreateExercise)

	return r.Run(":" + strconv.Itoa(port))
}
