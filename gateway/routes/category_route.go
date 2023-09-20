package routes

import (
	"github.com/gin-gonic/gin"
	category_controllers "github.com/scm-dev1dev5/mtm-community-forum/gateway/controllers"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/middleware"
	service "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	category_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/category"
)

func CategoryRoute(apiRouter *gin.RouterGroup, svc *service.ServiceClient) {

	categoryService := category_service.NewCategoryService(*svc)
	categoryController := category_controllers.NewCategoryController(categoryService)

	categoryRoute := apiRouter.Group("/")
	{
		categoryRoute.POST("/category", middleware.VerifyToken(svc), categoryController.CreateCategory)
		categoryRoute.GET("/categories", middleware.VerifyToken(svc), categoryController.GetCategories)
		categoryRoute.GET("/categories/count", middleware.VerifyToken(svc), categoryController.GetCategoryCount)
		categoryRoute.GET("/categories/type", middleware.VerifyToken(svc), categoryController.GetCategoryByType)
		categoryRoute.GET("/category/:category_id", middleware.VerifyToken(svc), categoryController.GetCategory)
		categoryRoute.PUT("/category/:category_id", middleware.VerifyToken(svc), categoryController.UpdateCategory)
		categoryRoute.DELETE("/category/:category_id", middleware.VerifyToken(svc), categoryController.DeleteCategory)
	}
}
