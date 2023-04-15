package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleCreateExercise(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
