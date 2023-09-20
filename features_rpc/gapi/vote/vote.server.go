package vote_gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/pb"
	vote_service "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/services/vote"
	"go.mongodb.org/mongo-driver/mongo"
)

type VoteServer struct {
	pb.UnimplementedVoteServiceServer
	voteCollection *mongo.Collection
	voteService    vote_service.VoteService
}

func NewGrpcVoteServer(voteCollection *mongo.Collection, voteService vote_service.VoteService) (*VoteServer, error) {
	voteServer := &VoteServer{
		voteCollection: voteCollection,
		voteService:    voteService,
	}

	return voteServer, nil
}
