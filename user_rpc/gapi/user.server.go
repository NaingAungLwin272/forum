package gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	userCollection *mongo.Collection
	userService    services.UserService
}

func NewGrpcUserServer(userCollection *mongo.Collection, userService services.UserService) (*UserServer, error) {
	userServer := &UserServer{
		userCollection: userCollection,
		userService:    userService,
	}

	return userServer, nil
}
