package http

import (
	"github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

func (h *handler) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api/users")

	// Entity:user
	group.POST("/", h.Create)
	group.GET("/", h.GetAll)
	group.GET("/:userId", h.Get)
	group.PUT("/userId", h.Update)
	group.DELETE("/userId", h.Delete)
	
	// Swagger
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
