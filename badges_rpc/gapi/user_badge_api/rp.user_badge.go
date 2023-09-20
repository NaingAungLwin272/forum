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

func (userBadgeServer *UserBadgeServer) CreateUserBadge(ctx context.Context, req *pb.CreateUserBadgeRequest) (*pb.UserBadgeResponse, error) {
	userBadge := &models.CreateUserBadgeRequest{
		User_Id:  req.GetUserId(),
		Badge_Id: req.GetBadgeId(),
	}

	newUserBadge, err := userBadgeServer.userBadgeService.CreateUserBadge(userBadge)

	if err != nil {
		if strings.Contains(err.Error(), "userId already exists") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserBadgeResponse{
		UserBadge: &pb.UserBadge{
			XId:       newUserBadge.Id.Hex(),
			BadgeId:   newUserBadge.Badge_Id,
			CreatedAt: timestamppb.New(newUserBadge.CreatedAt),
			UpdatedAt: timestamppb.New(newUserBadge.UpdatedAt),
		},
	}
	return res, nil
}

func (userBadgeServer *UserBadgeServer) GetUserBadge(ctx context.Context, req *pb.UserBadgeRequest) (*pb.UserBadgeResponse, error) {
	userBadgetId := req.GetUserId()
	BadgeId := req.GetBadgeId()

	userBadge, err := userBadgeServer.userBadgeService.GetUserBadge(userBadgetId, BadgeId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())

		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserBadgeResponse{
		UserBadge: &pb.UserBadge{
			XId:       userBadge.Id.Hex(),
			UserId:    userBadge.User_Id,
			BadgeId:   userBadge.Badge_Id,
			CreatedAt: timestamppb.New(userBadge.CreatedAt),
			UpdatedAt: timestamppb.New(userBadge.UpdatedAt),
		},
	}
	return res, nil
}

func (userBadgeServer *UserBadgeServer) GetUserBadges(ctx context.Context, req *pb.GetUserBadgesRequest) (*pb.UserBadgeResponseList, error) {
	var page = req.GetPage()
	var limit = req.GetLimit()

	userBadges, err := userBadgeServer.userBadgeService.GetUserBadges(int(page), int(limit))
	if err != nil {
		return nil, err
	}

	List := make([]*pb.UserBadge, 0, len(userBadges))
	for _, userBadge := range userBadges {
		response := &pb.UserBadge{
			XId:       userBadge.Id.Hex(),
			UserId:    userBadge.User_Id,
			BadgeId:   userBadge.Badge_Id,
			CreatedAt: timestamppb.New(userBadge.CreatedAt),
			UpdatedAt: timestamppb.New(userBadge.UpdatedAt),
		}
		List = append(List, response)
	}

	response := &pb.UserBadgeResponseList{
		UserBadges: List,
	}

	return response, nil
}

func (userBadgeServer *UserBadgeServer) GetUserBadgesOfUser(ctx context.Context, req *pb.GetUserBadgesOfUserRequest) (*pb.UserBadgeResponseList, error) {
	var userId = req.GetUserId()

	userBadges, err := userBadgeServer.userBadgeService.GetUserBadgesOfUser(userId)
	if err != nil {
		return nil, err
	}

	List := make([]*pb.UserBadge, 0, len(userBadges))
	for _, userBadge := range userBadges {
		response := &pb.UserBadge{
			XId:       userBadge.Id.Hex(),
			UserId:    userBadge.User_Id,
			BadgeId:   userBadge.Badge_Id,
			CreatedAt: timestamppb.New(userBadge.CreatedAt),
			UpdatedAt: timestamppb.New(userBadge.UpdatedAt),
		}
		List = append(List, response)
	}

	response := &pb.UserBadgeResponseList{
		UserBadges: List,
	}

	return response, nil
}

func (userBadgeServer *UserBadgeServer) UpdateUserBadge(ctx context.Context, req *pb.UpdateUserBadgeRequest) (*pb.UserBadgeResponse, error) {
	userId := req.GetUserId()

	userBadge := &models.UpdateUserBadge{
		User_Id:   req.UserId,
		Badge_Id:  *req.BadgeId,
		UpdatedAt: time.Now(),
	}

	updatedUserBadge, err := userBadgeServer.userBadgeService.UpdateUserBadge(userId, userBadge)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserBadgeResponse{
		UserBadge: &pb.UserBadge{
			XId:       updatedUserBadge.Id.Hex(),
			UserId:    updatedUserBadge.User_Id,
			BadgeId:   updatedUserBadge.Badge_Id,
			CreatedAt: timestamppb.New(updatedUserBadge.CreatedAt),
			UpdatedAt: timestamppb.New(updatedUserBadge.UpdatedAt),
		},
	}
	return res, nil
}

func (userBadgeServer *UserBadgeServer) DeleteUserBadge(ctx context.Context, req *pb.UserBadgeRequest) (*pb.DeleteUserBadgeResponse, error) {
	userId := req.GetUserId()

	err := userBadgeServer.userBadgeService.DeleteUserBadge(userId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteUserBadgeResponse{
		Success: true,
	}

	return res, nil
}

func (userBadgeServer *UserBadgeServer) GetBadgeCount(context context.Context, req *pb.BadgeRequestByUserId) (*pb.BadgeCountResponse, error) {
	userId := req.GetUserId()
	badges := userBadgeServer.userBadgeService.GetBadgeCount(userId)

	res := &pb.BadgeCountResponse{
		Count: int64(badges),
	}

	return res, nil
}
