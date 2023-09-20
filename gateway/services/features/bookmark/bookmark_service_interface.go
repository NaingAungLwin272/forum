package bookmark_service

import (
	"github.com/gin-gonic/gin"
	bookmark_pb "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
)

// type CustomCommentResponse struct {
// 	*bookmark_proto.CommentResponse
// 	UserProfile string `json:"user_profile"`
// }

type BookmarkServiceInterface interface {
	CreateBookmark(ctx *gin.Context) (*bookmark_pb.BookmarkResponse, error)
	GetBookmarksByUserId(ctx *gin.Context) []*bookmark_pb.CommentResponse
	DeleteBookmark(ctx *gin.Context) (*bookmark_pb.DeleteBookmarkResponse, error)
	GetBookmarksByUserIdQuestionId(ctx *gin.Context) ([]*bookmark_pb.Bookmark, error)
}
