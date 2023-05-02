package api

import (
	"fmt"
	"net/http"
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

func (s *Server) Start(port int) error {
	r := gin.Default()

	r.GET("/v1/exercises", createHandlerFunc(s.handlers.exercise.GetExercises))
	r.POST("/v1/exercises", createHandlerFunc(s.handlers.exercise.CreateExercise))
	r.POST("/v1/exercises/answers", createHandlerFunc(s.handlers.exercise.SubmitAnswers))

	return r.Run(":" + strconv.Itoa(port))
}

func createHandlerFunc(f func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := f(c); err != nil {
			if err, ok := err.(Error); ok {
				c.JSON(err.Status, err)
				return
			}

			fmt.Printf("server error: %s\n", err.Error())
			status := http.StatusInternalServerError
			c.JSON(status, NewError(status, http.StatusText(status)))
		}
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
