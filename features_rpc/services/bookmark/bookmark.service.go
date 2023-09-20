package bookmark_service

import "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/models"

type BookmarkService interface {
	CreateBookmark(*models.CreateBookmarkRequest) (*models.DBBookmark, error)
	GetBookmarks(page int, limit int) ([]*models.DBBookmark, error)
	GetBookmark(string) (*models.DBBookmark, error)
	UpdateBookmark(string, *models.UpdateBookmark) (*models.DBBookmark, error)
	DeleteBookmark(string) error
	GetBookmarksByUserId(string, int, int) ([]*models.DBBookmark, error)
	GetBookmarkCount(string) (count int64)
	GetBookmarksByUserIdQuestionId(string, string, int, int) ([]*models.DBBookmark, error)
}
