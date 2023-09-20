package category_service

import (
	"errors"

	"github.com/gin-gonic/gin"
	category_client "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	category_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/category"
	category_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/category/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryService struct {
	CategoryClient category_client.ServiceClient
}

// Category Processes
func (categorySvc *CategoryService) Create(ctx *gin.Context) (*category_proto.CategoryResponse, error) {
	data, err := category_proxy.CreateCategory(ctx, categorySvc.CategoryClient.Category)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (categorySvc *CategoryService) GetCategory(ctx *gin.Context) (*category_proto.CategoryResponse, error) {
	data, err := category_proxy.GetCategory(ctx, categorySvc.CategoryClient.Category)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (categorySvc *CategoryService) GetCategories(ctx *gin.Context) ([]*category_proto.Category, error) {
	data, err := category_proxy.GetCategories(ctx, categorySvc.CategoryClient.Category)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (categorySvc *CategoryService) UpdateCategory(ctx *gin.Context) (*category_proto.CategoryResponse, error) {
	data, err := category_proxy.UpdateCategory(ctx, categorySvc.CategoryClient.Category)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (categorySvc *CategoryService) DeleteCategory(ctx *gin.Context) (*category_proto.DeleteCategoryResponse, error) {
	data, err := category_proxy.DeleteCategory(ctx, categorySvc.CategoryClient.Category)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (categorySvc *CategoryService) GetCategoryByType(ctx *gin.Context) ([]*category_proto.Category, error) {
	data, err := category_proxy.GetCategoryByType(ctx, categorySvc.CategoryClient.Category)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (categorySvc *CategoryService) GetCategoryCount(ctx *gin.Context) (*category_proto.CategoryCountResponse, error) {
	data, err := category_proxy.GetCategoryCount(ctx, categorySvc.CategoryClient.Category)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func NewCategoryService(CategoryClient category_client.ServiceClient) CategoryServiceInterface {
	return &CategoryService{
		CategoryClient: CategoryClient,
	}
}
