package gapi

import (
	"context"
	"strings"
	"time"

	models "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (UserPointServer *UserPointServer) CreateUserPoint(ctx context.Context, req *pb.CreateUserPointRequest) (*pb.UserPointResponse, error) {
	UserPoint := &models.CreateUserPointRequest{
		UserId:        req.UserId,
		ReactionLevel: req.ReactionLevel,
		QaLevel:       req.QaLevel,
		QuestionCount: req.QuestionCount,
		AnswerCount:   req.AnswerCount,
		SolvedCount:   req.SolvedCount,
	}

	newUserPoint, err := UserPointServer.UserPointService.CreateUserPoint(UserPoint)

	if err != nil {
		if strings.Contains(err.Error(), "userId already exists") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserPointResponse{
		UserPoint: &pb.UserPoint{
			XId:           newUserPoint.Id.Hex(),
			UserId:        newUserPoint.UserId,
			ReactionLevel: newUserPoint.ReactionLevel,
			QaLevel:       newUserPoint.QaLevel,
			QuestionCount: newUserPoint.QuestionCount,
			AnswerCount:   newUserPoint.AnswerCount,
			SolvedCount:   newUserPoint.SolvedCount,
			CreatedAt:     timestamppb.New(newUserPoint.CreateAt),
			UpdatedAt:     timestamppb.New(newUserPoint.UpdatedAt),
		},
	}
	return res, nil
}

func (UserPointServer *UserPointServer) GetUserPoint(ctx context.Context, req *pb.UserPointRequest) (*pb.UserPointResponse, error) {
	userId := req.GetId()

	UserPoint, err := UserPointServer.UserPointService.GetUserPoint(userId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())

		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserPointResponse{
		UserPoint: &pb.UserPoint{
			XId:           UserPoint.Id.Hex(),
			UserId:        UserPoint.UserId,
			ReactionLevel: UserPoint.ReactionLevel,
			QaLevel:       UserPoint.QaLevel,
			QuestionCount: UserPoint.QuestionCount,
			AnswerCount:   UserPoint.AnswerCount,
			SolvedCount:   UserPoint.SolvedCount,
			CreatedAt:     timestamppb.New(UserPoint.CreateAt),
			UpdatedAt:     timestamppb.New(UserPoint.UpdatedAt),
		},
	}
	return res, nil
}

func (UserPointServer *UserPointServer) GetUserPoints(ctx context.Context, req *pb.GetUserPointsRequest) (*pb.UserPointResponseList, error) {
	var page = req.GetPage()
	var limit = req.GetLimit()

	UserPoints, err := UserPointServer.UserPointService.GetUserPoints(int(page), int(limit))
	if err != nil {
		return nil, err
	}

	cateList := make([]*pb.UserPoint, 0, len(UserPoints))
	for _, UserPoint := range UserPoints {
		response := &pb.UserPoint{
			XId:           UserPoint.Id.Hex(),
			UserId:        UserPoint.UserId,
			ReactionLevel: UserPoint.ReactionLevel,
			QaLevel:       UserPoint.QaLevel,
			QuestionCount: UserPoint.QuestionCount,
			AnswerCount:   UserPoint.AnswerCount,
			SolvedCount:   UserPoint.SolvedCount,
			CreatedAt:     timestamppb.New(UserPoint.CreateAt),
			UpdatedAt:     timestamppb.New(UserPoint.UpdatedAt),
		}
		cateList = append(cateList, response)
	}

	response := &pb.UserPointResponseList{
		UserPoints: cateList,
	}

	return response, nil
}

func (UserPointServer *UserPointServer) UpdateUserPoint(ctx context.Context, req *pb.UpdateUserPointRequest) (*pb.UserPointResponse, error) {

	UserPoint := &models.UpdateUserPoint{
		UserId:        req.GetUserId(),
		ReactionLevel: req.GetReactionLevel(),
		QaLevel:       req.GetQaLevel(),
		QuestionCount: req.GetQuestionCount(),
		AnswerCount:   req.GetAnswerCount(),
		SolvedCount:   req.GetSolvedCount(),
		UpdatedAt:     time.Now(),
	}

	updatedUserPoint, err := UserPointServer.UserPointService.UpdateUserPoint(UserPoint.UserId, UserPoint)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserPointResponse{
		UserPoint: &pb.UserPoint{
			XId:           updatedUserPoint.Id.Hex(),
			UserId:        updatedUserPoint.UserId,
			ReactionLevel: updatedUserPoint.ReactionLevel,
			QaLevel:       updatedUserPoint.QaLevel,
			QuestionCount: updatedUserPoint.QuestionCount,
			AnswerCount:   updatedUserPoint.AnswerCount,
			SolvedCount:   updatedUserPoint.SolvedCount,
			CreatedAt:     timestamppb.New(updatedUserPoint.CreateAt),
			UpdatedAt:     timestamppb.New(updatedUserPoint.UpdatedAt),
		},
	}
	return res, nil
}

func (UserPointServer *UserPointServer) DeleteUserPoint(ctx context.Context, req *pb.UserPointRequest) (*pb.DeleteUserPointResponse, error) {
	userId := req.GetId()

	err := UserPointServer.UserPointService.DeleteUserPoint(userId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteUserPointResponse{
		Success: true,
	}

	return res, nil
}

func (UserPointServer *UserPointServer) EvaluatePoints(ctx context.Context, req *pb.GetUserPointsRequest) (*pb.UserPointEvaluateResponse, error) {

	err := UserPointServer.UserPointService.EvaluatePoints()
	if err != nil {
		return nil, err
	}

	res := &pb.UserPointEvaluateResponse{
		Success: true,
	}
	return res, nil
}
