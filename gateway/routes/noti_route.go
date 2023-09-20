package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/scm-dev1dev5/mtm-community-forum/gateway/controllers"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/middleware"
	clients "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	noti_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/noti"
)

func NotiRoute(apiRouter *gin.RouterGroup, svc *clients.ServiceClient) {
	notiService := noti_service.NewNotiService(*svc)
	notiController := controllers.NewNotiController(notiService)
	apiRouter.GET("/notifications", middleware.VerifyToken(svc), notiController.GetNotis)
	apiRouter.GET("/user/:id/notifications", middleware.VerifyToken(svc), notiController.GetNotiByUserId)
	apiRouter.GET("/user/:id/noticount", middleware.VerifyToken(svc), notiController.GetNotiCount)
	apiRouter.PUT("/user/:id/notireset", middleware.VerifyToken(svc), notiController.MarkAllNotiAsRead)
	notiRoute := apiRouter.Group("/notification")
	{
		notiRoute.POST("", middleware.VerifyToken(svc), notiController.CreateNoti)
		notiRoute.GET("/:id", middleware.VerifyToken(svc), notiController.GetNoti)
		notiRoute.PUT("/:id", middleware.VerifyToken(svc), notiController.UpdateNoti)
		notiRoute.DELETE("/:id", middleware.VerifyToken(svc), notiController.DeleteNoti)
	}
}
