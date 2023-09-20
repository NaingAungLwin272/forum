package question_service

import (
	"errors"

	"github.com/gin-gonic/gin"
	clients "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/models"
	question_pb "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	features_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/proxy"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionService struct {
	QuestionClient clients.ServiceClient
}

// GetBookmarksByUserId implements BookmarkServiceInterface
func (QuestionSvc *QuestionService) GetQuestionsByUserId(ctx *gin.Context) []*question_pb.Question {
	data := features_proxy.GetQuestionsByUserId(ctx, QuestionSvc.QuestionClient.Question, QuestionSvc.QuestionClient.Comment)
	return data
}

// GetBookmarksByUserId implements BookmarkServiceInterface
func (QuestionSvc *QuestionService) GetQuestionById(ctx *gin.Context) (*models.QuestionDetail, error) {
	data, err := features_proxy.GetQuestionById(ctx, QuestionSvc.QuestionClient.Question, QuestionSvc.QuestionClient.Comment)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, err
	}
	return data, err
}

// Create implements UserServiceInterface.
func (QuestionSvc *QuestionService) Create(ctx *gin.Context) (*features_proxy.CreateQuestionResponse, error) {
	data, err := features_proxy.CreateQuestion(ctx, QuestionSvc.QuestionClient.Question, QuestionSvc.QuestionClient.Comment, QuestionSvc.QuestionClient.User, QuestionSvc.QuestionClient.UserPoint)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (QuestionSvc *QuestionService) GetQuestions(ctx *gin.Context) []*question_pb.Question {
	data := features_proxy.GetQuestions(ctx, QuestionSvc.QuestionClient.Question, QuestionSvc.QuestionClient.Comment)
	return data
}

func (QuestionSvc *QuestionService) FilterQuestion(ctx *gin.Context) []*question_pb.Question {
	data := features_proxy.FilterQuestion(ctx, QuestionSvc.QuestionClient.Question, QuestionSvc.QuestionClient.Comment)
	return data
}

// GetQuestionCountAll implements QuestionServiceInterface.
func (QuestionSvc *QuestionService) GetQuestionCountAll(ctx *gin.Context) (*question_pb.QuestionCountResponse, error) {
	data, err := features_proxy.GetQuestionCountAll(ctx, QuestionSvc.QuestionClient.Question)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, err
	}
	return data, err
}

func (QuestionSvc *QuestionService) GetFilteredQuestionCount(ctx *gin.Context) (*question_pb.QuestionCountResponse, error) {
	data, err := features_proxy.GetFilteredQuestionCount(ctx, QuestionSvc.QuestionClient.Question)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, err
	}
	return data, err
}

func NewQuestionService(QuestionClient clients.ServiceClient) QuestionServiceInterface {
	return &QuestionService{
		QuestionClient: QuestionClient,
	}
}
