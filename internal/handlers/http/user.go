package http

import (
	"net/http"

	"github.com/arezooq/open-utils/api"
	"github.com/arezooq/open-utils/errors"
	"github.com/gin-gonic/gin"
	"user-service/internal/models"
)

// Create
func (h *handler) Create(c *gin.Context) {
	req := api.New(c, "user-service", "v1")

	user := &models.User{}
	if err := req.BindJSON(user); err != nil {
		api.FromAppError(c, errors.ErrInvalidateInput, map[string]string{
			"detail": err.Error(),
		})

		return
	}

	user, appErr := h.userService.CreateUser(req, user)
	if appErr != nil {
		api.FromAppError(c, appErr, nil)
		return
	}

	api.Success(c, http.StatusCreated, "User created successfully", user)
}

// Get all
func (h *handler) GetAll(c *gin.Context) {
	req := api.New(c, "user-service", "v1")

	pagination := api.NewPaginationFromRequest(c)
	query := api.NewQueryFromRequest(c)

	users, total, appErr := h.userService.GetAllUsers(req, pagination, query)
	if appErr != nil {
		api.FromAppError(c, appErr, nil)
		return
	}

	pagination.SetTotal(total)
	api.Success(c, http.StatusOK, "Users fetched successfully", pagination.JSON(users))
}

// Get by id
func (h *handler) Get(c *gin.Context) {
	req := api.New(c, "user-service", "v1")

	userID := c.Param("uuid")
	if userID == "" {
		api.FromAppError(c, errors.ErrInvalidateInput, map[string]string{
			"detail": "missing user ID",
		})
		return
	}

	user, appErr := h.userService.GetUserById(req, userID)
	if appErr != nil {
		api.FromAppError(c, appErr, nil)
		return
	}

	api.Success(c, http.StatusOK, "User fetched successfully", user)
}

// Update
func (h *handler) Update(c *gin.Context) {
	req := api.New(c, "user-service", "v1")

	user := &models.UpdateProfile{}
	if err := req.BindJSON(user); err != nil {
		api.FromAppError(c, errors.ErrInvalidateInput, map[string]string{
			"detail": err.Error(),
		})
		return
	}

	userId := c.Param("uuid")
	if userId == "" {
		api.FromAppError(c, errors.ErrInvalidateInput, map[string]string{
			"detail": "missing user ID",
		})
		return
	}

	updatedUser, appErr := h.userService.UpdateUser(req, userId, user)
	if appErr != nil {
		api.FromAppError(c, appErr, nil)
		return
	}

	api.Success(c, http.StatusOK, "User updated successfully", updatedUser)
}

// Delete
func (h *handler) Delete(c *gin.Context) {
	req := api.New(c, "user-service", "v1")

	userId := c.Param("uuid")
	if userId == "" {
		api.FromAppError(c, errors.ErrInvalidateInput, map[string]string{
			"detail": "missing user ID",
		})
		return
	}

	appErr := h.userService.DeleteUser(req, userId)
	if appErr != nil {
		api.FromAppError(c, appErr, nil)
		return
	}

	api.Success(c, http.StatusOK, "User deleted successfully", nil)
}
