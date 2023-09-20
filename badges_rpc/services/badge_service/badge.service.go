package services

import models "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/models"

type BadgeService interface {
	CreateBadge(*models.CreateBadgeRequest) (*models.DBBadge, error)
	GetBadge(string) (*models.DBBadge, error)
	GetBadges(page int, limit int) ([]*models.DBBadge, error)
	UpdateBadge(string, *models.UpdateBadge) (*models.DBBadge, error)
	DeleteBadge(string) error
}
