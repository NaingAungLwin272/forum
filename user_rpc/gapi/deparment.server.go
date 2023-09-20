package gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type DeparmentServer struct {
	pb.UnimplementedDeparmentServiceServer
	deparmentCollection *mongo.Collection
	deparmentService    services.DeparmentService
}

func NewGrpcDeparmentServer(deparmentCollection *mongo.Collection, deparmentService services.DeparmentService) (*DeparmentServer, error) {
	deparmentServer := &DeparmentServer{
		deparmentCollection: deparmentCollection,
		deparmentService:    deparmentService,
	}

	return deparmentServer, nil
}
