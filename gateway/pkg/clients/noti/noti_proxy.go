package noti_proxy

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	mail_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/mail/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/noti/models"
	noti_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/noti/pb"
	user_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/pb"
)

func CreateNoti(ctx *gin.Context, nsc noti_proto.NotiServiceClient, msc mail_proto.MailServiceClient, usc user_proto.UserServiceClient) ([]*noti_proto.NotiResponse, error) {
	b := models.CreateNotiRequest{}
	if err := ctx.BindJSON(&b); err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("body can't be empty")
		} else {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return nil, err
		}
	}

	responses := make([]*noti_proto.NotiResponse, 0)
	for _, userID := range b.UserId {
		res, err := nsc.CreateNoti(context.Background(), &noti_proto.CreateNotiRequest{
			UserId:      userID,
			Type:        int64(b.Type),
			Name:        b.Name,
			Description: b.Description,
			Link:        &b.Link,
			Status:      *b.Status,
		})
		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return nil, err
		}
		responses = append(responses, res)

		ures, _ := usc.GetUser(context.Background(), &user_proto.UserRequest{
			XId: res.Noti.UserId,
		})
		if ures.MailSubscribe {
			displayName := ures.DisplayName
			msc.SendMail(context.Background(), &mail_proto.MailRequest{
				Email:   ures.Email,
				Type:    6,
				Link:    res.Noti.Link,
				Subject: fmt.Sprintf("%s - %s", res.Noti.Description, displayName),
			})
		}
	}

	return responses, nil
}

func DeleteNoti(ctx *gin.Context, nsc noti_proto.NotiServiceClient) *noti_proto.DeleteNotiResponse {
	notiId := ctx.Param("id")
	res, err := nsc.DeleteNoti(context.Background(), &noti_proto.NotiRequest{
		XId: notiId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}
	return res
}

func GetNotiByUserId(ctx *gin.Context, nsc noti_proto.NotiServiceClient) []*noti_proto.Noti {
	userId := ctx.Param("id")
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)
	b := noti_proto.UserIdRequest{
		UserId: userId,
		Page:   &convertedPage,
		Limit:  &convertedLimit,
	}

	stream, err := nsc.GetNotiByUserId(context.Background(), &b)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}

	var notis []*noti_proto.Noti
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return nil
		}

		notis = append(notis, res)
	}
	return notis
}

func GetNoti(ctx *gin.Context, nsc noti_proto.NotiServiceClient) *noti_proto.NotiResponse {
	notiId := ctx.Param("id")
	res, err := nsc.GetNoti(context.Background(), &noti_proto.NotiRequest{
		XId: notiId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}
	return res
}

func GetNotis(ctx *gin.Context, nsc noti_proto.NotiServiceClient) []*noti_proto.Noti {
	b := noti_proto.GetNotisRequest{}

	if err := ctx.ShouldBindJSON(&b); err != nil {
		if err == io.EOF {
			defaultPage := int64(0)
			defaultLimit := int64(10)
			b.Page = &defaultPage
			b.Limit = &defaultLimit
		} else {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return nil
		}
	}

	stream, err := nsc.GetNotis(context.Background(), &b)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}

	var notis []*noti_proto.Noti
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return nil
		}

		notis = append(notis, res)
	}

	return notis
}

func UpdateNoti(ctx *gin.Context, nsc noti_proto.NotiServiceClient) *noti_proto.NotiResponse {
	b := models.UpdateNoti{}
	notiId := ctx.Param("id")
	if err := ctx.BindJSON(&b); err != nil {
		if err == io.EOF {
			return nil
		} else {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return nil
		}
	}

	res, err := nsc.UpdateNoti(context.Background(), &noti_proto.UpdateNotiRequest{
		XId:         notiId,
		Name:        &b.Name,
		Description: &b.Description,
		Link:        &b.Link,
		Status:      &b.Status,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}
	return res
}

func GetNotiCount(ctx *gin.Context, nsc noti_proto.NotiServiceClient) *noti_proto.NotiCountResponse {
	userId := ctx.Param("id")
	b := noti_proto.NotiRequestByUserId{
		UserId: userId,
	}
	res, err := nsc.GetNotiCount(context.Background(), &b)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}
	return res
}

func MarkAllNotiAsRead(ctx *gin.Context, nsc noti_proto.NotiServiceClient) *noti_proto.DeleteNotiResponse {
	userId := ctx.Param("id")
	b := noti_proto.NotiRequestByUserId{
		UserId: userId,
	}
	res, err := nsc.MarkAllNotiAsRead(context.Background(), &b)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}
	return res
}
