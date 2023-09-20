package mention_service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	clients "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	mention_pb "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	features_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/proxy"
	"go.mongodb.org/mongo-driver/mongo"
)

type MentionService struct {
	MentionClient clients.ServiceClient
}

// GetMentionsByUserId implements MentionServiceInterface
func (MentionSvc *MentionService) GetMentionsByUserId(ctx *gin.Context) []*mention_pb.CommentResponse {
	data := features_proxy.GetMentionsByUserId(ctx, MentionSvc.MentionClient.Mention, MentionSvc.MentionClient.Comment, MentionSvc.MentionClient.User)
	fmt.Println(data, "data.........")
	return data
}

// Create implements UserServiceInterface.
func (MentionSvc *MentionService) CreateMention(ctx *gin.Context) (*mention_pb.MentionResponse, error) {
	data, err := features_proxy.CreateMention(ctx, MentionSvc.MentionClient.Mention, MentionSvc.MentionClient.Noti, MentionSvc.MentionClient.User, MentionSvc.MentionClient.Mail, MentionSvc.MentionClient.Question, MentionSvc.MentionClient.Comment)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func NewMentionService(MentionClient clients.ServiceClient) MentionServiceInterface {
	return &MentionService{
		MentionClient: MentionClient,
	}
}
