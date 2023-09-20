package features_proxy

import (
	"context"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/models"
	view_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
)

func CreateView(ctx *gin.Context, vsc view_proto.ViewServiceClient, qsc view_proto.QuestionServiceClient) (*view_proto.ViewResponse, error) {
	viewModel := models.CreateViewRequest{}

	if err := ctx.BindJSON(&viewModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	stream, _ := vsc.GetViewsByUserIdQuestionId(context.Background(), &view_proto.ViewRequestByUserIdQuestionId{
		UserId:     viewModel.User_Id,
		QuestionId: viewModel.Question_Id,
	})

	var views []*view_proto.View

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}

		views = append(views, res)
	}

	if len(views) > 0 {
		return nil, nil
	}

	res, err := vsc.CreateView(context.Background(), &view_proto.CreateViewRequest{
		UserId:     viewModel.User_Id,
		QuestionId: viewModel.Question_Id,
	})

	getViewResultByQuestionId, _ := qsc.GetQuestion(context.Background(), &view_proto.QuestionRequest{
		XId: res.QuestionId,
	})

	viewCountByQuestionId := getViewResultByQuestionId.ViewCount + 1

	qsc.UpdateQuestion(context.Background(), &view_proto.UpdateQuestionRequest{
		XId:        getViewResultByQuestionId.XId,
		VoteCount:  &getViewResultByQuestionId.VoteCount,
		ReplyCount: &getViewResultByQuestionId.ReplyCount,
		ViewCount:  &viewCountByQuestionId,
	})

	return res, err
}

func GetViews(ctx *gin.Context, vsc view_proto.ViewServiceClient) ([]*view_proto.View, error) {
	viewModel := models.ViewRequestByPage{}

	stream, err := vsc.GetViews(context.Background(), &view_proto.GetViewsRequest{
		Page:  &viewModel.Page,
		Limit: &viewModel.Limit,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var views []*view_proto.View

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}

		views = append(views, res)
	}

	return views, err
}

func GetView(ctx *gin.Context, vsc view_proto.ViewServiceClient) (*view_proto.ViewResponse, error) {
	viewId := ctx.Param("view_id")

	res, err := vsc.GetView(context.Background(), &view_proto.ViewRequest{
		XId: viewId,
	})

	return res, err
}

func UpdateView(ctx *gin.Context, vsc view_proto.ViewServiceClient) (*view_proto.ViewResponse, error) {
	viewId := ctx.Param("view_id")
	viewModel := models.UpdateView{}

	if err := ctx.BindJSON(&viewModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := vsc.UpdateView(context.Background(), &view_proto.UpdateViewRequest{
		XId:        viewId,
		UserId:     &viewModel.User_Id,
		QuestionId: &viewModel.Question_Id,
	})

	return res, err
}

func DeleteView(ctx *gin.Context, vsc view_proto.ViewServiceClient) (*view_proto.DeleteViewResponse, error) {
	viewId := ctx.Param("view_id")

	res, err := vsc.DeleteView(context.Background(), &view_proto.ViewRequest{
		XId: viewId,
	})

	return res, err
}

func GetViewsUserId(ctx *gin.Context, vsc view_proto.ViewServiceClient) []*view_proto.View {
	userId := ctx.Param("user_id")

	stream, err := vsc.GetViewsByUserId(context.Background(), &view_proto.ViewRequestByUserId{
		UserId: userId,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}

	var views []*view_proto.View
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return nil
		}

		views = append(views, res)
	}

	return views

}
