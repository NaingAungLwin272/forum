package routes

import (
	"github.com/gin-gonic/gin"
	vote_controllers "github.com/scm-dev1dev5/mtm-community-forum/gateway/controllers"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/middleware"
	service "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	bookmark_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/features/bookmark"
	comment_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/features/comment"
	mention_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/features/mention"
	question_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/features/question"
	view_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/features/view"
	vote_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/features/vote"
)

func FeatureRoute(apiRouter *gin.RouterGroup, svc *service.ServiceClient) {
	/**
	Vote Controller
	*/

	voteService := vote_service.NewVoteService(*svc)
	bookmarkService := bookmark_service.NewBookmarkService(*svc)
	mentionService := mention_service.NewMentionService(*svc)
	viewService := view_service.NewViewService(*svc)
	questionService := question_service.NewQuestionService(*svc)

	voteController := vote_controllers.NewVoteController(voteService)
	bookmarkController := vote_controllers.NewBookmarkController(bookmarkService)
	mentionController := vote_controllers.NewMentionController(mentionService)
	viewController := vote_controllers.NewViewController(viewService)
	questionController := vote_controllers.NewQuestionController(questionService)

	voteRoute := apiRouter.Group("/")
	{
		voteRoute.POST("/votes", middleware.VerifyToken(svc), voteController.CreateVote)
		voteRoute.GET("/users/:user_id/votes", middleware.VerifyToken(svc), voteController.GetVotesByUserId)
		voteRoute.DELETE("/votes/:vote_id", middleware.VerifyToken(svc), voteController.DeleteVote)
		voteRoute.GET("/votes/user/:user_id/question/:question_id", middleware.VerifyToken(svc), voteController.GetVotesByUserIdQuestionId)
	}
	bookmarkRoute := apiRouter.Group("/")
	{
		bookmarkRoute.POST("/bookmarks", middleware.VerifyToken(svc), bookmarkController.CreateBookmark)
		bookmarkRoute.GET("/users/:user_id/bookmarks", middleware.VerifyToken(svc), bookmarkController.GetBookmarkByUserId)
		bookmarkRoute.DELETE("/bookmarks/:bookmark_id", middleware.VerifyToken(svc), bookmarkController.DeleteBookmark)
		bookmarkRoute.GET("/bookmarks/user/:user_id/question/:question_id", middleware.VerifyToken(svc), bookmarkController.GetBookmarksByUserIdQuestionId)
	}
	mentionRoute := apiRouter.Group("/")
	{
		mentionRoute.POST("/mentions", middleware.VerifyToken(svc), mentionController.CreateMention)
		mentionRoute.GET("/users/:user_id/mentions", middleware.VerifyToken(svc), mentionController.GetMentionsByUserId)
	}
	viewRoute := apiRouter.Group("/")
	{
		viewRoute.POST("/views", middleware.VerifyToken(svc), viewController.CreateView)
		viewRoute.GET("/users/:user_id/views", middleware.VerifyToken(svc), viewController.GetViewsByUserId)
	}
	questionRoute := apiRouter.Group("/")
	{
		questionRoute.POST("/question", middleware.VerifyToken(svc), questionController.CreateQuestion)
		questionRoute.GET("/questions", middleware.VerifyToken(svc), questionController.GetQuestions)
		questionRoute.GET("/questions/:question_id", middleware.VerifyToken(svc), questionController.GetQuestionById)
		questionRoute.GET("/questions/count", middleware.VerifyToken(svc), questionController.GetQuestionCountAll)
		questionRoute.GET("/users/:user_id/questions", middleware.VerifyToken(svc), questionController.GetQuestionsByUserId)
		questionRoute.POST("questions/search", middleware.VerifyToken(svc), questionController.FilterQuestion)
		questionRoute.POST("questions/filteredquestioncount", middleware.VerifyToken(svc), questionController.GetFilteredQuestionCount)
	}

	commentService := comment_service.NewCommentService(*svc)
	commentController := vote_controllers.NewCommentController(commentService)

	commentRoute := apiRouter.Group("/")
	{
		commentRoute.GET("/comment/:comment_id", middleware.VerifyToken(svc), commentController.GetComment)
		commentRoute.POST("/comment", middleware.VerifyToken(svc), commentController.CreateComment)
		commentRoute.PUT("/comment/:comment_id", middleware.VerifyToken(svc), commentController.UpdateComment)
		commentRoute.DELETE("/comment/:comment_id", middleware.VerifyToken(svc), commentController.DeleteComment)
	}
	apiRouter.GET("/question/:question_id/comments", middleware.VerifyToken(svc), commentController.GetCommentByQuestionId)
	apiRouter.GET("/users/:user_id/comments", middleware.VerifyToken(svc), commentController.GetCommentsByUserId)
	apiRouter.GET("/users/:user_id/answers", middleware.VerifyToken(svc), commentController.GetAnswersByUserId)
	apiRouter.GET("/users/:user_id/solved", middleware.VerifyToken(svc), commentController.GetCommentsByUserIdWithSolved)
}
