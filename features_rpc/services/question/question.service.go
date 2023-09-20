package question_service

import models "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/models"

type QuestionService interface {
	CreateQuestion(*models.CreateQuestionRequest) (*models.DBQuestion, error)
	GetQuestions(page int, limit int, order string, sort string) ([]*models.DBQuestion, error)
	GetQuestion(string) (*models.DBQuestion, error)
	UpdateQuestion(string, *models.UpdateQuestion) (*models.DBQuestion, error)
	DeleteQuestion(string) error
	GetQuestionsByUserId(string, int, int, string, string) []*models.DBQuestion
	GetQuestionCount(string) (count int64)
	FilterQuestion(*models.FilterQuestionRequest) ([]*models.DBQuestion, error)
	GetQuestionCountAll() (count int64)
	GetFilteredQuestionCount(*models.FilterQuestionRequest) (count int64)
}
