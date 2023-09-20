package services

import (
	"context"

	models "github.com/scm-dev1dev5/mtm-community-forum/user_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/pb"
	// "github.com/scm-dev1dev5/mtm-community-forum/user_rpc/pb"
)

type UserService interface {
	CreateUser(*models.CreateUserRequest) (*models.DBUser, error)
	GetUser(string) (*models.DBUser, error)
	GetUsers(page int, limit int) ([]*models.DBUser, error)
	UpdateUser(string, *models.UpdateUser, context.Context) (*models.DBUser, error)
	DeleteUser(string) error
	FilterUser(*models.FilterUserRequest) ([]*models.DBUser, error)
	CreateUsersWithCsv([]*models.CreateUserRequest) ([]*models.DBUser, error)
	UploadImage(*pb.FileUploadRequest) (*pb.FileUploadResponse, error)
	GetUserCount(*models.FilterUserRequest) (count int64)
	GetUserByDisplayName(string) (*models.DBUser, error)
	// GetUserNotiCount(string) (count int64)
}
