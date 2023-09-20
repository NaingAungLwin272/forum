package routes

import (
	"github.com/gin-gonic/gin"
	user_controllers "github.com/scm-dev1dev5/mtm-community-forum/gateway/controllers"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/middleware"
	service "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	user_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/user"
)

func UserRoute(apiRouter *gin.RouterGroup, svc *service.ServiceClient) {
	userService := user_service.NewUserService(*svc)
	userController := user_controllers.NewUserController(userService)

	authRoute := apiRouter.Group("/auth")
	{
		authRoute.POST("/login", userController.Login)
		authRoute.POST("/forget-password", userController.ForgetPassword)
		authRoute.POST("/reset-password", userController.ResetPassword)

	}

	userRoute := apiRouter.Group("/")
	{
		userRoute.POST("/user", middleware.VerifyToken(svc), userController.CreateUser)
		userRoute.GET("/users", middleware.VerifyToken(svc), userController.GetUsers)
		userRoute.GET("/uname/:displayname", middleware.VerifyToken(svc), userController.GetUserByDisplayName)
		userRoute.GET("/users/count", middleware.VerifyToken(svc), userController.GetUserCount)
		userRoute.GET("users/:user_id", middleware.VerifyToken(svc), userController.GetUser)
		userRoute.PUT("users/:user_id", middleware.VerifyToken(svc), userController.UpdateUser)
		userRoute.DELETE("users/:user_id", middleware.VerifyToken(svc), userController.DeleteUser)
		userRoute.GET("users/:user_id/summary", middleware.VerifyToken(svc), userController.GetUserSummary)
		userRoute.POST("users/csv", middleware.VerifyToken(svc), userController.CreateUserWithCsv)
		userRoute.POST("users/search", middleware.VerifyToken(svc), userController.FilterUser)
		userRoute.POST("/users/count", middleware.VerifyToken(svc), userController.GetUserCount)
		userRoute.PUT("user/change-password", middleware.VerifyToken(svc), userController.ChangePassword)
	}

	teamRoute := apiRouter.Group("/")
	{
		teamRoute.POST("/team", middleware.VerifyToken(svc), userController.CreateTeam)
		teamRoute.GET("/teams", middleware.VerifyToken(svc), userController.GetTeams)
		teamRoute.GET("/teams/count", middleware.VerifyToken(svc), userController.GetTeamCount)
		teamRoute.GET("teams/:team_id", middleware.VerifyToken(svc), userController.GetTeam)
		teamRoute.PUT("teams/:team_id", middleware.VerifyToken(svc), userController.UpdateTeam)
		teamRoute.DELETE("teams/:team_id", middleware.VerifyToken(svc), userController.DeleteTeam)
		teamRoute.GET("/departments/:department_id/teams", middleware.VerifyToken(svc), userController.GetTeamsByDepartmentId)
	}

	departmentRoute := apiRouter.Group("/")
	{
		departmentRoute.POST("/department", middleware.VerifyToken(svc), userController.CreateDepartment)
		departmentRoute.GET("/departments", middleware.VerifyToken(svc), userController.GetDepartments)
		departmentRoute.GET("/departments/count", middleware.VerifyToken(svc), userController.GetDepartmentCount)
		departmentRoute.PUT("departments/:department_id", middleware.VerifyToken(svc), userController.UpdateDepartment)
		departmentRoute.GET("departments/:department_id", middleware.VerifyToken(svc), userController.GetDepartment)
		departmentRoute.DELETE("departments/:department_id", middleware.VerifyToken(svc), userController.DeleteDepartment)
	}
}
