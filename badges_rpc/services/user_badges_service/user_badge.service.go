package services

import models "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/models"

type UserBadgeService interface {
	CreateUserBadge(*models.CreateUserBadgeRequest) (*models.DBUserBadge, error)
	GetUserBadge(string, string) (*models.DBUserBadge, error)
	GetUserBadges(page int, limit int) ([]*models.DBUserBadge, error)
	GetUserBadgesOfUser(userId string) ([]*models.DBUserBadge, error)
	UpdateUserBadge(string, *models.UpdateUserBadge) (*models.DBUserBadge, error)
	DeleteUserBadge(string) error
	GetBadgeCount(string) (count int64)
}
