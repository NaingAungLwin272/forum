package category_service

import (
	"github.com/gin-gonic/gin"
	category_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/category/pb"
)

type CategoryServiceInterface interface {
	Create(ctx *gin.Context) (*category_proto.CategoryResponse, error)
	GetCategory(ctx *gin.Context) (*category_proto.CategoryResponse, error)
	GetCategories(ctx *gin.Context) ([]*category_proto.Category, error)
	UpdateCategory(ctx *gin.Context) (*category_proto.CategoryResponse, error)
	DeleteCategory(ctx *gin.Context) (*category_proto.DeleteCategoryResponse, error)
	GetCategoryByType(ctx *gin.Context) ([]*category_proto.Category, error)
	GetCategoryCount(ctx *gin.Context) (*category_proto.CategoryCountResponse, error)
}
