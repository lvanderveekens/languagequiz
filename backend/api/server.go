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

	r.GET("/v1/quizzes", createHandlerFunc(s.handlers.quiz.GetQuizzes))
	r.GET("/v1/quizzes/:id", createHandlerFunc(s.handlers.quiz.GetQuizByID))
	r.POST("/v1/quizzes", createHandlerFunc(s.handlers.quiz.CreateQuiz))
	r.POST("/v1/quizzes/:id/answers", createHandlerFunc(s.handlers.quiz.SubmitAnswers))
	r.POST("/v1/feedback", createHandlerFunc(s.handlers.feedback.ReceiveFeedback))

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
	quiz     *QuizHandler
	feedback *FeedbackHandler
}

func NewHandlers(quizHandler *QuizHandler, feedbackHandler *FeedbackHandler) *Handlers {
	return &Handlers{
		quiz:     quizHandler,
		feedback: feedbackHandler,
	}
}
