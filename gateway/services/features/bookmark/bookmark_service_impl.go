package bookmark_service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	clients "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	bookmark_pb "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	bookmark_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	features_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/proxy"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookmarkService struct {
	BookmarkClient clients.ServiceClient
}
type CustomCommentResponse struct {
	*bookmark_proto.CommentResponse
	UserProfile string `json:"user_profile"`
}

// GetBookmarksByUserId implements BookmarkServiceInterface
func (BookmarkSvc *BookmarkService) GetBookmarksByUserId(ctx *gin.Context) []*bookmark_proto.CommentResponse {
	data := features_proxy.GetBookmarksByUserId(ctx, BookmarkSvc.BookmarkClient.BookMark, BookmarkSvc.BookmarkClient.Comment, BookmarkSvc.BookmarkClient.User) // Pass CommentServiceClient argument
	fmt.Println(data, "data........")
	return data
}

// Create implements UserServiceInterface.
func (BookmarkSvc *BookmarkService) CreateBookmark(ctx *gin.Context) (*bookmark_pb.BookmarkResponse, error) {
	data, err := features_proxy.CreateBookmark(ctx, BookmarkSvc.BookmarkClient.BookMark)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

// DeleteBookmark implements bookmarkServiceInterface.
func (BookmarkSvc *BookmarkService) DeleteBookmark(ctx *gin.Context) (*bookmark_pb.DeleteBookmarkResponse, error) {
	data, err := features_proxy.DeleteBookmark(ctx, BookmarkSvc.BookmarkClient.BookMark)
	if data == nil {
		return nil, errors.New("failed to delete bookmark")
	}
	return data, err
}

func (BookmarkSvc *BookmarkService) GetBookmarksByUserIdQuestionId(ctx *gin.Context) ([]*bookmark_pb.Bookmark, error) {
	data := features_proxy.GetBookmarksByUserIdQuestionId(ctx, BookmarkSvc.BookmarkClient.BookMark)
	return data, nil
}

func NewBookmarkService(BookmarkClient clients.ServiceClient) BookmarkServiceInterface {
	return &BookmarkService{
		BookmarkClient: BookmarkClient,
	}
}
