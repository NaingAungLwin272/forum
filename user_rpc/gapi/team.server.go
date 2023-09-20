package gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type TeamServer struct {
	pb.UnimplementedTeamServiceServer
	teamCollection *mongo.Collection
	teamService    services.TeamService
}

func NewGrpcTeamServer(teamCollection *mongo.Collection, teamService services.TeamService) (*TeamServer, error) {
	teamServer := &TeamServer{
		teamCollection: teamCollection,
		teamService:    teamService,
	}

	return teamServer, nil
}
