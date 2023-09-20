package services

import models "github.com/scm-dev1dev5/mtm-community-forum/user_rpc/models"

type DeparmentService interface {
	CreateDeparment(*models.CreateDeparmentRequest) (*models.DBDeparment, error)
	GetDeparment(string) (*models.DBDeparment, error)
	GetDeparmentList(page int, limit int) []*models.DBDeparment
	UpdateDeparment(string, *models.UpdateDeparment) (*models.DBDeparment, error)
	DeleteDeparment(string) error
	GetDepartmentCount(page int, limit int) (count int64)
}
