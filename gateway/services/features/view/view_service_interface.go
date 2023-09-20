package view_service

import (
	"github.com/gin-gonic/gin"
	view_pb "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
)

type ViewServiceInterfae interface {
	Create(ctx *gin.Context) (*view_pb.ViewResponse, error)
	GetViewsByUserId(ctx *gin.Context) ([]*view_pb.View, error)
	// GetViewsByUserIdQuestionId(ctx *gin.Context) ([]*view_pb.View, error)
}
