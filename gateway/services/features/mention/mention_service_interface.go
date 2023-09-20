package mention_service

import (
	"github.com/gin-gonic/gin"
	mention_pb "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
)

type MentionServiceInterface interface {
	CreateMention(ctx *gin.Context) (*mention_pb.MentionResponse, error)
	GetMentionsByUserId(ctx *gin.Context) []*mention_pb.CommentResponse
}
