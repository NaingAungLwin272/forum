package question_gapi

import (
	"context"
	"strings"
	"time"

	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (questionServer *QuestionServer) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*pb.QuestionResponse, error) {
	question := &models.CreateQuestionRequest{
		UserId:      req.GetUserId(),
		Title:       req.GetTitle(),
		LanguageIds: req.GetLanguageIds(),
		TagIds:      req.GetTagIds(),
		ViewCount:   req.GetViewCount(),
		VoteCount:   req.GetVoteCount(),
		ReplyCount:  req.GetReplyCount(),
		IsDeleted:   req.GetIsDeleted(),
		UserIds:     req.GetUserIds(),
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
	}

	newQuestion, _ := questionServer.questionService.CreateQuestion(question)

	res := &pb.QuestionResponse{
		XId:           newQuestion.Id.Hex(),
		UserId:        newQuestion.UserId,
		Title:         newQuestion.Title,
		LanguageIds:   newQuestion.LanguageIds,
		TagIds:        newQuestion.TagIds,
		ViewCount:     newQuestion.ViewCount,
		VoteCount:     newQuestion.VoteCount,
		ReplyCount:    newQuestion.ReplyCount,
		SolutionCount: newQuestion.SolutionCount,
		UserIds:       newQuestion.UserIds,
		IsDeleted:     wrapperspb.Bool(newQuestion.IsDeleted),
		CreatedAt:     timestamppb.New(newQuestion.CreateAt),
		UpdatedAt:     timestamppb.New(newQuestion.UpdatedAt),
	}
	return res, nil
}

func (questionServer *QuestionServer) GetQuestions(req *pb.GetQuestionsRequest, stream pb.QuestionService_GetQuestionsServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()

	question_service, err := questionServer.questionService.GetQuestions(int(page), int(limit), req.GetOrder(), req.GetSort())
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, question := range question_service {
		stream.Send(&pb.Question{
			XId:           question.Id.Hex(),
			UserId:        question.UserId,
			Title:         question.Title,
			LanguageIds:   question.LanguageIds,
			TagIds:        question.TagIds,
			ViewCount:     question.ViewCount,
			VoteCount:     question.VoteCount,
			ReplyCount:    question.ReplyCount,
			SolutionCount: question.SolutionCount,
			UserIds:       question.UserIds,
			IsDeleted:     wrapperspb.Bool(question.IsDeleted),
			CreatedAt:     timestamppb.New(question.CreateAt),
			UpdatedAt:     timestamppb.New(question.UpdatedAt),
		})
	}

	return nil
}

func (questionServer *QuestionServer) GetQuestion(context context.Context, req *pb.QuestionRequest) (*pb.QuestionResponse, error) {
	questionId := req.GetXId()

	question, err := questionServer.questionService.GetQuestion(questionId)

	if err != nil {
		if strings.Contains(err.Error(), "question") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.QuestionResponse{
		XId:           question.Id.Hex(),
		UserId:        question.UserId,
		Title:         question.Title,
		LanguageIds:   question.LanguageIds,
		TagIds:        question.TagIds,
		ViewCount:     question.ViewCount,
		VoteCount:     question.VoteCount,
		ReplyCount:    question.ReplyCount,
		SolutionCount: question.SolutionCount,
		UserIds:       question.UserIds,
		IsDeleted:     wrapperspb.Bool(question.IsDeleted),
		CreatedAt:     timestamppb.New(question.CreateAt),
		UpdatedAt:     timestamppb.New(question.UpdatedAt),
	}

	return res, err
}

func (questionServer *QuestionServer) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*pb.QuestionResponse, error) {
	questionId := req.GetXId()
	question := &models.UpdateQuestion{
		Title:         req.GetTitle(),
		LanguageIds:   req.LanguageIds,
		TagIds:        req.TagIds,
		ViewCount:     req.GetViewCount(),
		VoteCount:     req.GetVoteCount(),
		ReplyCount:    req.GetReplyCount(),
		SolutionCount: req.GetSolutionCount(),
		UserIds:       req.GetUserIds(),
		IsDeleted:     req.GetIsDeleted(),
		UpdatedAt:     time.Now(),
	}

	updatedQuestion, err := questionServer.questionService.UpdateQuestion(questionId, question)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.QuestionResponse{
		XId:           updatedQuestion.Id.Hex(),
		UserId:        updatedQuestion.UserId,
		Title:         updatedQuestion.Title,
		LanguageIds:   updatedQuestion.LanguageIds,
		TagIds:        updatedQuestion.TagIds,
		ViewCount:     updatedQuestion.ViewCount,
		VoteCount:     updatedQuestion.VoteCount,
		ReplyCount:    updatedQuestion.ReplyCount,
		SolutionCount: updatedQuestion.SolutionCount,
		IsDeleted:     wrapperspb.Bool(updatedQuestion.IsDeleted),
		CreatedAt:     timestamppb.New(updatedQuestion.CreateAt),
		UpdatedAt:     timestamppb.New(updatedQuestion.UpdatedAt),
	}

	return res, nil
}

