package view_gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/pb"
	view_service "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/services/view"
	"go.mongodb.org/mongo-driver/mongo"
)

type ViewServer struct {
	pb.UnimplementedViewServiceServer
	viewCollection *mongo.Collection
	viewService    view_service.ViewService
}

func NewGrpcViewServer(viewCollection *mongo.Collection, viewService view_service.ViewService) (*ViewServer, error) {
	viewServer := &ViewServer{
		viewCollection: viewCollection,
		viewService:    viewService,
	}

	return viewServer, nil
}
