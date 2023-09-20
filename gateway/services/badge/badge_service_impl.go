package badge_service

import (
	"errors"

	"github.com/gin-gonic/gin"
	badge_client "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	badge_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/badge"
	badge_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/badge/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

type BadgeService struct {
	BadgeClient badge_client.ServiceClient
}

// Badge Processes
func (BadgeSvc *BadgeService) Create(ctx *gin.Context) (*badge_proto.BadgeResponse, error) {
	data, err := badge_proxy.CreateBadge(ctx, BadgeSvc.BadgeClient.Badge)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (BadgeSvc *BadgeService) GetBadge(ctx *gin.Context) (*badge_proto.BadgeResponse, error) {
	data, err := badge_proxy.GetBadge(ctx, BadgeSvc.BadgeClient.Badge)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (BadgeSvc *BadgeService) GetBadgesList(ctx *gin.Context) ([]*badge_proto.Badge, error) {
	data, err := badge_proxy.GetBadgeLists(ctx, BadgeSvc.BadgeClient.Badge)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (BadgeSvc *BadgeService) UpdateBadge(ctx *gin.Context) (*badge_proto.BadgeResponse, error) {
	data, err := badge_proxy.UpdateBadge(ctx, BadgeSvc.BadgeClient.Badge)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (BadgeSvc *BadgeService) DeleteBadge(ctx *gin.Context) (*badge_proto.DeleteBadgeResponse, error) {
	data, err := badge_proxy.DeleteBadge(ctx, BadgeSvc.BadgeClient.Badge)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

// UserBadge Processes
func (BadgeSvc *BadgeService) CreateUserBadge(ctx *gin.Context) (*badge_proto.UserBadgeResponse, error) {
	data, err := badge_proxy.CreateUserBadge(ctx, BadgeSvc.BadgeClient.UserBadge)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (BadgeSvc *BadgeService) GetUserBadgeByUserId(ctx *gin.Context) ([]*badge_proto.UserBadge, error) {
	data, err := badge_proxy.GetUserBadgeByUserId(ctx, BadgeSvc.BadgeClient.UserBadge)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

// UserPoint Processes
func (BadgeSvc *BadgeService) CreateUserPoint(ctx *gin.Context) (*badge_proto.UserPointResponse, error) {
	data, err := badge_proxy.CreateUserPoint(ctx, BadgeSvc.BadgeClient.UserPoint)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (BadgeSvc *BadgeService) GetUserPoint(ctx *gin.Context) (*badge_proto.UserPointResponse, error) {
	data, err := badge_proxy.GetUserPoint(ctx, BadgeSvc.BadgeClient.UserPoint)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (BadgeSvc *BadgeService) GetPointsList(ctx *gin.Context) ([]*badge_proto.UserPoint, error) {
	data, err := badge_proxy.GetPointsList(ctx, BadgeSvc.BadgeClient.UserPoint)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (BadgeSvc *BadgeService) UpdateUserPoint(ctx *gin.Context) *badge_proto.UserPointResponse {
	data := badge_proxy.UpdateUserPoint(ctx, BadgeSvc.BadgeClient.UserPoint)
	return data
}

func (BadgeSvc *BadgeService) DeleteUserPoint(ctx *gin.Context) (*badge_proto.DeleteUserPointResponse, error) {
	data, err := badge_proxy.DeleteUserPoint(ctx, BadgeSvc.BadgeClient.UserPoint)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (BadgeSvc *BadgeService) EvaluatePoints(ctx *gin.Context) (*badge_proto.UserPointEvaluateResponse, error) {
	data, err := badge_proxy.EvaluatePoints(ctx, BadgeSvc.BadgeClient.UserPoint)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func NewBadgeService(BadgeClient badge_client.ServiceClient) BadgeServiceInterface {
	return &BadgeService{
		BadgeClient: BadgeClient,
	}
}
