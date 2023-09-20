package features_proxy

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/models"
	bookmark_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	vote_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	user_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/pb"
)

func CreateBookmark(ctx *gin.Context, bsc bookmark_proto.BookmarkServiceClient) (*bookmark_proto.BookmarkResponse, error) {
	bookmark := models.CreateBookmarkRequest{}

	if err := ctx.BindJSON(&bookmark); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := bsc.CreateBookmark(context.Background(), &bookmark_proto.CreateBookmarkRequest{
		UserId:     bookmark.User_Id,
		CommentId:  bookmark.Comment_Id,
		QuestionId: bookmark.Question_Id,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	return res, err
}

func GetBookmarks(ctx *gin.Context, bsc bookmark_proto.BookmarkServiceClient) ([]*bookmark_proto.Bookmark, error) {
	bookmarkModel := models.BookmarkRequestByPage{}

	stream, err := bsc.GetBookmarks(context.Background(), &bookmark_proto.GetBookmarksRequest{
		Page:  &bookmarkModel.Page,
		Limit: &bookmarkModel.Limit,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var votes []*bookmark_proto.Bookmark

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}

		votes = append(votes, res)
	}

	return votes, err
}

func GetBookmark(ctx *gin.Context, bsc bookmark_proto.BookmarkServiceClient) (*bookmark_proto.BookmarkResponse, error) {
	bookmarkId := ctx.Param("bookmark_id")

	res, err := bsc.GetBookmark(context.Background(), &bookmark_proto.BookmarkRequest{
		XId: bookmarkId,
	})

	return res, err
}

func UpdateBookmark(ctx *gin.Context, bsc bookmark_proto.BookmarkServiceClient) (*bookmark_proto.BookmarkResponse, error) {
	bookmarkId := ctx.Param("bookmark_id")
	bookmarkModel := models.UpdateBookmark{}

	if err := ctx.BindJSON(&bookmarkModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := bsc.UpdateBookmark(context.Background(), &bookmark_proto.UpdateBookmarkRequest{
		XId:        bookmarkId,
		UserId:     &bookmarkModel.User_Id,
		CommentId:  &bookmarkModel.Comment_Id,
		QuestionId: &bookmarkModel.Question_Id,
	})

	return res, err
}

func DeleteBookmark(ctx *gin.Context, bsc bookmark_proto.BookmarkServiceClient) (*bookmark_proto.DeleteBookmarkResponse, error) {
	bookmarkId := ctx.Param("bookmark_id")

	res, err := bsc.DeleteBookmark(context.Background(), &bookmark_proto.BookmarkRequest{
		XId: bookmarkId,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	return res, err
}

type CustomCommentResponse struct {
	*bookmark_proto.CommentResponse
	UserProfile string `json:"user_profile"`
	UserName    string `json:"display_name"`
}

func GetBookmarksByUserId(ctx *gin.Context, bsc bookmark_proto.BookmarkServiceClient, csc bookmark_proto.CommentServiceClient, usc user_proto.UserServiceClient) []*bookmark_proto.CommentResponse {
	userId := ctx.Param("user_id")
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	stream, err := bsc.GetBookmarksByUserId(context.Background(), &bookmark_proto.BookmarkRequestByUserId{
		UserId: userId,
		Page:   &convertedPage,
		Limit:  &convertedLimit,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}

	var commentsData []*vote_proto.CommentResponse
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return nil
		}
		co, _ := csc.GetComment(context.Background(), &vote_proto.CommentRequest{
			XId: res.CommentId,
		})

		commentData := &vote_proto.CommentResponse{
			XId:         co.XId,
			UserId:      co.UserId,
			QuestionId:  co.QuestionId,
			ParentId:    co.ParentId,
			Sort:        co.Sort,
			Description: co.Description,
			VoteCount:   co.VoteCount,
			IsSolution:  co.IsSolution,
			IsDeleted:   co.IsDeleted,
			CreatedAt:   co.CreatedAt,
			UpdatedAt:   co.UpdatedAt,
		}

		commentsData = append(commentsData, commentData)
	}

	return commentsData
}

func GetBookmarksByUserIdQuestionId(ctx *gin.Context, bsc bookmark_proto.BookmarkServiceClient) []*bookmark_proto.Bookmark {
	userId := ctx.Param("user_id")
	questionId := ctx.Param("question_id")
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	stream, err := bsc.GetBookmarksByUserIdQuestionId(context.Background(), &bookmark_proto.BookmarkRequestByUserIdQuestionId{
		UserId:     userId,
		QuestionId: questionId,
		Page:       &convertedPage,
		Limit:      &convertedLimit,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}

	var bookmarks []*bookmark_proto.Bookmark
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return nil
		}

		bookmarks = append(bookmarks, res)
	}

	return bookmarks
}
