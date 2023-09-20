package comment_service

import (
	"github.com/gin-gonic/gin"
	comment_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
)

type CommentServiceInterface interface {
	GetComment(ctx *gin.Context) (*comment_proto.CommentResponse, error)
	Create(ctx *gin.Context) (*comment_proto.CommentResponse, error)
	DeleteComment(ctx *gin.Context) (*comment_proto.DeleteCommentResponse, error)
	UpdateComment(ctx *gin.Context) (*comment_proto.CommentResponse, error)
	GetCommentByQuestionId(ctx *gin.Context) ([]*comment_proto.Comment, error)
	GetCommentByUserId(ctx *gin.Context) []*comment_proto.CommentResponse
	GetAnswersByUserId(ctx *gin.Context) []*comment_proto.CommentResponse
	GetCommentsByUserIdWithSolved(ctx *gin.Context) []*comment_proto.CommentResponse
}
