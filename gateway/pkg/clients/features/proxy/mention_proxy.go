package features_proxy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/models"
	mention_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	vote_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	mail_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/mail/pb"
	noti_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/noti/pb"
	user_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/config"
)

func CreateMention(ctx *gin.Context, msc mention_proto.MentionServiceClient, nsc noti_proto.NotiServiceClient, usc user_proto.UserServiceClient, masc mail_proto.MailServiceClient, qsc vote_proto.QuestionServiceClient, cscc mention_proto.CommentServiceClient) (*mention_proto.MentionResponse, error) {
	config, err := config.LoadConfig()
	mention := models.CreateMentionRequest{}

	if err := ctx.BindJSON(&mention); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := msc.CreateMention(context.Background(), &mention_proto.CreateMentionRequest{
		UserId:     mention.User_Id,
		CommentId:  mention.Comment_Id,
		QuestionId: mention.Question_Id,
	})

	ures, _ := usc.GetUser(context.Background(), &user_proto.UserRequest{
		XId: res.UserId,
	})

	queRes, _ := qsc.GetQuestion(context.Background(), &vote_proto.QuestionRequest{
		XId: res.QuestionId,
	})
	commentLink := getCommentLink(mention)

	co, _ := cscc.GetComment(context.Background(), &vote_proto.CommentRequest{
		XId: res.CommentId,
	})
	mentionRes, _ := usc.GetUser(context.Background(), &user_proto.UserRequest{
		XId: co.UserId,
	})

	msg := fmt.Sprintf("<span class=\"mm\">%s</span> mentioned you at <span class=\"tt\">%s</span>", mentionRes.DisplayName, queRes.Title)
	nsc.CreateNoti(context.Background(), &noti_proto.CreateNotiRequest{
		UserId:      res.UserId,
		Type:        3,
		Name:        "Mention",
		Description: msg,
		Link:        &commentLink,
		Status:      true,
	})

	requestBodyData := map[string]interface{}{
		"user_id":           res.UserId,
		"created_user_name": mentionRes.DisplayName,
	}

	requestBodyJSON, err := json.Marshal(requestBodyData)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	fmt.Println(requestBodyData, "requestBodyData......")

	BOT_URL := config.BOT_URL

	go func() { http.Post(BOT_URL+"/mentioned", "application/json", bytes.NewBuffer(requestBodyJSON)) }()
	// go func () {http.Post("http://127.0.0.1:5000/mentioned", "application/json", bytes.NewBuffer(requestBodyJSON))} ()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	NotiFunc(ctx, mention.NotiToken, "Mention", "You Got A Mention")

	if ures.MailSubscribe {
		displayName := ures.DisplayName
		masc.SendMail(context.Background(), &mail_proto.MailRequest{
			Email:   ures.Email,
			Type:    3,
			Link:    &commentLink,
			Subject: fmt.Sprintf("Mention Mail- %s", displayName),
		})
	}
	return res, err
}

func getCommentLink(mention models.CreateMentionRequest) string {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at configs", err)
	}
	baseUrl := config.FRONT_END_URL
	return fmt.Sprintf("%s/qa-detail/%s#%s", baseUrl, mention.Question_Id, mention.Comment_Id)
}
func GetMentions(ctx *gin.Context, msc mention_proto.MentionServiceClient) ([]*mention_proto.Mention, error) {
	mentionModel := models.MentionRequestByPage{}

	stream, err := msc.GetMentions(context.Background(), &mention_proto.GetMentionsRequest{
		Page:  &mentionModel.Page,
		Limit: &mentionModel.Limit,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var mentions []*mention_proto.Mention

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}

		mentions = append(mentions, res)
	}

	return mentions, err
}

func GetMention(ctx *gin.Context, msc mention_proto.MentionServiceClient) (*mention_proto.MentionResponse, error) {
	mentionId := ctx.Param("mention_id")

	res, err := msc.GetMention(context.Background(), &mention_proto.MentionRequest{
		XId: mentionId,
	})

	return res, err
}

func UpdateMention(ctx *gin.Context, msc mention_proto.MentionServiceClient) (*mention_proto.MentionResponse, error) {
	mentionId := ctx.Param("mention_id")
	mentionModel := models.UpdateMention{}

	if err := ctx.BindJSON(&mentionModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := msc.UpdateMention(context.Background(), &mention_proto.UpdateMentionRequest{
		XId:        mentionId,
		UserId:     &mentionModel.User_Id,
		CommentId:  &mentionModel.Comment_Id,
		QuestionId: &mentionModel.Question_Id,
	})

	return res, err
}

func DeleteMention(ctx *gin.Context, msc mention_proto.MentionServiceClient) (*mention_proto.DeleteMentionResponse, error) {
	mentionId := ctx.Param("mention_id")

	res, err := msc.DeleteMention(context.Background(), &mention_proto.MentionRequest{
		XId: mentionId,
	})

	return res, err
}

func GetMentionsByUserId(ctx *gin.Context, msc mention_proto.MentionServiceClient, cscc mention_proto.CommentServiceClient, usc user_proto.UserServiceClient) []*mention_proto.CommentResponse {
	userId := ctx.Param("user_id")
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	stream, err := msc.GetMentionsByUserId(context.Background(), &mention_proto.MentionRequestByUserId{
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
