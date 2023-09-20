package comment_service

import "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/models"

type CommentService interface {
	CreateComment(*models.CreateCommentRequest) (*models.DBComment, error)
	GetComments(page int, limit int) ([]*models.DBComment, error)
	GetComment(string) (*models.DBComment, error)
	UpdateComment(string, *models.UpdateComment) (*models.DBComment, error)
	DeleteComment(string) error
	GetCommentByQuestionId(string) []*models.DBComment
	GetCommentsByUserId(string, int, int) ([]*models.DBComment, error)
	GetAnswersByUserId(string, int, int) ([]*models.DBComment, error)
	GetCommentsByUserIdWithSolved(string, int, int) ([]*models.DBComment, error)
	GetCommentCount(string) (count int64)
	GetCommentCountBySolved(string) (count int64)
	GetCommentCountByQuestionIdSolved(string) (count int64)
}
