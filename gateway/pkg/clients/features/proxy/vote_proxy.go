package features_proxy

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/models"
	vote_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"

	// features_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/proxy"
	mail_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/mail/pb"
	noti_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/noti/pb"
	user_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/config"
)

func CreateVote(ctx *gin.Context, vsc vote_proto.VoteServiceClient, qsc vote_proto.QuestionServiceClient, csc vote_proto.CommentServiceClient, nsc noti_proto.NotiServiceClient, usc user_proto.UserServiceClient, masc mail_proto.MailServiceClient) (*vote_proto.VoteResponse, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at configs", err)
	}
	baseUrl := config.FRONT_END_URL
	voteModel := models.CreateVoteRequest{}

	if err := ctx.BindJSON(&voteModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)

		return nil, err
	}

	res, err := vsc.CreateVote(context.Background(), &vote_proto.CreateVoteRequest{
		UserId:     voteModel.User_Id,
		CommentId:  voteModel.Comment_Id,
		QuestionId: voteModel.Question_Id,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	currentCommentId := res.CommentId
	commentRes, err := csc.GetComment(context.Background(), &vote_proto.CommentRequest{
		XId: currentCommentId,
	})

	queRes, _ := qsc.GetQuestion(context.Background(), &vote_proto.QuestionRequest{
		XId: res.QuestionId,
	})
	ureslog, _ := usc.GetUser(context.Background(), &user_proto.UserRequest{
		XId: res.UserId,
	})

	commentUserId := commentRes.UserId
	link := fmt.Sprintf("%s/qa-detail/%s#%s", baseUrl, res.QuestionId, res.CommentId)

	msg := fmt.Sprintf("<span class=\"mm\">%s</span> voted at <span class=\"mm\">%s</span> ", ureslog.DisplayName, queRes.Title)
	if res.UserId != commentRes.UserId {

		response, err := nsc.CreateNoti(context.Background(), &noti_proto.CreateNotiRequest{
			UserId:      commentUserId,
			Type:        5,
			Name:        "Vote",
			Description: msg,
			Link:        &link,
			Status:      true,
		})
		fmt.Println(err, "err.....")
		fmt.Println(response, "response....")

		NotiFunc(ctx, voteModel.NotiToken, "Vote", "You Got A Vote")
	}

	ures, _ := usc.GetUser(context.Background(), &user_proto.UserRequest{
		XId: commentRes.UserId,
	})

	if ures.MailSubscribe {
		displayName := ures.DisplayName
		masc.SendMail(context.Background(), &mail_proto.MailRequest{
			Email:   ures.Email,
			Type:    5,
			Link:    &link,
			Subject: fmt.Sprintf("You got vote - %s", displayName),
		})
	}

	getVoteResultByQuestionId, _ := qsc.GetQuestion(context.Background(), &vote_proto.QuestionRequest{
		XId: res.QuestionId,
	})

	getVoteResultByCommentId, _ := csc.GetComment(context.Background(), &vote_proto.CommentRequest{
		XId: res.CommentId,
	})

	voteCountByCommentId := getVoteResultByCommentId.VoteCount + 1
	csc.UpdateComment(context.Background(), &vote_proto.UpdateCommentRequest{
		XId:        getVoteResultByCommentId.XId,
		VoteCount:  &voteCountByCommentId,
		IsSolution: &getVoteResultByCommentId.IsSolution.Value,
	})

	voteCountByQuestionId := getVoteResultByQuestionId.VoteCount + 1
	qsc.UpdateQuestion(context.Background(), &vote_proto.UpdateQuestionRequest{
		XId:        getVoteResultByQuestionId.XId,
		VoteCount:  &voteCountByQuestionId,
		ReplyCount: &getVoteResultByQuestionId.ReplyCount,
		ViewCount:  &getVoteResultByQuestionId.ViewCount,
	})

	return res, err
}

func removeHTMLTags(input string) string {
	re := regexp.MustCompile("<[^>]*>")
	return re.ReplaceAllString(input, "")
}

func GetVotes(ctx *gin.Context, vsc vote_proto.VoteServiceClient) ([]*vote_proto.Vote, error) {
	voteModel := models.VoteRequestByPage{}

	stream, err := vsc.GetVotes(context.Background(), &vote_proto.GetVotesRequest{
		Page:  &voteModel.Page,
		Limit: &voteModel.Limit,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var votes []*vote_proto.Vote

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

func GetVote(ctx *gin.Context, vsc vote_proto.VoteServiceClient) (*vote_proto.VoteResponse, error) {
	voteId := ctx.Param("vote_id")

	res, err := vsc.GetVote(context.Background(), &vote_proto.VoteRequest{
		XId: voteId,
	})

	return res, err
}

func UpdateVotes(ctx *gin.Context, vsc vote_proto.VoteServiceClient) (*vote_proto.VoteResponse, error) {
	voteId := ctx.Param("vote_id")
	voteModel := models.UpdateVote{}

	if err := ctx.BindJSON(&voteModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := vsc.UpdateVote(context.Background(), &vote_proto.UpdateVoteRequest{
		XId:        voteId,
		UserId:     &voteModel.User_Id,
		CommentId:  &voteModel.Comment_Id,
		QuestionId: &voteModel.Question_Id,
	})

	return res, err
}

func DeleteVote(ctx *gin.Context, vsc vote_proto.VoteServiceClient, qsc vote_proto.QuestionServiceClient, csc vote_proto.CommentServiceClient) (*vote_proto.DeleteVoteResponse, error) {
	voteId := ctx.Param("vote_id")

	fmt.Println(voteId, "voteId,....")
	getVote, _ := vsc.GetVote(context.Background(), &vote_proto.VoteRequest{
		XId: voteId,
	})

	res, err := vsc.DeleteVote(context.Background(), &vote_proto.VoteRequest{
		XId: voteId,
	})

	getVoteResultByQuestionId, _ := qsc.GetQuestion(context.Background(), &vote_proto.QuestionRequest{
		XId: getVote.QuestionId,
	})

	getVoteResultByCommentId, _ := csc.GetComment(context.Background(), &vote_proto.CommentRequest{
		XId: getVote.CommentId,
	})

	voteCountByCommentId := getVoteResultByCommentId.VoteCount - 1
	csc.UpdateComment(context.Background(), &vote_proto.UpdateCommentRequest{
		XId:        getVoteResultByCommentId.XId,
		VoteCount:  &voteCountByCommentId,
		IsSolution: &getVoteResultByCommentId.IsSolution.Value,
	})

	voteCountByQuestionId := getVoteResultByQuestionId.VoteCount - 1
	qsc.UpdateQuestion(context.Background(), &vote_proto.UpdateQuestionRequest{
		XId:        getVoteResultByQuestionId.XId,
		VoteCount:  &voteCountByQuestionId,
		ReplyCount: &getVoteResultByQuestionId.ReplyCount,
		ViewCount:  &getVoteResultByQuestionId.ViewCount,
	})

	return res, err
}

func GetVotesByUserId(ctx *gin.Context, vsc vote_proto.VoteServiceClient, cscc vote_proto.CommentServiceClient, usc user_proto.UserServiceClient) []*vote_proto.CommentResponse {
	userId := ctx.Param("user_id")
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	stream, err := vsc.GetVotesByUserId(context.Background(), &vote_proto.VoteRequestByUserId{
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
		co, _ := cscc.GetComment(context.Background(), &vote_proto.CommentRequest{
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

func GetVotesByUserIdQuestionId(ctx *gin.Context, vsc vote_proto.VoteServiceClient) []*vote_proto.Vote {
	userId := ctx.Param("user_id")
	questionId := ctx.Param("question_id")
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	stream, err := vsc.GetVotesByUserIdQuestionId(context.Background(), &vote_proto.VoteRequestByUserIdQuestionId{
		UserId:     userId,
		QuestionId: questionId,
		Page:       &convertedPage,
		Limit:      &convertedLimit,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}

	var votes []*vote_proto.Vote
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return nil
		}

		votes = append(votes, res)
	}

	return votes
}
