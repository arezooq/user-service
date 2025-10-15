package http

import (
	"net/http"

	"github.com/arezooq/open-utils/api"
	"github.com/arezooq/open-utils/errors"
	"github.com/arezooq/open-utils/jwt"
	"github.com/gin-gonic/gin"
	"user-service/internal/models"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User info"
// @Success 201 {object} models.User
// @Router /users [post]
func (h *handler) Create(c *gin.Context) {
	req := api.New(c, "user-service", "v1")

	_, err := jwt.ExtractTokenFromHeader(req)
	if err != nil {
		api.FromAppError(c, errors.ErrUnauthorized, nil)
		return
	}

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

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
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

// GetUser godoc
// @Summary Get a user by ID
// @Description Get user details by user ID
// @Tags users
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} models.User
// @Router /users/{userId} [get]
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

// UpdateUser godoc
// @Summary Update a user
// @Description Update user fields like name, email, mobile
// @Tags users
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param user body models.UpdateProfile true "Updated user info"
// @Success 200 {object} models.User
// @Router /users/{userId} [put]
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

// DeleteUser godoc
// @Summary Delete a user
// @Description Remove user by ID
// @Tags users
// @Produce json
// @Param userId path string true "User ID"
// @Success 204
// @Router /users/{userId} [delete]
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
