package mention_gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/pb"
	mention_service "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/services/mention"
	"go.mongodb.org/mongo-driver/mongo"
)

type MentionServer struct {
	pb.UnimplementedMentionServiceServer
	mentionCollection *mongo.Collection
	mentionService    mention_service.MentionService
}

func NewGrpcMentionServer(mentionCollection *mongo.Collection, mentionService mention_service.MentionService) (*MentionServer, error) {
	mentionServer := &MentionServer{
		mentionCollection: mentionCollection,
		mentionService:    mentionService,
	}

	return mentionServer, nil
}
