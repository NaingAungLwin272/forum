package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	category_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/category"
)

type CategoryController struct {
	CategoryServiceInterface category_service.CategoryServiceInterface
}

// Category Processes
func (controller *CategoryController) CreateCategory(ctx *gin.Context) {
	data, err := controller.CategoryServiceInterface.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &data)
	}
}

func (controller *CategoryController) GetCategory(ctx *gin.Context) {
	data, err := controller.CategoryServiceInterface.GetCategory(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *CategoryController) GetCategories(ctx *gin.Context) {
	data, err := controller.CategoryServiceInterface.GetCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *CategoryController) UpdateCategory(ctx *gin.Context) {
	data, err := controller.CategoryServiceInterface.UpdateCategory(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &data)
	}
}

func (controller *CategoryController) DeleteCategory(ctx *gin.Context) {
	data, err := controller.CategoryServiceInterface.DeleteCategory(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *CategoryController) GetCategoryByType(ctx *gin.Context) {
	data, err := controller.CategoryServiceInterface.GetCategoryByType(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, &data)
}

func (controller *CategoryController) GetCategoryCount(ctx *gin.Context) {
	data, err := controller.CategoryServiceInterface.GetCategoryCount(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func NewCategoryController(CategoryServiceInterface category_service.CategoryServiceInterface) *CategoryController {
	return &CategoryController{
		CategoryServiceInterface: CategoryServiceInterface,
	}
}
