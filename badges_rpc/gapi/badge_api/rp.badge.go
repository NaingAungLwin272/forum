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

func (badgeServer *BadgeServer) CreateBadge(ctx context.Context, req *pb.CreateBadgeRequest) (*pb.BadgeResponse, error) {
	badge := &models.CreateBadgeRequest{
		Type:        int32(req.Type),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Level:       req.GetLevel(),
	}

	newBadge, err := badgeServer.badgeService.CreateBadge(badge)

	if err != nil {
		if strings.Contains(err.Error(), "name already exists") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.BadgeResponse{
		Badge: &pb.Badge{
			XId:         newBadge.Id.Hex(),
			Type:        pb.BadgeType(newBadge.Type),
			Name:        newBadge.Name,
			Description: newBadge.Description,
			Level:       newBadge.Level,
			CreatedAt:   timestamppb.New(newBadge.CreateAt),
			UpdatedAt:   timestamppb.New(newBadge.UpdatedAt),
		},
	}
	return res, nil
}

func (badgeServer *BadgeServer) GetBadge(ctx context.Context, req *pb.BadgeRequest) (*pb.BadgeResponse, error) {
	badgetId := req.GetId()

	badge, err := badgeServer.badgeService.GetBadge(badgetId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())

		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.BadgeResponse{
		Badge: &pb.Badge{
			XId:         badge.Id.Hex(),
			Type:        pb.BadgeType(badge.Type),
			Name:        badge.Name,
			Description: badge.Description,
			Level:       badge.Level,
			CreatedAt:   timestamppb.New(badge.CreateAt),
			UpdatedAt:   timestamppb.New(badge.UpdatedAt),
		},
	}
	return res, nil
}

func (badgeServer *BadgeServer) GetBadges(ctx context.Context, req *pb.GetBadgesRequest) (*pb.BadgeResponseList, error) {
	var page = req.GetPage()
	var limit = req.GetLimit()

	badges, err := badgeServer.badgeService.GetBadges(int(page), int(limit))
	if err != nil {
		return nil, err
	}

	cateList := make([]*pb.Badge, 0, len(badges))
	for _, badge := range badges {
		response := &pb.Badge{
			XId:         badge.Id.Hex(),
			Type:        pb.BadgeType(badge.Type),
			Name:        badge.Name,
			Description: badge.Description,
			Level:       badge.Level,
			CreatedAt:   timestamppb.New(badge.CreateAt),
			UpdatedAt:   timestamppb.New(badge.UpdatedAt),
		}
		cateList = append(cateList, response)
	}

	response := &pb.BadgeResponseList{
		Badges: cateList,
	}

	return response, nil
}

func (badgeServer *BadgeServer) UpdateBadge(ctx context.Context, req *pb.UpdateBadgeRequest) (*pb.BadgeResponse, error) {
	badgeId := req.GetXId()

	badge := &models.UpdateBadge{
		Type:        int32(req.GetType()),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Level:       req.GetLevel(),
		UpdatedAt:   time.Now(),
	}

	updatedBadge, err := badgeServer.badgeService.UpdateBadge(badgeId, badge)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.BadgeResponse{
		Badge: &pb.Badge{
			XId:         updatedBadge.Id.Hex(),
			Type:        pb.BadgeType(updatedBadge.Type),
			Name:        updatedBadge.Name,
			Description: updatedBadge.Description,
			Level:       updatedBadge.Level,
			CreatedAt:   timestamppb.New(updatedBadge.CreateAt),
			UpdatedAt:   timestamppb.New(updatedBadge.UpdatedAt),
		},
	}
	return res, nil
}

func (badgeServer *BadgeServer) DeleteBadge(ctx context.Context, req *pb.BadgeRequest) (*pb.DeleteBadgeResponse, error) {
	badgeId := req.GetId()

	err := badgeServer.badgeService.DeleteBadge(badgeId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteBadgeResponse{
		Success: true,
	}

	return res, nil
}
