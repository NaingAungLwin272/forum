package noti_service

import (
	"github.com/gin-gonic/gin"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/noti/pb"
)

type NotiServiceInterface interface {
	Create(ctx *gin.Context) ([]*pb.NotiResponse, error)
	GetNotis(ctx *gin.Context) ([]*pb.Noti, error)
	GetNoti(ctx *gin.Context) (*pb.NotiResponse, error)
	UpdateNoti(ctx *gin.Context) (*pb.NotiResponse, error)
	DeleteNoti(ctx *gin.Context) (*pb.DeleteNotiResponse, error)
	GetNotiByUserId(ctx *gin.Context) []*pb.Noti
	GetNotiCount(ctx *gin.Context) (*pb.NotiCountResponse, error)
	MarkAllNotiAsRead(ctx *gin.Context) (*pb.DeleteNotiResponse, error)
}
