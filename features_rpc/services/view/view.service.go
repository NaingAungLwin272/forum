package view_service

import "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/models"

type ViewService interface {
	CreateView(*models.CreateViewRequest) (*models.DBView, error)
	GetViews(page int, limit int) ([]*models.DBView, error)
	GetView(string) (*models.DBView, error)
	UpdateView(string, *models.UpdateView) (*models.DBView, error)
	DeleteView(string) error
	GetViewsByUserId(string) ([]*models.DBView, error)
	GetViewsByUserIdQuestionId(string, string, int, int) ([]*models.DBView, error)
}
