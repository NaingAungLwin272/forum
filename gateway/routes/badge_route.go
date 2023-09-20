package routes

import (
	"github.com/gin-gonic/gin"
	badge_controller "github.com/scm-dev1dev5/mtm-community-forum/gateway/controllers"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/middleware"
	clients "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	badge_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/badge"
)

func BadgeRoute(apiRouter *gin.RouterGroup, svc *clients.ServiceClient) {

	badgeService := badge_service.NewBadgeService(*svc)
	badgeController := badge_controller.NewBadgeController(badgeService)

	badgeRoute := apiRouter.Group("/")
	{
		badgeRoute.POST("/badge", middleware.VerifyToken(svc), badgeController.CreateBadge)
		badgeRoute.GET("/badge/:badgeId", middleware.VerifyToken(svc), badgeController.GetBadge)
		badgeRoute.GET("/badges", middleware.VerifyToken(svc), badgeController.GetBadgesList)
		badgeRoute.PUT("/badge/:badgeId", middleware.VerifyToken(svc), badgeController.UpdateBadge)
		badgeRoute.DELETE("/badge/:badgeId", middleware.VerifyToken(svc), badgeController.DeleteBadge)
	}

	userBadgeRoute := apiRouter.Group("/user_badge")
	{
		userBadgeRoute.POST("/", middleware.VerifyToken(svc), badgeController.CreateUserBadge)
	}

	userPointRoute := apiRouter.Group("/user_point")
	{
		userPointRoute.POST("/", middleware.VerifyToken(svc), badgeController.CreateUserPoint)
		userPointRoute.GET("/:pointId", middleware.VerifyToken(svc), badgeController.GetUserPoint)
		userPointRoute.POST("/points_list", middleware.VerifyToken(svc), badgeController.GetPointsList)
		userPointRoute.PUT("/:user_id", middleware.VerifyToken(svc), badgeController.UpdateUserPoint)
		userPointRoute.DELETE("/:pointId", middleware.VerifyToken(svc), badgeController.DeleteUserPoint)

	}
	apiRouter.GET("/users/:user_id/badges", middleware.VerifyToken(svc), badgeController.GetUserBadgeByUserId)
	apiRouter.GET("/users/evaluate_point", badgeController.EvaluatePoints)
}
