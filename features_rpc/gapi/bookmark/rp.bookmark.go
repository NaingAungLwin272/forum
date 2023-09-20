package bookmark_gapi

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

func (bookmarkServer *BookmarkServer) CreateBookmark(ctx context.Context, req *pb.CreateBookmarkRequest) (*pb.BookmarkResponse, error) {
	bookmark := &models.CreateBookmarkRequest{
		User_Id:     req.GetUserId(),
		Comment_Id:  req.GetCommentId(),
		Question_Id: req.GetQuestionId(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	newBookmark, err := bookmarkServer.bookmarkService.CreateBookmark(bookmark)
	if err != nil {
		if strings.Contains(err.Error(), "commentid already exists") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.BookmarkResponse{
		XId:        newBookmark.Id.Hex(),
		UserId:     newBookmark.User_Id,
		CommentId:  newBookmark.Comment_Id,
		QuestionId: newBookmark.Question_Id,
		CreatedAt:  timestamppb.New(newBookmark.CreatedAt),
		UpdatedAt:  timestamppb.New(newBookmark.UpdatedAt),
	}

	return res, nil
}

func (bookmarkServer *BookmarkServer) GetBookmarks(req *pb.GetBookmarksRequest, stream pb.BookmarkService_GetBookmarksServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()

	bookmarks, err := bookmarkServer.bookmarkService.GetBookmarks(int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, bookmark := range bookmarks {
		stream.Send(&pb.Bookmark{
			XId:        bookmark.Id.Hex(),
			UserId:     bookmark.User_Id,
			CommentId:  bookmark.Comment_Id,
			QuestionId: bookmark.Question_Id,
			CreatedAt:  timestamppb.New(bookmark.CreatedAt),
			UpdatedAt:  timestamppb.New(bookmark.UpdatedAt),
		})
	}

	return nil
}

func (bookmarkServer *BookmarkServer) GetBookmark(context context.Context, req *pb.BookmarkRequest) (*pb.BookmarkResponse, error) {
	bookmarkId := req.GetXId()
	bookmark, err := bookmarkServer.bookmarkService.GetBookmark(bookmarkId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists...") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.BookmarkResponse{
		XId:        bookmark.Id.Hex(),
		UserId:     bookmark.User_Id,
		CommentId:  bookmark.Comment_Id,
		QuestionId: bookmark.Question_Id,
		CreatedAt:  timestamppb.New(bookmark.CreatedAt),
		UpdatedAt:  timestamppb.New(bookmark.UpdatedAt),
	}

	return res, nil
}

func (bookmarkServer *BookmarkServer) DeleteBookmark(context context.Context, req *pb.BookmarkRequest) (*pb.DeleteBookmarkResponse, error) {
	bookmarkId := req.GetXId()

	if err := bookmarkServer.bookmarkService.DeleteBookmark(bookmarkId); err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteBookmarkResponse{
		Success: true,
	}

	return res, nil
}

func (bookmarkServer *BookmarkServer) UpdateBookmark(context context.Context, req *pb.UpdateBookmarkRequest) (*pb.BookmarkResponse, error) {
	bookmarkId := req.GetXId()
	bookmark := &models.UpdateBookmark{
		User_Id:     req.GetUserId(),
		Comment_Id:  req.GetCommentId(),
		Question_Id: req.GetQuestionId(),
		UpdatedAt:   time.Now(),
	}

	updatedBookmark, err := bookmarkServer.bookmarkService.UpdateBookmark(bookmarkId, bookmark)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.BookmarkResponse{
		XId:        updatedBookmark.Id.Hex(),
		UserId:     updatedBookmark.User_Id,
		CommentId:  updatedBookmark.Comment_Id,
		QuestionId: updatedBookmark.Question_Id,
		CreatedAt:  timestamppb.New(updatedBookmark.CreatedAt),
		UpdatedAt:  timestamppb.New(updatedBookmark.UpdatedAt),
	}

	return res, nil
}

func (bookmarkServer *BookmarkServer) GetBookmarksByUserId(req *pb.BookmarkRequestByUserId, stream pb.BookmarkService_GetBookmarksByUserIdServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()
	userId := req.GetUserId()
	bookmarks, err := bookmarkServer.bookmarkService.GetBookmarksByUserId(userId, int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, bookmark := range bookmarks {
		stream.Send(&pb.Bookmark{
			XId:        bookmark.Id.Hex(),
			UserId:     bookmark.User_Id,
			CommentId:  bookmark.Comment_Id,
			QuestionId: bookmark.Question_Id,
			CreatedAt:  timestamppb.New(bookmark.CreatedAt),
			UpdatedAt:  timestamppb.New(bookmark.UpdatedAt),
		})
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	return nil
}

func (bookmarkServer *BookmarkServer) GetBookmarkCount(context context.Context, req *pb.BookmarkRequestByUserId) (*pb.BookmarkCountResponse, error) {
	userId := req.GetUserId()
	bookmarks := bookmarkServer.bookmarkService.GetBookmarkCount(userId)

	res := &pb.BookmarkCountResponse{
		Count: int64(bookmarks),
	}

	return res, nil
}

func (bookmarkServer *BookmarkServer) GetBookmarksByUserIdQuestionId(req *pb.BookmarkRequestByUserIdQuestionId, stream pb.BookmarkService_GetBookmarksByUserIdQuestionIdServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()
	userId := req.GetUserId()
	questionId := req.GetQuestionId()
	bookmarks, err := bookmarkServer.bookmarkService.GetBookmarksByUserIdQuestionId(userId, questionId, int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, bookmark := range bookmarks {
		stream.Send(&pb.Bookmark{
			XId:        bookmark.Id.Hex(),
			UserId:     bookmark.User_Id,
			CommentId:  bookmark.Comment_Id,
			QuestionId: bookmark.Question_Id,
			CreatedAt:  timestamppb.New(bookmark.CreatedAt),
			UpdatedAt:  timestamppb.New(bookmark.UpdatedAt),
		})
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	return nil
}
