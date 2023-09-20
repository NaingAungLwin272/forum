package vote_service

import (
	"github.com/gin-gonic/gin"
	vote_pb "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
)

type VoteServiceInterface interface {
	Create(ctx *gin.Context) (*vote_pb.VoteResponse, error)
	GetVotesByUserId(ctx *gin.Context) []*vote_pb.CommentResponse
	DeleteVote(ctx *gin.Context) (*vote_pb.DeleteVoteResponse, error)
	GetVotesByUserIdQuestionId(ctx *gin.Context) ([]*vote_pb.Vote, error)
}
