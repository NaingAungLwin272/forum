package badge_proxy

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/badge/models"
	badge_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/badge/pb"
)

// Badge Processes
func CreateBadge(ctx *gin.Context, bsc badge_proto.BadgeServiceClient) (*badge_proto.BadgeResponse, error) {
	badgeModel := models.CreateBadgeRequest{}

	if err := ctx.BindJSON(&badgeModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := bsc.CreateBadge(context.Background(), &badge_proto.CreateBadgeRequest{
		Name:        badgeModel.Name,
		Description: badgeModel.Description,
		Level:       int32(badgeModel.Level),
	})

	return res, err
}

func GetBadge(ctx *gin.Context, bsc badge_proto.BadgeServiceClient) (*badge_proto.BadgeResponse, error) {
	badgeId := ctx.Param("badgeId")

	res, err := bsc.GetBadge(context.Background(), &badge_proto.BadgeRequest{
		Id: badgeId,
	})

	return res, err
}

func GetBadgeLists(ctx *gin.Context, bsc badge_proto.BadgeServiceClient) ([]*badge_proto.Badge, error) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)
	res, err := bsc.GetBadges(context.Background(), &badge_proto.GetBadgesRequest{
		Page:  &convertedPage,
		Limit: &convertedLimit,
	})

	var badges []*badge_proto.Badge
	badges = append(badges, res.GetBadges()...)

	return badges, err

}

func UpdateBadge(ctx *gin.Context, bsc badge_proto.BadgeServiceClient) (*badge_proto.BadgeResponse, error) {
	badgeId := ctx.Param("badgeId")
	badgeModel := models.UpdateBadge{}

	if err := ctx.BindJSON(&badgeModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := bsc.UpdateBadge(context.Background(), &badge_proto.UpdateBadgeRequest{
		XId:         badgeId,
		Name:        &badgeModel.Name,
		Description: &badgeModel.Description,
		Level:       &badgeModel.Level,
	})

	return res, err
}

func DeleteBadge(ctx *gin.Context, bsc badge_proto.BadgeServiceClient) (*badge_proto.DeleteBadgeResponse, error) {
	badgeId := ctx.Param("badgeId")

	res, err := bsc.DeleteBadge(context.Background(), &badge_proto.BadgeRequest{
		Id: badgeId,
	})

	return res, err
}

// UserBadge Processes
func CreateUserBadge(ctx *gin.Context, bsc badge_proto.UserBadgeServiceClient) (*badge_proto.UserBadgeResponse, error) {
	userBadgeModel := models.CreateUserBadgeRequest{}

	if err := ctx.BindJSON(&userBadgeModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := bsc.CreateUserBadge(context.Background(), &badge_proto.CreateUserBadgeRequest{
		UserId:  userBadgeModel.User_Id,
		BadgeId: userBadgeModel.Badge_Id,
	})

	return res, err
}

func GetUserBadgeByUserId(ctx *gin.Context, bsc badge_proto.UserBadgeServiceClient) ([]*badge_proto.UserBadge, error) {
	userId := ctx.Param("user_id")

	res, err := bsc.GetUserBadgesOfUser(context.Background(), &badge_proto.GetUserBadgesOfUserRequest{
		UserId: userId,
	})

	var userBadges []*badge_proto.UserBadge
	userBadges = append(userBadges, res.GetUserBadges()...)

	return userBadges, err

}

// UserPoint Processes
func CreateUserPoint(ctx *gin.Context, bsc badge_proto.UserPointServiceClient) (*badge_proto.UserPointResponse, error) {
	pointModel := models.CreateUserPointRequest{}

	if err := ctx.BindJSON(&pointModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := bsc.CreateUserPoint(context.Background(), &badge_proto.CreateUserPointRequest{
		UserId:        pointModel.UserId,
		ReactionLevel: pointModel.ReactionLevel,
		QaLevel:       pointModel.QaLevel,
		QuestionCount: pointModel.QuestionCount,
		AnswerCount:   pointModel.AnswerCount,
		SolvedCount:   pointModel.SolvedCount,
	})

	return res, err
}

func GetUserPoint(ctx *gin.Context, bsc badge_proto.UserPointServiceClient) (*badge_proto.UserPointResponse, error) {
	pointId := ctx.Param("user_id")

	res, err := bsc.GetUserPoint(context.Background(), &badge_proto.UserPointRequest{
		Id: pointId,
	})

	return res, err
}

func GetPointsList(ctx *gin.Context, bsc badge_proto.UserPointServiceClient) ([]*badge_proto.UserPoint, error) {
	pointModel := models.PointRequestByPage{}
	res, err := bsc.GetUserPoints(context.Background(), &badge_proto.GetUserPointsRequest{
		Page:  &pointModel.Page,
		Limit: &pointModel.Limit,
	})

	var badges []*badge_proto.UserPoint
	badges = append(badges, res.GetUserPoints()...)

	return badges, err

}

func UpdateUserPoint(ctx *gin.Context, bsc badge_proto.UserPointServiceClient) *badge_proto.UserPointResponse {
	userId := ctx.Param("user_id")
	isSolved, _ := strconv.ParseBool(ctx.Query("is_solved"))

	userPoint, userPointErr := bsc.GetUserPoint(context.Background(), &badge_proto.UserPointRequest{
		Id: userId,
	})

	if userPointErr != nil {
		ctx.AbortWithError(http.StatusBadGateway, userPointErr)
		return nil
	}

	if isSolved {
		bsc.UpdateUserPoint(context.Background(), &badge_proto.UpdateUserPointRequest{
			UserId:        userPoint.UserPoint.UserId,
			ReactionLevel: userPoint.UserPoint.ReactionLevel,
			QaLevel:       userPoint.UserPoint.QaLevel,
			QuestionCount: userPoint.UserPoint.QuestionCount,
			AnswerCount:   userPoint.UserPoint.AnswerCount,
			SolvedCount:   userPoint.UserPoint.SolvedCount + 1,
		})
	} else {
		bsc.UpdateUserPoint(context.Background(), &badge_proto.UpdateUserPointRequest{
			UserId:        userPoint.UserPoint.UserId,
			ReactionLevel: userPoint.UserPoint.ReactionLevel,
			QaLevel:       userPoint.UserPoint.QaLevel,
			QuestionCount: userPoint.UserPoint.QuestionCount,
			AnswerCount:   userPoint.UserPoint.AnswerCount,
			SolvedCount:   userPoint.UserPoint.SolvedCount - 1,
		})
	}

	return nil
}

func DeleteUserPoint(ctx *gin.Context, bsc badge_proto.UserPointServiceClient) (*badge_proto.DeleteUserPointResponse, error) {
	pointId := ctx.Param("pointId")

	res, err := bsc.DeleteUserPoint(context.Background(), &badge_proto.UserPointRequest{
		Id: pointId,
	})

	return res, err
}

func EvaluatePoints(ctx *gin.Context, bsc badge_proto.UserPointServiceClient) (*badge_proto.UserPointEvaluateResponse, error) {
	page := int64(0)
	limit := int64(0)

	res, err := bsc.EvaluatePoints(context.Background(), &badge_proto.GetUserPointsRequest{
		Page:  &page,
		Limit: &limit,
	})

	return res, err
}
