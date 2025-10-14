package http_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	httphandler "user-service/internal/handlers/http"

	"github.com/arezooq/open-utils/api"
	"github.com/arezooq/open-utils/uuid"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"user-service/internal/models"
)

type MockUserService struct {
	CreateUserFn            func(req *api.Request, user *models.User) (*models.User, error)
	GetAllUsersFn           func(req *api.Request, pagination *api.PaginationParams, query *api.QueryParams) ([]models.User, int64, error)
	GetUserByIdFn           func(req *api.Request, uuid string) (*models.User, error)
	UpdateUserFn            func(req *api.Request, uuid string, user *models.UpdateProfile) (*models.User, error)
	DeleteUserFn            func(req *api.Request, uuid string) error
	GetUserByEmailOrPhoneFn func(identifier string) (*models.User, error)
}

// Implement all methods
func (m *MockUserService) CreateUser(req *api.Request, user *models.User) (*models.User, error) {
	return m.CreateUserFn(req, user)
}

func (m *MockUserService) GetAllUsers(req *api.Request, pagination *api.PaginationParams, query *api.QueryParams) ([]models.User, int64, error) {
	return m.GetAllUsersFn(req, pagination, query)
}

func (m *MockUserService) GetUserById(req *api.Request, uuid string) (*models.User, error) {
	return m.GetUserByIdFn(req, uuid)
}

func (m *MockUserService) UpdateUser(req *api.Request, uuid string, user *models.UpdateProfile) (*models.User, error) {
	return m.UpdateUserFn(req, uuid, user)
}

func (m *MockUserService) DeleteUser(req *api.Request, uuid string) error {
	return m.DeleteUserFn(req, uuid)
}

func (m *MockUserService) GetUserByEmailOrPhone(identifier string) (*models.User, error) {
	return m.GetUserByEmailOrPhoneFn(identifier)
}

// Create
func TestCreateUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockUserService{
		CreateUserFn: func(req *api.Request, user *models.User) (*models.User, error) {
			return &models.User{ID: uuid.UUIDString(), Email: user.Email}, nil
		},
	}

	h := httphandler.InitUserHandler(mockSvc)

	// JSON body
	body, _ := json.Marshal(models.User{Email: "test@example.com", Password: "pass"})
	w := httptest.NewRecorder()
	reqHTTP := httptest.NewRequest("POST", "/api/users", bytes.NewReader(body))
	reqHTTP.Header.Set("Content-Type", "application/json")

	r := gin.New()
	h.RegisterRoutes(r)
	r.ServeHTTP(w, reqHTTP)

	assert.Equal(t, 201, w.Code)
	assert.Contains(t, w.Body.String(), "test@example.com")
}

// Get all
func TestGetAllUsers_WithFiltersAndPagination_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockUserService{
		GetAllUsersFn: func(req *api.Request, p *api.PaginationParams, q *api.QueryParams) ([]models.User, int64, error) {
			users := []models.User{
				{ID: uuid.UUIDString(), Email: "alice@example.com", Username: "alice"},
				{ID: uuid.UUIDString(), Email: "bob@example.com", Username: "bob"},
				{ID: uuid.UUIDString(), Email: "charlie@example.com", Username: "charlie"},
			}

			for _, f := range q.Filters {
				if f.Field == "username" {
					filtered := []models.User{}
					for _, u := range users {
						if u.Username == f.Value {
							filtered = append(filtered, u)
						}
					}
					users = filtered
				}
			}

			total := int64(len(users))
			return users, total, nil
		},
	}

	h := httphandler.InitUserHandler(mockSvc)

	w := httptest.NewRecorder()
	reqHTTP := httptest.NewRequest(
		"GET",
		"/api/users?page=1&limit=2&username=alice&search=alice&order=email",
		nil,
	)
	reqHTTP.Header.Set("Content-Type", "application/json")

	r := gin.New()
	h.RegisterRoutes(r)
	r.ServeHTTP(w, reqHTTP)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "alice@example.com")
	assert.NotContains(t, w.Body.String(), "bob@example.com")
	assert.NotContains(t, w.Body.String(), "charlie@example.com")
}

// Update
func TestUpdateUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockUserService{
		UpdateUserFn: func(req *api.Request, uuid string, user *models.UpdateProfile) (*models.User, error) {
			return &models.User{ID: uuid, Email: user.Email}, nil
		},
	}

	h := httphandler.InitUserHandler(mockSvc)

	body, _ := json.Marshal(&models.UpdateProfile{FName: "John", LName: "Doe", Email: "john@example.com"})
	w := httptest.NewRecorder()
	reqHTTP := httptest.NewRequest("PUT", "/api/users/"+uuid.UUIDString(), bytes.NewReader(body))
	reqHTTP.Header.Set("Content-Type", "application/json")

	r := gin.New()
	h.RegisterRoutes(r)
	r.ServeHTTP(w, reqHTTP)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "john@example.com")
}

// Delete
func TestDeleteUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockUserService{
		DeleteUserFn: func(req *api.Request, uuid string) error {
			return nil
		},
	}

	h := httphandler.InitUserHandler(mockSvc)

	w := httptest.NewRecorder()
	reqHTTP := httptest.NewRequest("DELETE", "/api/users/"+uuid.UUIDString(), nil)
	reqHTTP.Header.Set("Content-Type", "application/json")

	r := gin.New()
	h.RegisterRoutes(r)
	r.ServeHTTP(w, reqHTTP)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "success")
}
