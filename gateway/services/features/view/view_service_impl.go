package view_service

import (
	"errors"

	"github.com/gin-gonic/gin"
	clients "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	view_pb "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	features_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/proxy"
	"go.mongodb.org/mongo-driver/mongo"
)

type ViewService struct {
	ViewClient clients.ServiceClient
}

// GetViewsByUserId implements ViewServiceInterface
func (ViewSvc *ViewService) GetViewsByUserId(ctx *gin.Context) ([]*view_pb.View, error) {
	data := features_proxy.GetViewsUserId(ctx, ViewSvc.ViewClient.View)
	if data == nil {
		return nil, errors.New("no view with that userid doesn`t exists")
	}
	return data, nil
}

// Create implements UserServiceInterface.
func (ViewSvc *ViewService) Create(ctx *gin.Context) (*view_pb.ViewResponse, error) {
	data, err := features_proxy.CreateView(ctx, ViewSvc.ViewClient.View, ViewSvc.ViewClient.Question)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

// func (ViewSvc *ViewService) GetViewsByUserIdQuestionId(ctx *gin.Context) ([]*view_pb.View, error) {
// 	data := features_proxy.GetViewsByUserIdQuestionId(ctx, ViewSvc.ViewClient.View)
// 	return data, nil
// }

func NewViewService(ViewClient clients.ServiceClient) ViewServiceInterfae {
	return &ViewService{
		ViewClient: ViewClient,
	}
}
