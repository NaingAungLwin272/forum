package utils

import (
	"context"
	"errors"
	"log"

	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/auth"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/metadata"
)

func ToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

func CheckAuth(ctx context.Context, jwtManager *auth.JWTManager) (*auth.UserClaims, error) {

	metaData, ok := metadata.FromIncomingContext(ctx)
	values := metaData["authorization"]
	if !ok {
		log.Println("--> metadata is not provided")
		return nil, errors.New("metadata is not provided")
	}

	values = metaData["authorization"]
	if len(values) == 0 {
		log.Println("--> authorization token is not provided")
		return nil, errors.New("authorization token is not provided")
	}

	accessToken := values[0]

	log.Println(accessToken, "accessToken is here")

	claims, ok, err := jwtManager.Verify(accessToken)

	if err != nil {
		log.Println("--> access token is invalid")
		return nil, errors.New("access token is invalid")
	}

	return claims, nil
}