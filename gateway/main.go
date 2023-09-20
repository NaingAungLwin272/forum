package main

import (
	"log"

	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/routes"

	"github.com/gin-contrib/cors"
	clients "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
)

func main() {
	config, err := config.LoadConfig()
	route := gin.Default()

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AddAllowHeaders("authorization")
	corConfig.AllowHeaders = []string{"Content-Type", "Authorization", "X-Api-Key"}
	route.Use(cors.New(corConfig))

	authClient, userClient, userTeamClient, userDeparmentClient, categoryClient, notiClient, badgeClient, userBadgeClient, userPointClient, MailClient, voteClient, bookmarkClient, mentionClient, viewClient, questionClient, commentClient := clients.InitServiceClient(&config)
	svc := &clients.ServiceClient{
		//auth
		Auth: authClient,
		//user
		User:       userClient,
		Team:       userTeamClient,
		Department: userDeparmentClient,
		Category:   categoryClient,
		Noti:       notiClient,
		//badge
		Badge:     badgeClient,
		UserBadge: userBadgeClient,
		UserPoint: userPointClient,
		Mail:      MailClient,
		Vote:      voteClient,
		BookMark:  bookmarkClient,
		Mention:   mentionClient,
		View:      viewClient,
		Question:  questionClient,
		Comment:   commentClient,
	}

	if err != nil {
		log.Fatalln("Failed at configs", err)
	}
	apiRouter := route.Group("/api")

	routes.UserRoute(apiRouter, svc)
	routes.CategoryRoute(apiRouter, svc)
	routes.NotiRoute(apiRouter, svc)
	routes.BadgeRoute(apiRouter, svc)
	routes.MailRoute(apiRouter, svc)
	routes.FeatureRoute(apiRouter, svc)

	// Start the server
	route.Run(":3000")
}
