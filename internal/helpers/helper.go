package helper

import (
	"time"
	"user-service/internal/models"

	"github.com/arezooq/open-utils/jwt"
	"github.com/arezooq/open-utils/uuid"
	"github.com/gin-gonic/gin"
)

type AuditFields struct {
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
}

// Pre Create
func PreCreate(c *gin.Context, user *models.User) error {
	token, err := jwt.ExtractTokenFromHeader(c)

	userID, err := jwt.ExtractUserIDFromToken(token)
	if err != nil {
		return err
	}

	user.ID = uuid.UUIDString()

	now := time.Now().UTC()

	user.CreatedBy = userID
	user.CreatedAt = now
	user.Status = 0

	return nil
}

// Pre Update
func PreUpdate(c *gin.Context, user *models.UpdateProfile) error {
	token, err := jwt.ExtractTokenFromHeader(c)

	userID, err := jwt.ExtractUserIDFromToken(token)
	if err != nil {
		return err
	}

	now := time.Now().UTC()

	user.UpdatedBy = userID
	user.UpdatedAt = now

	return nil
}

// Pre Delete
func PreDelete(c *gin.Context, user *models.User) error {
	token, err := jwt.ExtractTokenFromHeader(c)

	userID, err := jwt.ExtractUserIDFromToken(token)
	if err != nil {
		return err
	}

	now := time.Now().UTC()

	user.UpdatedBy = userID
	user.UpdatedAt = now
	user.Status = -1

	return nil
}
