package gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/pb"
	services "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/services/user_point_service"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserPointServer struct {
	pb.UnimplementedUserPointServiceServer
	UserPointCollection *mongo.Collection
	UserPointService    services.UserPointService
}

func NewGrpcUserPointServer(UserPointCollection *mongo.Collection, UserPointService services.UserPointService) (*UserPointServer, error) {
	UserPointServer := &UserPointServer{
		UserPointCollection: UserPointCollection,
		UserPointService:    UserPointService,
	}

	return UserPointServer, nil
}