func (questionServer *QuestionServer) DeleteQuestion(ctx context.Context, req *pb.QuestionRequest) (*pb.DeleteQuestionResponse, error) {
	questionId := req.GetXId()
	if err := questionServer.questionService.DeleteQuestion(questionId); err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	res := &pb.DeleteQuestionResponse{
		Success: true,
	}
	return res, nil
}

func (questionServer *QuestionServer) GetQuestionByUserId(req *pb.QuestionResquestByUserId, stream pb.QuestionService_GetQuestionByUserIdServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()
	userId := req.GetUserId()
	questions := questionServer.questionService.GetQuestionsByUserId(userId, int(page), int(limit), req.GetOrder(), req.GetSort())

	for _, question := range questions {
		stream.Send(&pb.Question{
			XId:           question.Id.Hex(),
			UserId:        question.UserId,
			Title:         question.Title,
			LanguageIds:   question.LanguageIds,
			TagIds:        question.TagIds,
			ViewCount:     question.ViewCount,
			ReplyCount:    question.ReplyCount,
			VoteCount:     question.VoteCount,
			SolutionCount: question.SolutionCount,
			UserIds:       question.UserIds,
			IsDeleted:     wrapperspb.Bool(question.IsDeleted),
			CreatedAt:     timestamppb.New(question.CreateAt),
			UpdatedAt:     timestamppb.New(question.UpdatedAt),
		})

	}

	return nil
}

func (questionServer *QuestionServer) FilterQuestion(req *pb.FilterQuestionRequest, stream pb.QuestionService_FilterQuestionServer) error {
	question := &models.FilterQuestionRequest{
		Page:        req.GetPage(),
		Limit:       req.GetLimit(),
		LanguageIds: req.GetLanguageId(),
		TagIds:      req.GetTagId(),
		UserId:      req.GetUserId(),
		Title:       req.GetTile(),
		Sort:        req.GetSort(),
		Order:       req.GetOrder(),
	}

	filteredQuestion, err := questionServer.questionService.FilterQuestion(question)
	if err != nil {
		return nil
	}
	for _, filterQuestion := range filteredQuestion {
		if filterQuestion.IsDeleted {
			return nil
		}
		stream.Send(&pb.Question{
			XId:           filterQuestion.Id.Hex(),
			UserId:        filterQuestion.UserId,
			Title:         filterQuestion.Title,
			LanguageIds:   filterQuestion.LanguageIds,
			TagIds:        filterQuestion.TagIds,
			ViewCount:     filterQuestion.ViewCount,
			VoteCount:     filterQuestion.VoteCount,
			ReplyCount:    filterQuestion.ReplyCount,
			SolutionCount: filterQuestion.SolutionCount,
			UserIds:       filterQuestion.UserIds,
			IsDeleted:     wrapperspb.Bool(filterQuestion.IsDeleted),
			CreatedAt:     timestamppb.New(filterQuestion.CreateAt),
			UpdatedAt:     timestamppb.New(filterQuestion.UpdatedAt),
		})
	}
	return nil

}

func (questionServer *QuestionServer) GetQuestionCount(context context.Context, req *pb.QuestionResquestByUserId) (*pb.QuestionCountResponse, error) {
	userId := req.GetUserId()
	questions := questionServer.questionService.GetQuestionCount(userId)

	res := &pb.QuestionCountResponse{
		Count: int64(questions),
	}

	return res, nil
}

func (questionServer *QuestionServer) GetFilteredQuestionCount(context context.Context, req *pb.FilterQuestionRequest) (*pb.QuestionCountResponse, error) {
	question := &models.FilterQuestionRequest{
		LanguageIds: req.GetLanguageId(),
		TagIds:      req.GetTagId(),
		UserId:      req.GetUserId(),
		Title:       req.GetTile(),
		Sort:        req.GetSort(),
		Order:       req.GetOrder(),
	}

	filteredQuestion := questionServer.questionService.GetFilteredQuestionCount(question)

	res := &pb.QuestionCountResponse{
		Count: int64(filteredQuestion),
	}

	return res, nil

}

func (questionServer *QuestionServer) GetQuestionCountAll(context context.Context, req *pb.GetQuestionCountRequest) (*pb.QuestionCountResponse, error) {
	questions := questionServer.questionService.GetQuestionCountAll()

	res := &pb.QuestionCountResponse{
		Count: int64(questions),
	}

	return res, nil
}
