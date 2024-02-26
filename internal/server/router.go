package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HealthCheckHandler)
	r.POST("/sleep", s.SleepHandler)

	return r
}
