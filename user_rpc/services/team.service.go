package services

import models "github.com/scm-dev1dev5/mtm-community-forum/user_rpc/models"

type TeamService interface {
	CreateTeam(*models.CreateTeamRequest) (*models.DBTeam, error)
	GetTeam(string) (*models.DBTeam, error)
	GetTeamList(page int, limit int) []*models.DBTeam
	UpdateTeam(string, *models.UpdateTeam) (*models.DBTeam, error)
	DeleteTeam(string) error
	GetTeamByDeparmentId(string) ([]*models.DBTeam, error)
	GetTeamCount(page int, limit int) (count int64)
}
