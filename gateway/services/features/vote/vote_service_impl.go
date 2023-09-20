package vote_service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	clients "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	vote_pb "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	features_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/proxy"
	"go.mongodb.org/mongo-driver/mongo"
)

type VoteService struct {
	VoteClient     clients.ServiceClient
	BookmarkClient clients.ServiceClient
}

// GetVotesByUserId implements VoteServiceInterface
func (VoteSvc *VoteService) GetVotesByUserId(ctx *gin.Context) []*vote_pb.CommentResponse {
	data := features_proxy.GetVotesByUserId(ctx, VoteSvc.VoteClient.Vote, VoteSvc.VoteClient.Comment, VoteSvc.VoteClient.User)
	return data
}

// Create implements UserServiceInterface.
func (VoteSvc *VoteService) Create(ctx *gin.Context) (*vote_pb.VoteResponse, error) {
	data, err := features_proxy.CreateVote(ctx, VoteSvc.VoteClient.Vote, VoteSvc.VoteClient.Question, VoteSvc.VoteClient.Comment, VoteSvc.VoteClient.Noti, VoteSvc.VoteClient.User, VoteSvc.VoteClient.Mail)
	fmt.Println(err)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

// DeleteVote implements voteServiceInterface.
func (VoteSvc *VoteService) DeleteVote(ctx *gin.Context) (*vote_pb.DeleteVoteResponse, error) {
	data, err := features_proxy.DeleteVote(ctx, VoteSvc.VoteClient.Vote, VoteSvc.VoteClient.Question, VoteSvc.VoteClient.Comment)
	if data == nil {
		return nil, errors.New("failed to delete comment")
	}
	return data, err
}

func (VoteSvc *VoteService) GetVotesByUserIdQuestionId(ctx *gin.Context) ([]*vote_pb.Vote, error) {
	data := features_proxy.GetVotesByUserIdQuestionId(ctx, VoteSvc.VoteClient.Vote)
	return data, nil
}

func NewVoteService(VoteClient clients.ServiceClient) VoteServiceInterface {
	return &VoteService{
		VoteClient: VoteClient,
	}
}
