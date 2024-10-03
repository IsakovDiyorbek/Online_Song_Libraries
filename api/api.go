package api

import (
	"github.com/Online_Song_Libraries/api/handler"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"

)

func NewGin(h *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/songs", h.AddSong)

	return router
}
