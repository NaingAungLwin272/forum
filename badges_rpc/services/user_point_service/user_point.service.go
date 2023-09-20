package services

import models "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/models"

type UserPointService interface {
	CreateUserPoint(*models.CreateUserPointRequest) (*models.DBUserPoint, error)
	GetUserPoint(string) (*models.DBUserPoint, error)
	GetUserPoints(page int, limit int) ([]*models.DBUserPoint, error)
	UpdateUserPoint(string, *models.UpdateUserPoint) (*models.DBUserPoint, error)
	DeleteUserPoint(string) error
	EvaluatePoints() error
}
