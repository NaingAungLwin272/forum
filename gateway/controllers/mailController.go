package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mail_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/mail"
)

type MailController struct {
	MailServiceInterface mail_service.MailServiceInterface
}

func (controller *MailController) SendMail(ctx *gin.Context) {

	data, err := controller.MailServiceInterface.SendMail(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *MailController) ForgetPasswordMail(ctx *gin.Context) {

	data, err := controller.MailServiceInterface.ForgetPasswordMail(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}


func NewMailController(MailServiceInterface mail_service.MailServiceInterface) *MailController {
	return &MailController{
		MailServiceInterface: MailServiceInterface,
	}
}
