package category_proxy

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/category/models"
	category_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/category/pb"
)

func CreateCategory(ctx *gin.Context, csc category_proto.CategoryServiceClient) (*category_proto.CategoryResponse, error) {
	categoryModel := models.CreateCategoryRequest{}

	if err := ctx.BindJSON(&categoryModel); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}

	res, err := csc.CreateCategory(context.Background(), &category_proto.CreateCategoryRequest{
		Type: int32(categoryModel.Type),
		Name: categoryModel.Name,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	return res, err
}

func GetCategories(ctx *gin.Context, csc category_proto.CategoryServiceClient) ([]*category_proto.Category, error) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	res, err := csc.GetCategories(context.Background(), &category_proto.GetCategoriesRequest{
		Page:  &convertedPage,
		Limit: &convertedLimit,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var categories []*category_proto.Category
	categories = append(categories, res.GetCategories()...)

	return categories, err
}

func GetCategoryByType(ctx *gin.Context, csc category_proto.CategoryServiceClient) ([]*category_proto.Category, error) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	cateType := ctx.Query("type")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)
	convertedCate, _ := strconv.ParseInt(cateType, 0, 64)

	res, err := csc.GetCategoryByType(context.Background(), &category_proto.GetCategoryByTypeRequest{
		Type:  int32(convertedCate),
		Page:  &convertedPage,
		Limit: &convertedLimit,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var categories []*category_proto.Category
	categories = append(categories, res.GetCategories()...)

	return categories, err
}

func GetCategory(ctx *gin.Context, csc category_proto.CategoryServiceClient) (*category_proto.CategoryResponse, error) {

	categoryId := ctx.Param("category_id")

	res, err := csc.GetCategory(context.Background(), &category_proto.CategoryRequest{
		Id: categoryId,
	})

	return res, err
}

func UpdateCategory(ctx *gin.Context, csc category_proto.CategoryServiceClient) (*category_proto.CategoryResponse, error) {
	categoryModel := category_proto.UpdateCategoryRequest{}
	categoryId := ctx.Param("category_id")

	if err := ctx.BindJSON(&categoryModel); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	res, err := csc.UpdateCategory(context.Background(), &category_proto.UpdateCategoryRequest{
		XId:  categoryId,
		Name: categoryModel.Name,
		Type: categoryModel.Type,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	return res, err
}

func DeleteCategory(ctx *gin.Context, bsc category_proto.CategoryServiceClient) (*category_proto.DeleteCategoryResponse, error) {
	categoryId := ctx.Param("category_id")

	res, err := bsc.DeleteCategory(context.Background(), &category_proto.CategoryRequest{
		Id: categoryId,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
	}

	return res, err
}

func GetCategoryCount(ctx *gin.Context, csc category_proto.CategoryServiceClient) (*category_proto.CategoryCountResponse, error) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)
	res, err := csc.GetCategoryCount(context.Background(), &category_proto.GetCategoriesRequest{
		Page:  &convertedPage,
		Limit: &convertedLimit,
	})
	return res, err
}
