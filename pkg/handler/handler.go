package handler

import (
	"mas/pkg/service"

	_ "mas/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(cors.Default())

	api := router.Group("/api")
	{
		api.GET("/processedfiles", h.getProcessedFiles)
		api.GET("/processeddata", h.getProcessedData)
		api.GET("/errorsdata", h.getErrorsData)
	}

	return router
}
