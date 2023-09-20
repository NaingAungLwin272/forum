package gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/noti_rpc/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/noti_rpc/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotiServer struct {
	pb.UnimplementedNotiServiceServer
	notiCollection *mongo.Collection
	notiService    services.NotiService
}

func NewGrpcNotiServer(notiCollection *mongo.Collection, notiService services.NotiService) (*NotiServer, error) {
	notiServer := &NotiServer{
		notiCollection: notiCollection,
		notiService:    notiService,
	}

	return notiServer, nil
}
