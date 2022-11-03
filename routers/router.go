package routers

import (
	"github.com/gin-gonic/gin"
	"init_project/internal/service"
)

func LoadRouters(s *service.Service) *gin.Engine {
	router := gin.Default()
	router.GET("/healthcheck", s.HealthCheck)
	return router
}

