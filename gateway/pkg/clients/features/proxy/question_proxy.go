package features_proxy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/timestamp"
	badge_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/badge/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/models"
	question_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/config"

	user_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateQuestionResponse struct {
	QuestionRes *question_proto.QuestionResponse
	CommentId   string
}

func CreateQuestion(ctx *gin.Context, qsc question_proto.QuestionServiceClient, csc question_proto.CommentServiceClient, usc user_proto.UserServiceClient, bsc badge_proto.UserPointServiceClient) (*CreateQuestionResponse, error) {
	config, err := config.LoadConfig()
	question := models.CreateQuestionRequest{}

	if err := ctx.BindJSON(&question); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := qsc.CreateQuestion(context.Background(), &question_proto.CreateQuestionRequest{
		UserId:      question.UserId,
		Title:       question.Title,
		LanguageIds: question.LanguageIds,
		TagIds:      question.TagIds,
		UserIds:     question.UserIds,
		IsDeleted:   &question.IsDeleted,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}

	if res != nil {
		userPoint, err := bsc.GetUserPoint(context.Background(), &badge_proto.UserPointRequest{
			Id: res.UserId,
		})

		if err != nil {
			ctx.AbortWithError(http.StatusNotFound, err)
			return nil, err
		}

		if userPoint != nil {
			bsc.UpdateUserPoint(context.Background(), &badge_proto.UpdateUserPointRequest{
				UserId:        userPoint.UserPoint.UserId,
				ReactionLevel: userPoint.UserPoint.ReactionLevel,
				QaLevel:       userPoint.UserPoint.QaLevel,
				QuestionCount: userPoint.UserPoint.QuestionCount + 1,
				AnswerCount:   userPoint.UserPoint.AnswerCount,
				SolvedCount:   userPoint.UserPoint.SolvedCount,
			})
		}

		userData, err := usc.GetUser(context.Background(), &user_proto.UserRequest{
			XId: res.UserId,
		})

		requestBodyData := map[string]interface{}{
			"user_name": userData.Name,
			"title": question.Title,
			"question_id": res.XId,
		}

		requestBodyJSON, err := json.Marshal(requestBodyData)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return nil, err
		}

		if question.IsMentioned {
			BOT_URL := config.BOT_URL
			go func () { http.Post(BOT_URL+"/mentioned-with-language", "application/json", bytes.NewBuffer(requestBodyJSON)) }()
			// go func () { http.Post("http://127.0.0.1:5000/mentioned-with-language", "application/json", bytes.NewBuffer(requestBodyJSON)) }()
		}
	}

	commentRes, err := csc.CreateComment(context.Background(), &question_proto.CreateCommentRequest{
		UserId:      question.UserId,
		QuestionId:  res.XId,
		ParentId:    "",
		Sort:        1,
		Description: question.Description,
		VoteCount:   0,
		IsSolution:  false,
		IsDeleted:   false,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}
	user, _ := usc.GetUser(context.Background(), &user_proto.UserRequest{
		XId: res.UserId,
	})

	var role int32
	switch user.Role {
	case "manager":
		role = 1
	case "bse":
		role = 2
	case "leader":
		role = 3
	case "sub leader":
		role = 4
	case "senior":
		role = 5
	case "junior":
		role = 6
	default:
		ctx.AbortWithError(http.StatusBadRequest, status.Errorf(codes.InvalidArgument, "Invalid role value"))
		return nil, status.Errorf(codes.InvalidArgument, "Invalid role value")
	}

	dobTime, err := convertProtoTimestampToTime(user.Dob)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	fmt.Println("lpp", res.CreatedAt)
	_, er := usc.UpdateUser(context.Background(), &user_proto.UpdateUserRequest{
		XId:           question.UserId,
		Name:          &user.Name,
		Email:         &user.Email,
		Profile:       &user.Profile,
		Phone:         &user.Phone,
		DisplayName:   &user.DisplayName,
		Role:          &role,
		DepartmentId:  &user.DepartmentId,
		TeamId:        &user.TeamId,
		AboutMe:       &user.AboutMe,
		Address:       &user.Address,
		MailSubscribe: &user.MailSubscribe,
		Dob:           &timestamp.Timestamp{Seconds: dobTime.Unix(), Nanos: int32(dobTime.Nanosecond())}, // Convert time.Time to *timestamp.Timestamp
		LastPost:      res.CreatedAt,
	})

	if er != nil {
		fmt.Println("Error updating user last post time:", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}
	example, err := &CreateQuestionResponse{
		QuestionRes: res,
		CommentId:   commentRes.XId,
	}, nil

	return example, nil
}

func convertProtoTimestampToTime(ts *timestamp.Timestamp) (time.Time, error) {
	return time.Unix(ts.Seconds, int64(ts.Nanos)).UTC(), nil
}

func GetQuestions(ctx *gin.Context, qsc question_proto.QuestionServiceClient, csc question_proto.CommentServiceClient) []*question_proto.Question {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	sort := ctx.Query("sort")
	order := ctx.Query("order")

	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	stream, err := qsc.GetQuestions(context.Background(), &question_proto.GetQuestionsRequest{
		Page:  &convertedPage,
		Limit: &convertedLimit,
		Sort:  &sort,
		Order: &order,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var questions []*question_proto.Question

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}

		questions = append(questions, res)
	}

	return questions
}

func GetQuestion(ctx *gin.Context, qsc question_proto.QuestionServiceClient) (*question_proto.QuestionResponse, error) {
	questionId := ctx.Param("question_id")

	res, err := qsc.GetQuestion(context.Background(), &question_proto.QuestionRequest{
		XId: questionId,
	})

	return res, err
}

func GetQuestionById(ctx *gin.Context, qsc question_proto.QuestionServiceClient, csc question_proto.CommentServiceClient) (*models.QuestionDetail, error) {
	questionId := ctx.Param("question_id")

	res, err := qsc.GetQuestion(context.Background(), &question_proto.QuestionRequest{
		XId: questionId,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return nil, err
		}
		ctx.AbortWithError(http.StatusNotFound, err)
		return nil, err
	}

	stream, err := csc.GetCommentByQuestionId(context.Background(), &question_proto.QuestionIdRequest{
		QuestionId: questionId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}

	var comments []*question_proto.Comment
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return nil, err
		}

		comments = append(comments, res)
	}

	questionDetail := &models.QuestionDetail{
		Id:          res.GetXId(),
		UserId:      res.GetUserId(),
		Title:       res.GetTitle(),
		LanguageIds: res.GetLanguageIds(),
		TagIds:      res.GetTagIds(),
		ViewCount:   res.GetViewCount(),
		VoteCount:   res.GetVoteCount(),
		ReplyCount:  res.GetReplyCount(),
		CreateAt:    res.GetCreatedAt().AsTime(),
		UpdatedAt:   res.GetUpdatedAt().AsTime(),
		Comments:    comments,
	}

	return questionDetail, err
}

func GetQuestionsByUserId(ctx *gin.Context, qsc question_proto.QuestionServiceClient, csc question_proto.CommentServiceClient) []*question_proto.Question {
	userId := ctx.Param("user_id")
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	sort := ctx.Query("sort")
	order := ctx.Query("order")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)
	stream, err := qsc.GetQuestionByUserId(context.Background(), &question_proto.QuestionResquestByUserId{
		UserId: userId,
		Page:   &convertedPage,
		Limit:  &convertedLimit,
		Sort:   &sort,
		Order:  &order,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}

	var questions []*question_proto.Question
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return nil
		}

		questions = append(questions, res)
	}

	return questions
}

