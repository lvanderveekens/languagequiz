package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(port int) error {
	r := gin.Default()

	r.POST("/v1/exercises", s.handleCreateExercise)

	return r.Run(":" + strconv.Itoa(port))
}
