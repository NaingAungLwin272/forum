package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/scm-dev1dev5/mtm-community-forum/gateway/controllers"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/middleware"
	clients "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	mail_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/mail"
)

func MailRoute(apiRouter *gin.RouterGroup, svc *clients.ServiceClient) {

	mailService := mail_service.NewMailService(*svc)
	mailController := controllers.NewMailController(mailService)
	apiRouter.POST("/email/send-message", middleware.VerifyToken(svc), mailController.SendMail)
	apiRouter.POST("/email/forget-password", middleware.VerifyToken(svc), mailController.ForgetPasswordMail)
}
