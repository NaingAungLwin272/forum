package comment_gapi

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

func (commentServer *CommentServer) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error) {
	comment := &models.CreateCommentRequest{
		User_Id:     req.GetUserId(),
		Question_Id: req.GetQuestionId(),
		Parent_Id:   req.GetParentId(),
		Sort:        req.GetSort(),
		Description: req.GetDescription(),
		Vote_Count:  req.GetVoteCount(),
		Is_Solution: req.IsSolution,
		Is_Deleted:  req.GetIsDeleted(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	newComment, _ := commentServer.commentService.CreateComment(comment)

	res := &pb.CommentResponse{
		XId:         newComment.Id.Hex(),
		UserId:      newComment.User_Id,
		QuestionId:  newComment.Question_Id,
		ParentId:    newComment.Parent_Id,
		Sort:        newComment.Sort,
		Description: newComment.Description,
		VoteCount:   newComment.Vote_Count,
		IsSolution:  wrapperspb.Bool(newComment.Is_Solution),
		IsDeleted:   wrapperspb.Bool(newComment.Is_Deleted),
		CreatedAt:   timestamppb.New(newComment.CreatedAt),
		UpdatedAt:   timestamppb.New(newComment.UpdatedAt),
	}

	return res, nil
}

func (commentServer *CommentServer) GetComments(req *pb.GetCommentsRequest, stream pb.CommentService_GetCommentsServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()

	comments, err := commentServer.commentService.GetComments(int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, comment := range comments {
		stream.Send(&pb.Comment{
			XId:         comment.Id.Hex(),
			UserId:      comment.User_Id,
			QuestionId:  comment.Question_Id,
			ParentId:    comment.Parent_Id,
			Sort:        comment.Sort,
			Description: comment.Description,
			VoteCount:   comment.Vote_Count,
			IsSolution:  wrapperspb.Bool(comment.Is_Solution),
			IsDeleted:   wrapperspb.Bool(comment.Is_Deleted),
			CreatedAt:   timestamppb.New(comment.CreatedAt),
			UpdatedAt:   timestamppb.New(comment.UpdatedAt),
		})
	}

	return nil
}

func (commentServer *CommentServer) GetComment(context context.Context, req *pb.CommentRequest) (*pb.CommentResponse, error) {
	commentId := req.GetXId()

	comment, err := commentServer.commentService.GetComment(commentId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists...") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if comment.Is_Deleted {
		return nil, status.Errorf(codes.NotFound, "Comment not found")
	}

	res := &pb.CommentResponse{
		XId:         comment.Id.Hex(),
		UserId:      comment.User_Id,
		QuestionId:  comment.Question_Id,
		ParentId:    comment.Parent_Id,
		Sort:        comment.Sort,
		Description: comment.Description,
		VoteCount:   comment.Vote_Count,
		IsSolution:  wrapperspb.Bool(comment.Is_Solution),
		IsDeleted:   wrapperspb.Bool(comment.Is_Deleted),
		CreatedAt:   timestamppb.New(comment.CreatedAt),
		UpdatedAt:   timestamppb.New(comment.UpdatedAt),
	}

	return res, nil
}

func (commentServer *CommentServer) DeleteComment(context context.Context, req *pb.CommentRequest) (*pb.DeleteCommentResponse, error) {
	commentId := req.GetXId()

	comment, err := commentServer.commentService.GetComment(commentId)
	if comment.Is_Deleted {
		return nil, status.Errorf(codes.NotFound, "Comment not found")
	}
	delete_err := commentServer.commentService.DeleteComment(commentId)
	if delete_err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteCommentResponse{
		Success: true,
	}

	return res, nil
}

func (commentServer *CommentServer) UpdateComment(context context.Context, req *pb.UpdateCommentRequest) (*pb.CommentResponse, error) {
	commentId := req.GetXId()

	comment := &models.UpdateComment{
		Sort:        req.GetSort(),
		Description: req.GetDescription(),
		Vote_Count:  req.GetVoteCount(),
		Is_Solution: req.GetIsSolution(),
		Is_Deleted:  req.GetIsDeleted(),
		UpdatedAt:   time.Now(),
	}

	updatedComment, err := commentServer.commentService.UpdateComment(commentId, comment)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CommentResponse{
		XId:         updatedComment.Id.Hex(),
		UserId:      updatedComment.User_Id,
		QuestionId:  updatedComment.Question_Id,
		Sort:        updatedComment.Sort,
		Description: updatedComment.Description,
		VoteCount:   updatedComment.Vote_Count,
		IsSolution:  wrapperspb.Bool(updatedComment.Is_Solution),
		IsDeleted:   wrapperspb.Bool(updatedComment.Is_Deleted),
		CreatedAt:   timestamppb.New(updatedComment.CreatedAt),
		UpdatedAt:   timestamppb.New(updatedComment.UpdatedAt),
	}

	return res, nil
}

func (commentServer CommentServer) GetCommentByQuestionId(req *pb.QuestionIdRequest, stream pb.CommentService_GetCommentByQuestionIdServer) error {
	questionId := req.GetQuestionId()
	comments := commentServer.commentService.GetCommentByQuestionId(questionId)

	for _, comment := range comments {
		stream.Send(&pb.Comment{
			XId:         comment.Id.Hex(),
			UserId:      comment.User_Id,
			QuestionId:  comment.Question_Id,
			ParentId:    comment.Parent_Id,
			Sort:        comment.Sort,
			Description: comment.Description,
			VoteCount:   comment.Vote_Count,
			IsSolution:  wrapperspb.Bool(comment.Is_Solution),
			IsDeleted:   wrapperspb.Bool(comment.Is_Deleted),
			CreatedAt:   timestamppb.New(comment.CreatedAt),
			UpdatedAt:   timestamppb.New(comment.UpdatedAt),
		})
	}

	return nil
}

func (commentServer *CommentServer) GetCommentsByUserId(req *pb.CommentResquestByUserId, stream pb.CommentService_GetCommentsByUserIdServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()
	userId := req.GetUserId()
	comments, err := commentServer.commentService.GetCommentsByUserId(userId, int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, comment := range comments {
		stream.Send(&pb.Comment{
			XId:         comment.Id.Hex(),
			UserId:      comment.User_Id,
			QuestionId:  comment.Question_Id,
			ParentId:    comment.Parent_Id,
			VoteCount:   comment.Vote_Count,
			IsSolution:  wrapperspb.Bool(comment.Is_Solution),
			IsDeleted:   wrapperspb.Bool(comment.Is_Deleted),
			Description: comment.Description,
			CreatedAt:   timestamppb.New(comment.CreatedAt),
			UpdatedAt:   timestamppb.New(comment.UpdatedAt),
		})
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	return nil
}

func (commentServer *CommentServer) GetAnswersByUserId(req *pb.CommentResquestByUserId, stream pb.CommentService_GetAnswersByUserIdServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()
	userId := req.GetUserId()
	comments, err := commentServer.commentService.GetAnswersByUserId(userId, int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, comment := range comments {
		stream.Send(&pb.Comment{
			XId:         comment.Id.Hex(),
			UserId:      comment.User_Id,
			QuestionId:  comment.Question_Id,
			ParentId:    comment.Parent_Id,
			VoteCount:   comment.Vote_Count,
			IsSolution:  wrapperspb.Bool(comment.Is_Solution),
			IsDeleted:   wrapperspb.Bool(comment.Is_Deleted),
			Description: comment.Description,
			CreatedAt:   timestamppb.New(comment.CreatedAt),
			UpdatedAt:   timestamppb.New(comment.UpdatedAt),
		})
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	return nil
}

func (commentServer *CommentServer) GetCommentsByUserIdWithSolved(req *pb.CommentResquestByUserId, stream pb.CommentService_GetCommentsByUserIdWithSolvedServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()
	userId := req.GetUserId()
	comments, err := commentServer.commentService.GetCommentsByUserIdWithSolved(userId, int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, comment := range comments {
		stream.Send(&pb.Comment{
			XId:         comment.Id.Hex(),
			UserId:      comment.User_Id,
			QuestionId:  comment.Question_Id,
			VoteCount:   comment.Vote_Count,
			IsSolution:  wrapperspb.Bool(comment.Is_Solution),
			IsDeleted:   wrapperspb.Bool(comment.Is_Deleted),
			Description: comment.Description,
			CreatedAt:   timestamppb.New(comment.CreatedAt),
			UpdatedAt:   timestamppb.New(comment.UpdatedAt),
		})
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	return nil
}

func (commentServer *CommentServer) GetCommentCount(context context.Context, req *pb.CommentResquestByUserId) (*pb.CommentCountResponse, error) {
	userId := req.GetUserId()
	comments := commentServer.commentService.GetCommentCount(userId)

	res := &pb.CommentCountResponse{
		Count: int64(comments),
	}

	return res, nil
}

func (commentServer *CommentServer) GetCommentCountBySolved(context context.Context, req *pb.CommentResquestByUserId) (*pb.CommentCountResponse, error) {
	userId := req.GetUserId()
	comments := commentServer.commentService.GetCommentCountBySolved(userId)

	res := &pb.CommentCountResponse{
		Count: int64(comments),
	}

	return res, nil
}

func (commentServer *CommentServer) GetCommentCountByQuestionIdSolved(context context.Context, req *pb.QuestionIdRequest) (*pb.CommentCountResponse, error) {
	userId := req.GetQuestionId()
	comments := commentServer.commentService.GetCommentCountByQuestionIdSolved(userId)

	res := &pb.CommentCountResponse{
		Count: int64(comments),
	}

	return res, nil
}
