package vote_gapi

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

func (voteServer *VoteServer) CreateVote(ctx context.Context, req *pb.CreateVoteRequest) (*pb.VoteResponse, error) {
	vote := &models.CreateVoteRequest{
		User_Id:     req.GetUserId(),
		Comment_Id:  req.GetCommentId(),
		Question_Id: req.GetQuestionId(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	newVote, err := voteServer.voteService.CreateVote(vote)
	if err != nil {
		if strings.Contains(err.Error(), "commentid already exists") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.VoteResponse{
		XId:        newVote.Id.Hex(),
		UserId:     newVote.User_Id,
		CommentId:  newVote.Comment_Id,
		QuestionId: newVote.Question_Id,
		CreatedAt:  timestamppb.New(newVote.CreatedAt),
		UpdatedAt:  timestamppb.New(newVote.UpdatedAt),
	}

	return res, nil
}

func (voteServer *VoteServer) GetVotes(req *pb.GetVotesRequest, stream pb.VoteService_GetVotesServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()

	votes, err := voteServer.voteService.GetVotes(int(page), int(limit))

	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, vote := range votes {
		stream.Send(&pb.Vote{
			XId:        vote.Id.Hex(),
			UserId:     vote.User_Id,
			CommentId:  vote.Comment_Id,
			QuestionId: vote.Question_Id,
			CreatedAt:  timestamppb.New(vote.CreatedAt),
			UpdatedAt:  timestamppb.New(vote.UpdatedAt),
		})
	}

	return nil
}

func (voteServer *VoteServer) GetVote(context context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	voteId := req.GetXId()

	vote, err := voteServer.voteService.GetVote(voteId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists...") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.VoteResponse{
		XId:        vote.Id.Hex(),
		UserId:     vote.User_Id,
		CommentId:  vote.Comment_Id,
		QuestionId: vote.Question_Id,
		CreatedAt:  timestamppb.New(vote.CreatedAt),
		UpdatedAt:  timestamppb.New(vote.UpdatedAt),
	}

	return res, nil
}

func (voteServer *VoteServer) DeleteVote(context context.Context, req *pb.VoteRequest) (*pb.DeleteVoteResponse, error) {
	voteId := req.GetXId()

	if err := voteServer.voteService.DeleteVote(voteId); err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
	}

	res := &pb.DeleteVoteResponse{
		Success: true,
	}

	return res, nil
}

func (voteServer *VoteServer) UpdateVote(context context.Context, req *pb.UpdateVoteRequest) (*pb.VoteResponse, error) {
	voteId := req.GetXId()

	vote := &models.UpdateVote{
		User_Id:     req.GetUserId(),
		Comment_Id:  req.GetCommentId(),
		Question_Id: req.GetQuestionId(),
		UpdatedAt:   time.Now(),
	}

	updatedVote, err := voteServer.voteService.UpdateVote(voteId, vote)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.VoteResponse{
		XId:        updatedVote.Id.Hex(),
		UserId:     updatedVote.User_Id,
		CommentId:  updatedVote.Comment_Id,
		QuestionId: updatedVote.Question_Id,
		CreatedAt:  timestamppb.New(updatedVote.CreatedAt),
		UpdatedAt:  timestamppb.New(updatedVote.UpdatedAt),
	}

	return res, nil
}

func (voteServer *VoteServer) GetVotesByUserId(req *pb.VoteRequestByUserId, stream pb.VoteService_GetVotesByUserIdServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()
	userId := req.GetUserId()
	votes, err := voteServer.voteService.GetVotesByUserId(userId, int(page), int(limit))

	if err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}

	for _, vote := range votes {
		stream.Send(&pb.Vote{
			XId:        vote.Id.Hex(),
			UserId:     vote.User_Id,
			CommentId:  vote.Comment_Id,
			QuestionId: vote.Question_Id,
			CreatedAt:  timestamppb.New(vote.CreatedAt),
			UpdatedAt:  timestamppb.New(vote.UpdatedAt),
		})
	}

	return nil
}

func (voteServer *VoteServer) GetVoteCount(context context.Context, req *pb.VoteRequestByUserId) (*pb.VoteCountResponse, error) {
	userId := req.GetUserId()
	votes := voteServer.voteService.GetVoteCount(userId)

	res := &pb.VoteCountResponse{
		Count: int64(votes),
	}

	return res, nil
}

func (voteServer *VoteServer) GetVotesByUserIdQuestionId(req *pb.VoteRequestByUserIdQuestionId, stream pb.VoteService_GetVotesByUserIdQuestionIdServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()
	userId := req.GetUserId()
	questionId := req.GetQuestionId()
	votes, err := voteServer.voteService.GetVotesByUserIdQuestionId(userId, questionId, int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, vote := range votes {
		stream.Send(&pb.Vote{
			XId:        vote.Id.Hex(),
			UserId:     vote.User_Id,
			CommentId:  vote.Comment_Id,
			QuestionId: vote.Question_Id,
			CreatedAt:  timestamppb.New(vote.CreatedAt),
			UpdatedAt:  timestamppb.New(vote.UpdatedAt),
		})
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	return nil
}
