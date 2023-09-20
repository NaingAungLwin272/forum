package question_service

import (
	"github.com/gin-gonic/gin"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/models"
	question_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	features_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/proxy"
)

type CreateQuestionResponse struct {
	QuestionRes *question_proto.QuestionResponse
	CommentId   string
}
type QuestionServiceInterface interface {
	Create(ctx *gin.Context) (*features_proxy.CreateQuestionResponse, error)
	GetQuestions(ctx *gin.Context) []*question_proto.Question
	GetQuestionById(ctx *gin.Context) (*models.QuestionDetail, error)
	GetQuestionsByUserId(ctx *gin.Context) []*question_proto.Question
	FilterQuestion(ctx *gin.Context) []*question_proto.Question
	GetQuestionCountAll(ctx *gin.Context) (*question_proto.QuestionCountResponse, error)
	GetFilteredQuestionCount(ctx *gin.Context) (*question_proto.QuestionCountResponse, error)
}
