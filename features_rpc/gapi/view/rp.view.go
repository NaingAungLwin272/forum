package view_gapi

import (
	"context"
	"strings"
	"time"

	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (viewServer *ViewServer) CreateView(ctx context.Context, req *pb.CreateViewRequest) (*pb.ViewResponse, error) {
	view := &models.CreateViewRequest{
		User_Id:     req.GetUserId(),
		Question_Id: req.GetQuestionId(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	newView, err := viewServer.viewService.CreateView(view)
	if err != nil {
		if strings.Contains(err.Error(), "questionid already exists") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ViewResponse{
		XId:        newView.Id.Hex(),
		UserId:     newView.User_Id,
		QuestionId: newView.Question_Id,
		CreatedAt:  timestamppb.New(newView.CreatedAt),
		UpdatedAt:  timestamppb.New(newView.UpdatedAt),
	}
	return res, nil
}

func (viewServer *ViewServer) GetViews(req *pb.GetViewsRequest, stream pb.ViewService_GetViewsServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()

	views, err := viewServer.viewService.GetViews(int(page), int(limit))

	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, view := range views {
		stream.Send(&pb.View{
			XId:        view.Id.Hex(),
			UserId:     view.User_Id,
			QuestionId: view.Question_Id,
			CreatedAt:  timestamppb.New(view.CreatedAt),
			UpdatedAt:  timestamppb.New(view.UpdatedAt),
		})
	}

	return nil
}

func (viewServer *ViewServer) GetView(context context.Context, req *pb.ViewRequest) (*pb.ViewResponse, error) {
	viewId := req.GetXId()

	view, err := viewServer.viewService.GetView(viewId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists...") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ViewResponse{
		UserId:     view.User_Id,
		QuestionId: view.Question_Id,
		CreatedAt:  timestamppb.New(view.CreatedAt),
		UpdatedAt:  timestamppb.New(view.UpdatedAt),
	}

	return res, nil
}

func (viewServer *ViewServer) DeleteView(context context.Context, req *pb.ViewRequest) (*pb.DeleteViewResponse, error) {
	viewId := req.GetXId()

	if err := viewServer.viewService.DeleteView(viewId); err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
	}

	res := &pb.DeleteViewResponse{
		Success: true,
	}

	return res, nil
}

func (viewServer *ViewServer) UpdateView(context context.Context, req *pb.UpdateViewRequest) (*pb.ViewResponse, error) {
	viewId := req.GetXId()

	view := &models.UpdateView{
		User_Id:     req.GetUserId(),
		Question_Id: req.GetQuestionId(),
		UpdatedAt:   time.Now(),
	}

	updatedView, err := viewServer.viewService.UpdateView(viewId, view)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ViewResponse{
		UserId:     updatedView.User_Id,
		QuestionId: updatedView.Question_Id,
		CreatedAt:  timestamppb.New(updatedView.CreatedAt),
		UpdatedAt:  timestamppb.New(updatedView.UpdatedAt),
	}

	return res, nil
}

func (viewServer *ViewServer) GetViewsByUserId(req *pb.ViewRequestByUserId, stream pb.ViewService_GetViewsByUserIdServer) error {
	userId := req.GetUserId()
	views, err := viewServer.viewService.GetViewsByUserId(userId)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, view := range views {
		stream.Send(&pb.View{
			XId:        view.Id.Hex(),
			UserId:     view.User_Id,
			QuestionId: view.Question_Id,
			CreatedAt:  timestamppb.New(view.CreatedAt),
			UpdatedAt:  timestamppb.New(view.UpdatedAt),
		})
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	return nil
}

func (viewServer *ViewServer) GetViewsByUserIdQuestionId(req *pb.ViewRequestByUserIdQuestionId, stream pb.ViewService_GetViewsByUserIdQuestionIdServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()
	userId := req.GetUserId()
	questionId := req.GetQuestionId()
	views, err := viewServer.viewService.GetViewsByUserIdQuestionId(userId, questionId, int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, view := range views {
		stream.Send(&pb.View{
			XId:        view.Id.Hex(),
			UserId:     view.User_Id,
			QuestionId: view.Question_Id,
			CreatedAt:  timestamppb.New(view.CreatedAt),
			UpdatedAt:  timestamppb.New(view.UpdatedAt),
		})
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	return nil
}
