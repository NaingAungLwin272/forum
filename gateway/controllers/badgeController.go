package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	badge_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/badge"
)

type BadgeController struct {
	BadgeServiceInterface badge_service.BadgeServiceInterface
}

// Badge Processes
func (controller *BadgeController) CreateBadge(ctx *gin.Context) {
	data, err := controller.BadgeServiceInterface.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *BadgeController) GetBadge(ctx *gin.Context) {
	data, err := controller.BadgeServiceInterface.GetBadge(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *BadgeController) GetBadgesList(ctx *gin.Context) {
	data, err := controller.BadgeServiceInterface.GetBadgesList(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &data)
	}
}

func (controller *BadgeController) UpdateBadge(ctx *gin.Context) {
	data, err := controller.BadgeServiceInterface.UpdateBadge(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *BadgeController) DeleteBadge(ctx *gin.Context) {
	data, err := controller.BadgeServiceInterface.DeleteBadge(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

// UserBadge Processes
func (controller *BadgeController) CreateUserBadge(ctx *gin.Context) {
	data, err := controller.BadgeServiceInterface.CreateUserBadge(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *BadgeController) GetUserBadgeByUserId(ctx *gin.Context) {
	data, err := controller.BadgeServiceInterface.GetUserBadgeByUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &data)
	}
}

// UserPoint Processes
func (controller *BadgeController) CreateUserPoint(ctx *gin.Context) {
	data, err := controller.BadgeServiceInterface.CreateUserPoint(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *BadgeController) GetUserPoint(ctx *gin.Context) {
	data, err := controller.BadgeServiceInterface.GetUserPoint(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *BadgeController) GetPointsList(ctx *gin.Context) {
	data, err := controller.BadgeServiceInterface.GetPointsList(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *BadgeController) UpdateUserPoint(ctx *gin.Context) {
	data := controller.BadgeServiceInterface.UpdateUserPoint(ctx)
	ctx.JSON(http.StatusOK, &data)
}

func (controller *BadgeController) DeleteUserPoint(ctx *gin.Context) {
	data, err := controller.BadgeServiceInterface.DeleteUserPoint(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *BadgeController) EvaluatePoints(ctx *gin.Context) {
	data, err := controller.BadgeServiceInterface.EvaluatePoints(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func NewBadgeController(BadgeServiceInterface badge_service.BadgeServiceInterface) *BadgeController {
	return &BadgeController{
		BadgeServiceInterface: BadgeServiceInterface,
	}
}
