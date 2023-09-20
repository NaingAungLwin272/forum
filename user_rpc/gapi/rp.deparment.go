package gapi

import (
	"context"
	"strings"
	"time"

	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/consts"
	models "github.com/scm-dev1dev5/mtm-community-forum/user_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (deparmentServer *DeparmentServer) CreateDeparment(ctx context.Context, req *pb.CreateDeparmentRequest) (*pb.DeparmentResponse, error) {
	deparment := &models.CreateDeparmentRequest{
		Name: req.GetName(),
	}
	newDeparment, err := deparmentServer.deparmentService.CreateDeparment(deparment)
	if err != nil {
		if strings.Contains(err.Error(), "department") {
			return nil, status.Errorf(codes.AlreadyExists, consts.DepartmentExists)
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	res := &pb.DeparmentResponse{
		XId:       newDeparment.Id.Hex(),
		Name:      newDeparment.Name,
		CreatedAt: timestamppb.New(newDeparment.CreateAt),
		UpdatedAt: timestamppb.New(newDeparment.UpdatedAt),
	}

	return res, nil
}

func (deparmentServer *DeparmentServer) GetDeparment(ctx context.Context, req *pb.DeparmentRequest) (*pb.DeparmentResponse, error) {
	deparmentId := req.GetXId()
	deparment, err := deparmentServer.deparmentService.GetDeparment(deparmentId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())

		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	res := &pb.DeparmentResponse{
		XId:       deparment.Id.Hex(),
		Name:      deparment.Name,
		CreatedAt: timestamppb.New(deparment.CreateAt),
		UpdatedAt: timestamppb.New(deparment.UpdatedAt),
	}
	return res, nil
}

func (deparmentServer *DeparmentServer) GetDeparments(req *pb.GetDeparmentsRequest, stream pb.DeparmentService_GetDeparmentsServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()
	deparments := deparmentServer.deparmentService.GetDeparmentList(int(page), int(limit))

	for _, deparment := range deparments {
		stream.Send(&pb.Deparment{
			XId:       deparment.Id.Hex(),
			Name:      deparment.Name,
			CreatedAt: timestamppb.New(deparment.CreateAt),
			UpdatedAt: timestamppb.New(deparment.UpdatedAt),
		})
	}
	return nil
}

func (deparmentServer *DeparmentServer) UpdateDeparment(ctx context.Context, req *pb.DeparmentUpdateRequest) (*pb.DeparmentResponse, error) {
	deparmentId := req.GetXId()
	deparment := &models.UpdateDeparment{
		Name:      req.GetName(),
		UpdatedAt: time.Now(),
	}
	updatedDeparment, err := deparmentServer.deparmentService.UpdateDeparment(deparmentId, deparment)
	if err != nil {
		if strings.Contains(err.Error(), "name_1") {
			return nil, status.Errorf(codes.AlreadyExists, consts.DepartmentExists)
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeparmentResponse{
		XId:       updatedDeparment.Id.Hex(),
		Name:      updatedDeparment.Name,
		CreatedAt: timestamppb.New(updatedDeparment.CreateAt),
		UpdatedAt: timestamppb.New(updatedDeparment.UpdatedAt),
	}
	return res, nil
}

func (deparmentServer *DeparmentServer) DeleteDeparment(ctx context.Context, req *pb.DeparmentRequest) (*pb.DeleteDeparmentResponse, error) {
	deparmentId := req.GetXId()
	err := deparmentServer.deparmentService.DeleteDeparment(deparmentId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	res := &pb.DeleteDeparmentResponse{
		Success: true,
	}
	return res, nil
}

func (deparmentServer *DeparmentServer) GetDepartmentCount(context context.Context, req *pb.GetDeparmentsRequest) (*pb.DepartmentCountResponse, error) {
	var page = req.GetPage()
	var limit = req.GetLimit()
	department := deparmentServer.deparmentService.GetDepartmentCount(int(page), int(limit))
	res := &pb.DepartmentCountResponse{
		Count: int64(department),
	}
	return res, nil
}
