package services

import (
	"github.com/arezooq/open-utils/api"
	"github.com/arezooq/open-utils/errors"
	"github.com/arezooq/open-utils/logger"
	"github.com/arezooq/open-utils/security"
	"user-service/internal/helpers"
	"user-service/internal/models"
	"user-service/internal/repositories/postgres"
)

type userService struct {
	userRepo *postgres.UserRepository
	log      *logger.Logger
}

// NewUserService
func NewUserService(userRepo *postgres.UserRepository, log *logger.Logger) UserService {
	return &userService{userRepo: userRepo, log: log}
}

func (u *userService) CreateUser(req *api.Request, user *models.User) (*models.User, error) {

	existing, _ := u.userRepo.GetUserByEmailOrPhone(user.Email)
	if existing != nil {
		u.log.Warn("User already exist: " + user.Email)
		return nil, errors.ErrDuplicate
	}

	if err := helper.PreCreate(req.Ctx, user); err != nil {
		return nil, errors.ErrUnauthorized
	}

	hashedPassword, errPass := security.HashPassword(user.Password)
	if errPass != nil {
		u.log.Error("Failed to hash password for: " + errPass.Error())
		return nil, errors.ErrInternal
	}

	user.Password = hashedPassword

	user, err := u.userRepo.Create(user)

	if err != nil {
		u.log.Error("Failed to save user: " + err.Error())
		return nil, errors.ErrInternal
	}
	u.log.Info("Create new user successfully: " + user.Email)
	return user, nil
}

func (u *userService) GetAllUsers(req *api.Request, pagination *api.PaginationParams, query *api.QueryParams) ([]models.User, int64, error) {
	users, count, err := u.userRepo.GetAll(pagination, query.Filters, query.Search, query.Orders)
	if err != nil {
		u.log.Error("Failed to fetch users: " + err.Error())
		return nil, 0, errors.ErrInternal
	}

	u.log.Info("Fetched all users successfully")
	return users, count, nil
}

func (u *userService) GetUserById(req *api.Request, uuid string) (*models.User, error) {
	user, err := u.userRepo.GetById(uuid)
	if err != nil {
		u.log.Error("Failed to fetch user: " + err.Error())
		return nil, errors.ErrInternal
	}

	u.log.Info("Fetched user successfully")
	return user, nil
}

func (u *userService) UpdateUser(req *api.Request, uuid string, user *models.UpdateProfile) (*models.User, error) {
	
	if err := helper.PreUpdate(req.Ctx, user); err != nil {
		return nil, errors.ErrUnauthorized
	}

	_, errExist := u.userRepo.GetById(uuid)
	if errExist != nil {
		u.log.Warn("User not found by id: " + uuid)
		return nil, errors.ErrNotFound
	}

	updates := map[string]any{
		"f_name":     user.FName,
		"l_name":     user.LName,
		"username":   user.Username,
		"mobile":     user.Mobile,
		"email":      user.Email,
		"updated_by": user.UpdatedBy,
		"updated_at": user.UpdatedAt,
	}

	updateUser, err := u.userRepo.Update(uuid, updates)

	if err != nil {
		u.log.Error("Failed to update user: " + err.Error())
		return nil, errors.ErrInternal
	}
	u.log.Info("Updated user successfully: " + user.Email)
	return updateUser, nil
}

func (u *userService) DeleteUser(req *api.Request, uuid string) error {
	user, errExist := u.userRepo.GetById(uuid)
	if errExist != nil {
		u.log.Warn("User not found by id: " + uuid)
		return errors.ErrNotFound
	}

	if err := helper.PreDelete(req.Ctx, user); err != nil {
		return errors.ErrUnauthorized
	}

	err := u.userRepo.Delete(uuid)

	if err != nil {
		u.log.Error("Failed to delete user: " + err.Error())
		return errors.ErrInternal
	}
	u.log.Info("Deleted user successfully")
	return nil
}

func (u *userService) GetUserByEmailOrPhone(identifier string) (*models.User, error) {
	user, err := u.userRepo.GetUserByEmailOrPhone(identifier)
	if err != nil {
		u.log.Error("Failed to fetch user: " + err.Error())
		return nil, errors.ErrInternal
	}

	u.log.Info("Fetched user successfully")
	return user, nil
}
