package badge_service

import (
	"github.com/gin-gonic/gin"
	badge_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/badge/pb"
)

type BadgeServiceInterface interface {
	//Badge Processes
	Create(ctx *gin.Context) (*badge_proto.BadgeResponse, error)
	GetBadge(ctx *gin.Context) (*badge_proto.BadgeResponse, error)
	GetBadgesList(ctx *gin.Context) ([]*badge_proto.Badge, error)
	UpdateBadge(ctx *gin.Context) (*badge_proto.BadgeResponse, error)
	DeleteBadge(ctx *gin.Context) (*badge_proto.DeleteBadgeResponse, error)

	//UserBadge Processes
	CreateUserBadge(ctx *gin.Context) (*badge_proto.UserBadgeResponse, error)
	GetUserBadgeByUserId(ctx *gin.Context) ([]*badge_proto.UserBadge, error)

	//UserPoint Processes
	CreateUserPoint(ctx *gin.Context) (*badge_proto.UserPointResponse, error)
	GetUserPoint(ctx *gin.Context) (*badge_proto.UserPointResponse, error)
	GetPointsList(ctx *gin.Context) ([]*badge_proto.UserPoint, error)
	UpdateUserPoint(ctx *gin.Context) *badge_proto.UserPointResponse
	DeleteUserPoint(ctx *gin.Context) (*badge_proto.DeleteUserPointResponse, error)
	EvaluatePoints(ctx *gin.Context) (*badge_proto.UserPointEvaluateResponse, error)
}
