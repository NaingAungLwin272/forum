package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	noti_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/noti"
)

type NotiController struct {
	NotiServiceInterface noti_service.NotiServiceInterface
}

func (controller *NotiController) CreateNoti(ctx *gin.Context) {

	data, err := controller.NotiServiceInterface.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *NotiController) GetNotis(ctx *gin.Context) {
	data, err := controller.NotiServiceInterface.GetNotis(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, data)
	}

}

func (controller *NotiController) GetNoti(ctx *gin.Context) {
	data, err := controller.NotiServiceInterface.GetNoti(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, data)
	}

}

func (controller *NotiController) UpdateNoti(ctx *gin.Context) {
	data, err := controller.NotiServiceInterface.UpdateNoti(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, data)
	}
}

func (controller *NotiController) DeleteNoti(ctx *gin.Context) {
	data, err := controller.NotiServiceInterface.DeleteNoti(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, data)
	}

}

func (controller *NotiController) GetNotiByUserId(ctx *gin.Context) {
	data := controller.NotiServiceInterface.GetNotiByUserId(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (controller *NotiController) GetNotiCount(ctx *gin.Context) {
	data, err := controller.NotiServiceInterface.GetNotiCount(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, data)
	}
}

func (controller *NotiController) MarkAllNotiAsRead(ctx *gin.Context) {
	data, err := controller.NotiServiceInterface.MarkAllNotiAsRead(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, data)
	}
}

func NewNotiController(NotiServiceInterface noti_service.NotiServiceInterface) *NotiController {
	return &NotiController{
		NotiServiceInterface: NotiServiceInterface,
	}
}
