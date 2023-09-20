package mention_service

import "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/models"

type MentionService interface {
	CreateMention(*models.CreateMentionRequest) (*models.DBMention, error)
	GetMentions(page int, limit int) ([]*models.DBMention, error)
	GetMention(string) (*models.DBMention, error)
	UpdateMention(string, *models.UpdateMention) (*models.DBMention, error)
	DeleteMention(string) error
	GetMentionsByUserId(string, int, int) []*models.DBMention
	GetMentionCount(string) (count int64)
}
