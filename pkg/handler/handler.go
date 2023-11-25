package handler

import (
	"mas/pkg/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())

	api := router.Group("/api")
	{
		api.GET("/processedfiles", h.getProcessedFiles)
		api.GET("/processeddata", h.getProcessedData)
		api.GET("/errorsdata", h.getErrorsData)
	}

	return router
}
