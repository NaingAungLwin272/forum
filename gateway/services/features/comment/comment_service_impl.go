package comment_service

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	comment_client "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	comment_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	comment_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/proxy"
)

type CommentService struct {
	CommentClient  comment_client.ServiceClient
	BookmarkClient comment_client.ServiceClient
}

// GetComment implements CommentServiceInterface.
func (commentSvc *CommentService) GetComment(ctx *gin.Context) (*comment_proto.CommentResponse, error) {
	commentID := ctx.Param("comment_id")
	data := comment_proxy.GetComment(commentSvc.CommentClient.Comment, commentID)
	if data == nil {
		return nil, errors.New("failed to retrieve comment")
	}

	return data, nil
}

// Create implements CommentServiceInterface.
func (commentSvc *CommentService) Create(ctx *gin.Context) (*comment_proto.CommentResponse, error) {
	data, err := comment_proxy.CreateComment(ctx, commentSvc.CommentClient.Comment, commentSvc.CommentClient.Question, commentSvc.CommentClient.Noti, commentSvc.CommentClient.Mail, commentSvc.CommentClient.User, commentSvc.CommentClient.UserPoint)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}
	return data, err
}

// UpdateComment implements CommentServiceInterface.
func (commentSvc *CommentService) UpdateComment(ctx *gin.Context) (*comment_proto.CommentResponse, error) {
	data := comment_proxy.UpdateComment(ctx, commentSvc.CommentClient.Comment, commentSvc.CommentClient.Noti, commentSvc.CommentClient.User, commentSvc.CommentClient.Mail, commentSvc.CommentClient.Question)
	if data == nil {
		return nil, errors.New("failed to update comment")
	}
	return data, nil
}

// DeleteComment implements CommentServiceInterface.
func (commentSvc *CommentService) DeleteComment(ctx *gin.Context) (*comment_proto.DeleteCommentResponse, error) {
	data := comment_proxy.DeleteComment(ctx, commentSvc.CommentClient.Comment, commentSvc.CommentClient.Question)
	if data == nil {
		return nil, errors.New("failed to delete comment")
	}
	return data, nil
}

// GetCommentByQuestionId implements CommentServiceInterface.
func (commentSvc *CommentService) GetCommentByQuestionId(ctx *gin.Context) ([]*comment_proto.Comment, error) {
	questionID := ctx.Param("question_id")
	data := comment_proxy.GetCommentByQuestionId(ctx, commentSvc.CommentClient.Comment, questionID)
	if data == nil {
		return nil, errors.New("failed to retrieve comments")
	}

	return data, nil
}

func (commentSvc *CommentService) GetCommentByUserId(ctx *gin.Context) []*comment_proto.CommentResponse {
	data := comment_proxy.GetCommentsByUserId(ctx, commentSvc.CommentClient.Comment, commentSvc.CommentClient.User)
	return data
}

// GetAnswersByUserId implements CommentServiceInterface.
func (commentSvc *CommentService) GetAnswersByUserId(ctx *gin.Context) []*comment_proto.CommentResponse {
	data := comment_proxy.GetAnswersByUserId(ctx, commentSvc.CommentClient.Comment, commentSvc.CommentClient.User)
	return data
}

func (commentSvc *CommentService) GetCommentsByUserIdWithSolved(ctx *gin.Context) []*comment_proto.CommentResponse {
	data := comment_proxy.GetCommentsByUserIdWithSolved(ctx, commentSvc.CommentClient.Comment, commentSvc.CommentClient.User)
	return data
}

func NewCommentService(CommentClient comment_client.ServiceClient) CommentServiceInterface {
	return &CommentService{
		CommentClient: CommentClient,
	}
}
