package auth_gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/auth"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userCollection *mongo.Collection
	jwtManager     *auth.JWTManager
	forgetPassword *auth.JWTManager
	changePassword *auth.JWTManager
}

func NewGrpcAuthServer(userCollection *mongo.Collection, jwtManager *auth.JWTManager, forgetPassword *auth.JWTManager, changePassword *auth.JWTManager) (*AuthServer, error) {
	authServer := &AuthServer{
		userCollection: userCollection,
		jwtManager:     jwtManager,
		forgetPassword: forgetPassword,
		changePassword: changePassword,
	}

	return authServer, nil
}
