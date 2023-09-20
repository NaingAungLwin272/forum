package mail_service

import (
	"github.com/gin-gonic/gin"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/mail/pb"
)

type MailServiceInterface interface {
	SendMail(ctx *gin.Context) (*pb.MailResponse, error)
	ForgetPasswordMail(ctx *gin.Context) (*pb.MailResponse, error)
}
