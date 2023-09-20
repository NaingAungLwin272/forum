package mail_proxy

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	mail_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/mail/pb"
)

func ForgetPasswordMail(ctx *gin.Context, msc mail_proto.MailServiceClient) *mail_proto.MailResponse {
	var req mail_proto.ForgetMailRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {

		if err == io.EOF {
			return nil
		} else {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return nil
		}
	}

	res, err := msc.ForgetPasswordMail(context.Background(), &req)
	fmt.Println(err, "error.....")
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}

	return res
}

func SendMail(ctx *gin.Context, msc mail_proto.MailServiceClient) *mail_proto.MailResponse {
	var req mail_proto.MailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if err == io.EOF {
			return nil
		} else {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return nil
		}
	}
	res, err := msc.SendMail(context.Background(), &req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}

	return res
}
