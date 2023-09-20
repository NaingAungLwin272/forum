package services

import models "github.com/scm-dev1dev5/mtm-community-forum/noti_rpc/models"

type NotiService interface {
	CreateNoti(*models.CreateNotiRequest) (*models.DBNoti, error)
	GetNoti(string) (*models.DBNoti, error)
	GetNotis(page int, limit int) ([]*models.DBNoti, error)
	UpdateNoti(string, *models.UpdateNoti) (*models.DBNoti, error)
	DeleteNoti(string) error
	GetNotiByUserId(string, int, int) []*models.DBNoti
	GetNotiCount(string) (count int64)
	MarkAllNotiAsRead(string) error
	GetNotiForUserSummary(string) (count int64)
}
