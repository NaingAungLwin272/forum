package gapi

import (
	"context"
	"strings"
	"time"

	models "github.com/scm-dev1dev5/mtm-community-forum/noti_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/noti_rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (notiServer *NotiServer) CreateNoti(ctx context.Context, req *pb.CreateNotiRequest) (*pb.NotiResponse, error) {
	noti := &models.CreateNotiRequest{
		UserId:      req.GetUserId(),
		Type:        int(req.GetType()),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Link:        req.GetLink(),
		Status:      &req.Status,
	}

	newPost, err := notiServer.notiService.CreateNoti(noti)
	if err != nil {
		if strings.Contains(err.Error(), "noti already exists") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.NotiResponse{
		Noti: &pb.Noti{
			XId:         newPost.Id.Hex(),
			UserId:      newPost.UserId,
			Type:        int64(newPost.Type),
			Name:        newPost.Name,
			Description: newPost.Description,
			Link:        &newPost.Link,
			Status:      *newPost.Status,
			CreatedAt:   timestamppb.New(newPost.CreateAt),
			UpdatedAt:   timestamppb.New(newPost.UpdatedAt),
		},
	}
	return res, nil
}

func (notiServer *NotiServer) GetNoti(ctx context.Context, req *pb.NotiRequest) (*pb.NotiResponse, error) {
	notiId := req.GetXId()
	noti, err := notiServer.notiService.GetNoti(notiId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())

		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.NotiResponse{
		Noti: &pb.Noti{
			XId:         noti.Id.Hex(),
			UserId:      noti.UserId,
			Type:        int64(noti.Type),
			Name:        noti.Name,
			Description: noti.Description,
			Link:        &noti.Link,
			Status:      *noti.Status,
			CreatedAt:   timestamppb.New(noti.CreateAt),
			UpdatedAt:   timestamppb.New(noti.UpdatedAt),
		},
	}
	return res, nil
}

func (notiServer *NotiServer) GetNotis(req *pb.GetNotisRequest, stream pb.NotiService_GetNotisServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()

	notis, err := notiServer.notiService.GetNotis(int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, noti := range notis {
		stream.Send(&pb.Noti{
			XId:         noti.Id.Hex(),
			UserId:      noti.UserId,
			Type:        int64(noti.Type),
			Name:        noti.Name,
			Description: noti.Description,
			Link:        &noti.Link,
			Status:      *noti.Status,
			CreatedAt:   timestamppb.New(noti.CreateAt),
			UpdatedAt:   timestamppb.New(noti.UpdatedAt),
		})
	}

	return nil
}

func (notiServer *NotiServer) UpdateNoti(ctx context.Context, req *pb.UpdateNotiRequest) (*pb.NotiResponse, error) {
	notiId := req.GetXId()

	noti := &models.UpdateNoti{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Link:        req.GetLink(),
		Status:      req.GetStatus(),
		UpdatedAt:   time.Now(),
	}

	updatedNoti, err := notiServer.notiService.UpdateNoti(notiId, noti)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.NotiResponse{
		Noti: &pb.Noti{
			XId:         updatedNoti.Id.Hex(),
			UserId:      updatedNoti.UserId,
			Type:        int64(updatedNoti.Type),
			Name:        updatedNoti.Name,
			Description: updatedNoti.Description,
			Link:        &updatedNoti.Link,
			Status:      *updatedNoti.Status,
			CreatedAt:   timestamppb.New(updatedNoti.CreateAt),
			UpdatedAt:   timestamppb.New(updatedNoti.UpdatedAt),
		},
	}
	return res, nil
}

func (notiServer *NotiServer) DeleteNoti(ctx context.Context, req *pb.NotiRequest) (*pb.DeleteNotiResponse, error) {
	notiId := req.GetXId()

	if err := notiServer.notiService.DeleteNoti(notiId); err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteNotiResponse{
		Success: true,
	}

	return res, nil
}

func (notiServer *NotiServer) GetNotiByUserId(req *pb.UserIdRequest, stream pb.NotiService_GetNotiByUserIdServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()
	notiId := req.GetUserId()
	notis := notiServer.notiService.GetNotiByUserId(notiId, int(page), int(limit))

	for _, noti := range notis {
		stream.Send(&pb.Noti{
			XId:         noti.Id.Hex(),
			UserId:      noti.UserId,
			Type:        int64(noti.Type),
			Name:        noti.Name,
			Description: noti.Description,
			Link:        &noti.Link,
			Status:      *noti.Status,
			CreatedAt:   timestamppb.New(noti.CreateAt),
			UpdatedAt:   timestamppb.New(noti.UpdatedAt),
		})
	}

	return nil
}

func (notiServer *NotiServer) GetNotiCount(context context.Context, req *pb.NotiRequestByUserId) (*pb.NotiCountResponse, error) {
	userId := req.GetUserId()
	notis := notiServer.notiService.GetNotiCount(userId)

	res := &pb.NotiCountResponse{
		Count: int64(notis),
	}

	return res, nil
}

func (NotiServer *NotiServer) MarkAllNotiAsRead(context context.Context, req *pb.NotiRequestByUserId) (*pb.DeleteNotiResponse, error) {
	userId := req.GetUserId()
	err := NotiServer.notiService.MarkAllNotiAsRead(userId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	res := &pb.DeleteNotiResponse{
		Success: true,
	}

	return res, nil
}

func (notiServer *NotiServer) GetNotiForUserSummary(context context.Context, req *pb.NotiRequestByUserId) (*pb.NotiCountResponse, error) {
	userId := req.GetUserId()
	notis := notiServer.notiService.GetNotiForUserSummary(userId)

	res := &pb.NotiCountResponse{
		Count: int64(notis),
	}

	return res, nil
}
