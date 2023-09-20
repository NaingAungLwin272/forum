package vote_service

import "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/models"

type VoteService interface {
	CreateVote(*models.CreateVoteRequest) (*models.DBVote, error)
	GetVotes(page int, limit int) ([]*models.DBVote, error)
	GetVote(string) (*models.DBVote, error)
	UpdateVote(string, *models.UpdateVote) (*models.DBVote, error)
	DeleteVote(string) error
	GetVotesByUserId(string, int, int) ([]*models.DBVote, error)
	GetVoteCount(string) (count int64)
	GetVotesByUserIdQuestionId(string, string, int, int) ([]*models.DBVote, error)
}
