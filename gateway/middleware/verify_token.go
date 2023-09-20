package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	user_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/pb"
)

func VerifyToken(svc *service.ServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string

		authorization := ctx.Request.Header.Get("Authorization")

		if len(authorization) != 0 {
			access_token = authorization
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		data, err := user_proto.AuthServiceClient.VerifyToken(svc.Auth, context.Background(), &user_proto.VerifyTokenRequest{
			AccessToken: access_token,
		})

		if err != nil || data.IsTokenVerified == false {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		log.Println("Passed..")
		ctx.Next()
	}
}
