package noti_service

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	noti_client "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	noti_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/noti"
	noti_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/noti/pb"
)

type NotiService struct {
	NotiClient noti_client.ServiceClient
	MailCient  noti_client.ServiceClient
}

// GetNoti implements NotiServiceInterface.
func (notiSvc *NotiService) GetNoti(ctx *gin.Context) (*noti_proto.NotiResponse, error) {
	data := noti_proxy.GetNoti(ctx, notiSvc.NotiClient.Noti)
	if data == nil {
		return nil, errors.New("failed to retrieve notification")
	}

	return data, nil
}

// GetNotis implements NotiServiceInterface.
func (notiSvc *NotiService) GetNotis(ctx *gin.Context) ([]*noti_proto.Noti, error) {
	data := noti_proxy.GetNotis(ctx, notiSvc.NotiClient.Noti)
	if data == nil {
		return nil, errors.New("failed to retrieve notifications")
	}

	return data, nil

}

// Create implements NotiServiceInterface.
func (notiSvc *NotiService) Create(ctx *gin.Context) ([]*noti_proto.NotiResponse, error) {
	data, err := noti_proxy.CreateNoti(ctx, notiSvc.NotiClient.Noti, notiSvc.NotiClient.Mail, notiSvc.NotiClient.User)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}
	return data, err
}

// UpdateNoti implements NotiServiceInterface.
func (notiSvc *NotiService) UpdateNoti(ctx *gin.Context) (*noti_proto.NotiResponse, error) {
	data := noti_proxy.UpdateNoti(ctx, notiSvc.NotiClient.Noti)
	if data == nil {
		return nil, errors.New("failed to update notification")
	}

	return data, nil
}

// DeleteNoti implements NotiServiceInterface.
func (notiSvc *NotiService) DeleteNoti(ctx *gin.Context) (*noti_proto.DeleteNotiResponse, error) {
	data := noti_proxy.DeleteNoti(ctx, notiSvc.NotiClient.Noti)
	if data == nil {
		return nil, errors.New("failed to delete notification")
	}

	return data, nil
}

// GetNotiByUserId implements NotiServiceInterface.
func (notiSvc *NotiService) GetNotiByUserId(ctx *gin.Context) []*noti_proto.Noti {
	data := noti_proxy.GetNotiByUserId(ctx, notiSvc.NotiClient.Noti)
	return data
}

// GetNotiCount implements NotiServiceInterface.
func (notiSvc *NotiService) GetNotiCount(ctx *gin.Context) (*noti_proto.NotiCountResponse, error) {
	data := noti_proxy.GetNotiCount(ctx, notiSvc.NotiClient.Noti)
	if data == nil {
		return nil, errors.New("failed to retrieve notifications")
	}

	return data, nil
}

// MarkAllNotiAsRead implements NotiServiceInterface.
func (notiSvc *NotiService) MarkAllNotiAsRead(ctx *gin.Context) (*noti_proto.DeleteNotiResponse, error) {
	data := noti_proxy.MarkAllNotiAsRead(ctx, notiSvc.NotiClient.Noti)
	if data == nil {
		return nil, errors.New("failed to retrieve notifications")
	}

	return data, nil
}

func NewNotiService(NotiClient noti_client.ServiceClient) NotiServiceInterface {
	return &NotiService{
		NotiClient: NotiClient,
	}
}
