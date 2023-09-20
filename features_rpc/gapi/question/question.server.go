package question_gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/pb"
	question_service "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/services/question"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionServer struct {
	pb.UnimplementedQuestionServiceServer
	questionCollection *mongo.Collection
	questionService    question_service.QuestionService
}

func NewGrpcQuestionServer(questionCollection *mongo.Collection, questionService question_service.QuestionService) (*QuestionServer, error) {
	questionServer := &QuestionServer{
		questionCollection: questionCollection,
		questionService:    questionService,
	}

	return questionServer, nil
}
