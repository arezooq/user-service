package helper

import (
	"github.com/arezooq/open-utils/api"
	"time"
	"user-service/internal/models"

	"github.com/arezooq/open-utils/jwt"
	"github.com/arezooq/open-utils/uuid"
)

type AuditFields struct {
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
}

// Pre Create
func PreCreate(req *api.Request, user *models.User) error {
	token, err := jwt.ExtractTokenFromHeader(req)

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
func PreUpdate(req *api.Request, user *models.UpdateProfile) error {
	token, err := jwt.ExtractTokenFromHeader(req)

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
func PreDelete(req *api.Request, user *models.User) error {
	token, err := jwt.ExtractTokenFromHeader(req)

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
