package mention_gapi

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

func (mentionServer *MentionServer) CreateMention(ctx context.Context, req *pb.CreateMentionRequest) (*pb.MentionResponse, error) {
	mention := &models.CreateMentionRequest{
		User_Id:     req.GetUserId(),
		Comment_Id:  req.GetCommentId(),
		Question_Id: req.GetQuestionId(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	newMention, _ := mentionServer.mentionService.CreateMention(mention)

	res := &pb.MentionResponse{
		XId:        newMention.Id.Hex(),
		UserId:     newMention.User_Id,
		CommentId:  newMention.Comment_Id,
		QuestionId: newMention.Question_Id,
		CreatedAt:  timestamppb.New(newMention.CreatedAt),
	}

	return res, nil
}

func (mentionServer *MentionServer) GetMentions(req *pb.GetMentionsRequest, stream pb.MentionService_GetMentionsServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()

	mentions, err := mentionServer.mentionService.GetMentions(int(page), int(limit))

	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, mention := range mentions {
		stream.Send(&pb.Mention{
			XId:        mention.Id.Hex(),
			UserId:     mention.User_Id,
			CommentId:  mention.Comment_Id,
			QuestionId: mention.Question_Id,
			CreatedAt:  timestamppb.New(mention.CreatedAt),
			UpdatedAt:  timestamppb.New(mention.UpdatedAt),
		})
	}

	return nil
}

func (mentionServer *MentionServer) GetMention(context context.Context, req *pb.MentionRequest) (*pb.MentionResponse, error) {
	mentionId := req.GetXId()

	mention, err := mentionServer.mentionService.GetMention(mentionId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists...") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.MentionResponse{
		UserId:     mention.User_Id,
		CommentId:  mention.Comment_Id,
		QuestionId: mention.Question_Id,
		CreatedAt:  timestamppb.New(mention.CreatedAt),
		UpdatedAt:  timestamppb.New(mention.UpdatedAt),
	}

	return res, nil
}

func (mentionServer *MentionServer) DeleteMention(context context.Context, req *pb.MentionRequest) (*pb.DeleteMentionResponse, error) {
	mentionId := req.GetXId()

	if err := mentionServer.mentionService.DeleteMention(mentionId); err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
	}

	res := &pb.DeleteMentionResponse{
		Success: true,
	}

	return res, nil
}

func (mentionServer *MentionServer) UpdateMention(context context.Context, req *pb.UpdateMentionRequest) (*pb.MentionResponse, error) {
	mentionId := req.GetXId()

	mention := &models.UpdateMention{
		User_Id:     req.GetUserId(),
		Comment_Id:  req.GetCommentId(),
		Question_Id: req.GetQuestionId(),
		UpdatedAt:   time.Now(),
	}

	updatedMention, err := mentionServer.mentionService.UpdateMention(mentionId, mention)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.MentionResponse{
		UserId:     updatedMention.User_Id,
		CommentId:  updatedMention.Comment_Id,
		QuestionId: updatedMention.Question_Id,
		CreatedAt:  timestamppb.New(updatedMention.CreatedAt),
		UpdatedAt:  timestamppb.New(updatedMention.UpdatedAt),
	}

	return res, nil
}

func (mentionServer *MentionServer) GetMentionsByUserId(req *pb.MentionRequestByUserId, stream pb.MentionService_GetMentionsByUserIdServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()
	userId := req.GetUserId()
	mentions := mentionServer.mentionService.GetMentionsByUserId(userId, int(page), int(limit))

	for _, mention := range mentions {
		stream.Send(&pb.Mention{
			XId:        mention.Id.Hex(),
			UserId:     mention.User_Id,
			CommentId:  mention.Comment_Id,
			QuestionId: mention.Question_Id,
			CreatedAt:  timestamppb.New(mention.CreatedAt),
			UpdatedAt:  timestamppb.New(mention.UpdatedAt),
		})
	}

	return nil
}

func (mentionServer *MentionServer) GetMentionCount(context context.Context, req *pb.MentionRequestByUserId) (*pb.MentionCountResponse, error) {
	userId := req.GetUserId()
	comments := mentionServer.mentionService.GetMentionCount(userId)

	res := &pb.MentionCountResponse{
		Count: int64(comments),
	}

	return res, nil

}
