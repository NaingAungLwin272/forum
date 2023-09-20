package gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/pb"
	services "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/services/user_badges_service"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserBadgeServer struct {
	pb.UnimplementedUserBadgeServiceServer
	userBadgeCollection *mongo.Collection
	userBadgeService    services.UserBadgeService
}

func NewGrpcUserBadgeServer(userBadgeCollection *mongo.Collection, userBadgeService services.UserBadgeService) (*UserBadgeServer, error) {
	userBadgeServer := &UserBadgeServer{
		userBadgeCollection: userBadgeCollection,
		userBadgeService:    userBadgeService,
	}

	return userBadgeServer, nil
}
