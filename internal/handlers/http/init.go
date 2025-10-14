package http

import (
	"user-service/internal/services"
	"github.com/gin-gonic/gin"
)

type HandlerUserInterface interface {
	RegisterRoutes(router *gin.Engine)

	// Entity:user
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type handler struct {
	userService services.UserService
}

func InitUserHandler(userService services.UserService) HandlerUserInterface {
	return &handler{userService: userService}
}