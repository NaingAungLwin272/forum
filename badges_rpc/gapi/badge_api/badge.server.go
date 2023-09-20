package gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/pb"
	services "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/services/badge_service"
	"go.mongodb.org/mongo-driver/mongo"
)

type BadgeServer struct {
	pb.UnimplementedBadgeServiceServer
	badgeCollection *mongo.Collection
	badgeService    services.BadgeService
}

func NewGrpcBadgeServer(badgeCollection *mongo.Collection, badgeService services.BadgeService) (*BadgeServer, error) {
	badgeServer := &BadgeServer{
		badgeCollection: badgeCollection,
		badgeService:    badgeService,
	}

	return badgeServer, nil
}
