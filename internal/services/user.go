package services

import (
	"github.com/arezooq/open-utils/api"
	"user-service/internal/models"
)

type UserService interface {
	CreateUser(req *api.Request, user *models.User) (*models.User, error)
	GetAllUsers(req *api.Request, pagination *api.PaginationParams, query *api.QueryParams) ([]models.User, int64, error)
	GetUserById(req *api.Request, uuid string) (*models.User, error)
	UpdateUser(req *api.Request, uuid string, user *models.UpdateProfile) (*models.User, error)
	DeleteUser(req *api.Request, uuid string) error
	GetUserByEmailOrPhone(identifier string) (*models.User, error)
}