func DeleteQuestion(ctx *gin.Context, qsc question_proto.QuestionServiceClient) (*question_proto.DeleteQuestionResponse, error) {
	questionId := ctx.Param("question_id")

	res, err := qsc.DeleteQuestion(context.Background(), &question_proto.QuestionRequest{
		XId: questionId,
	})

	return res, err
}

func UpdateQuestions(ctx *gin.Context, qsc question_proto.QuestionServiceClient, vsc question_proto.VoteServiceClient) (*question_proto.QuestionResponse, error) {
	questionId := ctx.Param("question_id")
	questionModel := models.UpdateQuestion{}

	if err := ctx.BindJSON(&questionModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := qsc.UpdateQuestion(context.Background(), &question_proto.UpdateQuestionRequest{
		XId:         questionId,
		Title:       &questionModel.Title,
		LanguageIds: questionModel.LanguageIds,
		TagIds:      questionModel.TagIds,
		ViewCount:   &questionModel.ViewCount,
		VoteCount:   &questionModel.VoteCount,
		ReplyCount:  &questionModel.ReplyCount,
		IsDeleted:   &questionModel.IsDeleted,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	return res, err
}

func FilterQuestion(ctx *gin.Context, qsc question_proto.QuestionServiceClient, csc question_proto.CommentServiceClient) []*question_proto.Question {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	sort := ctx.Query("sort")
	order := ctx.Query("order")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)
	questionModel := models.FilterQuestionRequest{}

	if err := ctx.BindJSON(&questionModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	stream, err := qsc.FilterQuestion(context.Background(), &question_proto.FilterQuestionRequest{
		Page:       &convertedPage,
		Limit:      &convertedLimit,
		LanguageId: questionModel.LanguageIds,
		TagId:      questionModel.TagIds,
		UserId:     questionModel.UserId,
		Tile:       &questionModel.Title,
		Order:      &order,
		Sort:       &sort,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var questions []*question_proto.Question
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}

		// comment := GetCommentByQuestionId(ctx, csc, res.XId)

		// for _, comment := range comment {
		// 	if comment.IsSolution.Value {
		// 		res.IsSolution = wrapperspb.Bool(true)
		// 	} else {
		// 		res.IsSolution = wrapperspb.Bool(false)
		// 	}
		// }

		questions = append(questions, res)
	}

	return questions
}

func GetFilteredQuestionCount(ctx *gin.Context, qsc question_proto.QuestionServiceClient) (*question_proto.QuestionCountResponse, error) {
	sort := ctx.Query("sort")
	order := ctx.Query("order")
	questionModel := models.FilterQuestionRequest{}

	if err := ctx.BindJSON(&questionModel); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	res, err := qsc.GetFilteredQuestionCount(context.Background(), &question_proto.FilterQuestionRequest{
		LanguageId: questionModel.LanguageIds,
		TagId:      questionModel.TagIds,
		UserId:     questionModel.UserId,
		Tile:       &questionModel.Title,
		Order:      &order,
		Sort:       &sort,
	})

	return res, err
}

func GetQuestionCountAll(ctx *gin.Context, qsc question_proto.QuestionServiceClient) (*question_proto.QuestionCountResponse, error) {
	res, err := qsc.GetQuestionCountAll(context.Background(), &question_proto.GetQuestionCountRequest{})

	return res, err
}
