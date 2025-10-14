package http

import "github.com/gin-gonic/gin"

func (h *handler) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api/users")

	// Entity:user
	group.POST("/", h.Create)
	group.GET("/", h.GetAll)
	group.GET("/:userId", h.Get)
	group.PUT("/userId", h.Update)
	group.DELETE("/userId", h.Delete)
}
