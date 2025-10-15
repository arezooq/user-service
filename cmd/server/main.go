package main

import (
	"github.com/arezooq/open-utils/logger"
	"github.com/gin-gonic/gin"
	"os"
	_ "user-service/docs"
	"user-service/internal/handlers/http"
	"user-service/internal/repositories/postgres"
	"user-service/internal/services"
)

// @title User Service API
// @version 1.0
// @description This is the User Service API documentation.
// @BasePath /api/users
func main() {
	port := os.Getenv("PORT")

	logger := logger.New("user-service")

	// Postgres
	pgDB, err := postgres.InitPostgres()
	if err != nil {
		logger.Fatal("Failed to init postgres: " + err.Error())
	}

	userRepo := postgres.NewUserRepository(pgDB, logger)

	userService := services.NewUserService(userRepo, logger)
	userHandler := http.InitUserHandler(userService)

	r := gin.Default()
	userHandler.RegisterRoutes(r)
	logger.Info("Server started on port " + port)
	r.Run(":" + port)
}
