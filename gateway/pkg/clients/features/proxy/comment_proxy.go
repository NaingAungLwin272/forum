package features_proxy

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	badge_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/badge/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/models"
	comment_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	vote_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	mail_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/mail/pb"
	noti_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/noti/pb"
	user_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/config"
)

func CreateComment(ctx *gin.Context, csc comment_proto.CommentServiceClient, qsc comment_proto.QuestionServiceClient, nsc noti_proto.NotiServiceClient, masc mail_proto.MailServiceClient, usc user_proto.UserServiceClient, bsc badge_proto.UserPointServiceClient) (*comment_proto.CommentResponse, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at configs", err)
	}
	baseUrl := config.FRONT_END_URL
	b := models.CreateCommentRequest{}
	if err := ctx.BindJSON(&b); err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("body can't be empty")
		} else {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return nil, err
		}
	}
	res, err := csc.CreateComment(context.Background(), &comment_proto.CreateCommentRequest{
		UserId:      b.User_Id,
		QuestionId:  b.Question_Id,
		ParentId:    b.Parent_Id,
		Sort:        b.Sort,
		Description: b.Description,
		VoteCount:   b.Vote_Count,
		IsSolution:  b.Is_Solution,
		IsDeleted:   b.Is_Deleted,
	})

	if res != nil {
		userPoint, err := bsc.GetUserPoint(context.Background(), &badge_proto.UserPointRequest{
			Id: res.UserId,
		})

		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return nil, err
		}

		if userPoint != nil {
			bsc.UpdateUserPoint(context.Background(), &badge_proto.UpdateUserPointRequest{
				UserId:        userPoint.UserPoint.UserId,
				ReactionLevel: userPoint.UserPoint.ReactionLevel,
				QaLevel:       userPoint.UserPoint.QaLevel,
				QuestionCount: userPoint.UserPoint.QuestionCount,
				AnswerCount:   userPoint.UserPoint.AnswerCount + 1,
				SolvedCount:   userPoint.UserPoint.SolvedCount,
			})
		}
	}

	parent := res.ParentId
	commentUserId := res.UserId
	queRes, _ := qsc.GetQuestion(context.Background(), &vote_proto.QuestionRequest{
		XId: res.QuestionId,
	})

	uniqueUserIDs := make(map[string]bool)

	for _, existingUserID := range queRes.UserIds {
		uniqueUserIDs[existingUserID] = true
	}

	if !uniqueUserIDs[commentUserId] {
		queRes.UserIds = append(queRes.UserIds, commentUserId)
		uniqueUserIDs[commentUserId] = true
	}

	qsc.UpdateQuestion(context.Background(), &comment_proto.UpdateQuestionRequest{
		XId:           queRes.XId,
		ViewCount:     &queRes.ViewCount,
		ReplyCount:    &queRes.ReplyCount,
		VoteCount:     &queRes.VoteCount,
		SolutionCount: queRes.SolutionCount,
		UserIds:       queRes.UserIds,
	})

	if parent != "" && commentUserId != "" {
		cres, _ := csc.GetComment(context.Background(), &comment_proto.CommentRequest{
			XId: parent,
		})

		cures, _ := csc.GetComment(context.Background(), &comment_proto.CommentRequest{
			XId: res.XId,
		})
		parentUserId := cres.UserId
		ures, _ := usc.GetUser(context.Background(), &user_proto.UserRequest{
			XId: parentUserId,
		})

		replyUserRes, _ := usc.GetUser(context.Background(), &user_proto.UserRequest{
			XId: cures.UserId,
		})

		link := fmt.Sprintf("%s/qa-detail/%s#%s", baseUrl, res.QuestionId, res.XId)
		msg := fmt.Sprintf("<span class=\"mm\">%s</span> reply you at <span class=\"mm\">%s</span> ", replyUserRes.DisplayName, queRes.Title)
		if res.UserId != queRes.UserId {
			if parentUserId != "" {
				nsc.CreateNoti(context.Background(), &noti_proto.CreateNotiRequest{
					UserId:      parentUserId,
					Type:        2,
					Name:        "Reply",
					Description: msg,
					Link:        &link,
					Status:      true,
				})
				fmt.Println(b.NotiToken, "notiToken..")
				NotiFunc(ctx, b.NotiToken, "Comment", "You Got A Comment") // added real time noti

				if ures.MailSubscribe {
					displayName := ures.DisplayName
					masc.SendMail(context.Background(), &mail_proto.MailRequest{
						Email:   ures.Email,
						Type:    2,
						Link:    &link,
						Subject: fmt.Sprintf("Reply Mail- %s", displayName),
					})
				}
			}
		} else if res.UserId == queRes.UserId {
			nsc.CreateNoti(context.Background(), &noti_proto.CreateNotiRequest{
				UserId:      parentUserId,
				Type:        2,
				Name:        "Reply",
				Description: msg,
				Link:        &link,
				Status:      true,
			})
			if ures.MailSubscribe {
				displayName := ures.DisplayName
				masc.SendMail(context.Background(), &mail_proto.MailRequest{
					Email:   ures.Email,
					Type:    2,
					Link:    &link,
					Subject: fmt.Sprintf("Reply Mail- %s", displayName),
				})
			}
		}
	}
	getCommentResultByQuestionId, _ := qsc.GetQuestion(context.Background(), &comment_proto.QuestionRequest{
		XId: res.QuestionId,
	})

	commentCountByQuestionId := getCommentResultByQuestionId.ReplyCount + 1
	qsc.UpdateQuestion(context.Background(), &comment_proto.UpdateQuestionRequest{
		XId:        getCommentResultByQuestionId.XId,
		VoteCount:  &getCommentResultByQuestionId.VoteCount,
		ViewCount:  &getCommentResultByQuestionId.ViewCount,
		ReplyCount: &commentCountByQuestionId,
	})

	return res, err
}

func UpdateComment(ctx *gin.Context, csc comment_proto.CommentServiceClient, nsc noti_proto.NotiServiceClient, usc user_proto.UserServiceClient, masc mail_proto.MailServiceClient, qsc vote_proto.QuestionServiceClient) *comment_proto.CommentResponse {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at configs", err)
	}
	baseUrl := config.FRONT_END_URL
	b := models.UpdateComment{}
	commentId := ctx.Param("comment_id")
	if err := ctx.BindJSON(&b); err != nil {
		if err == io.EOF {
			return nil
		} else {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return nil
		}
	}

	commentResult, _ := csc.GetComment(context.Background(), &comment_proto.CommentRequest{
		XId: commentId,
	})

	res, err := csc.UpdateComment(context.Background(), &comment_proto.UpdateCommentRequest{
		XId:         commentId,
		Sort:        &b.Sort,
		Description: &b.Description,
		VoteCount:   &commentResult.VoteCount,
		IsSolution:  &b.Is_Solution,
		IsDeleted:   &b.Is_Deleted,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	queRes, _ := qsc.GetQuestion(context.Background(), &vote_proto.QuestionRequest{
		XId: res.QuestionId,
	})

	if res.IsSolution.Value {
		qsc.UpdateQuestion(context.Background(), &vote_proto.UpdateQuestionRequest{
			XId:           queRes.XId,
			ViewCount:     &queRes.ViewCount,
			ReplyCount:    &queRes.ReplyCount,
			VoteCount:     &queRes.VoteCount,
			SolutionCount: queRes.SolutionCount + 1,
		})
		NotiFunc(ctx, b.NotiToken, "Solution", "You Got A Solution") // added real time noti
	} else {
		qsc.UpdateQuestion(context.Background(), &vote_proto.UpdateQuestionRequest{
			XId:           queRes.XId,
			ViewCount:     &queRes.ViewCount,
			ReplyCount:    &queRes.ReplyCount,
			VoteCount:     &queRes.VoteCount,
			SolutionCount: queRes.SolutionCount - 1,
		})
	}

	link := fmt.Sprintf("%s/qa-detail/%s#%s", baseUrl, res.QuestionId, res.XId)
	msg := fmt.Sprintf("Your answer at <span class=\"mm\">%s</span> is marked as solution", queRes.Title)
	if res.UserId != queRes.UserId {
		if res.IsSolution != nil && res.IsSolution.Value {
			nsc.CreateNoti(context.Background(), &noti_proto.CreateNotiRequest{
				UserId:      res.UserId,
				Type:        4,
				Name:        "Solved",
				Description: msg,
				Link:        &link,
				Status:      true,
			})
			ures, _ := usc.GetUser(context.Background(), &user_proto.UserRequest{
				XId: res.UserId,
			})
			if ures.MailSubscribe {
				displayName := ures.DisplayName
				masc.SendMail(context.Background(), &mail_proto.MailRequest{
					Email:   ures.Email,
					Type:    4,
					Link:    &link,
					Subject: fmt.Sprintf("Your answer marked as solution - %s", displayName),
				})
			}
		}
	}
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}
	return res
}

func DeleteComment(ctx *gin.Context, csc comment_proto.CommentServiceClient, qsc comment_proto.QuestionServiceClient) *comment_proto.DeleteCommentResponse {
	commentId := ctx.Param("comment_id")

	getComment, _ := csc.GetComment(context.Background(), &comment_proto.CommentRequest{
		XId: commentId,
	})

	res, err := csc.DeleteComment(context.Background(), &comment_proto.CommentRequest{
		XId: commentId,
	})

	getVoteResultByQuestionId, _ := qsc.GetQuestion(context.Background(), &comment_proto.QuestionRequest{
		XId: getComment.QuestionId,
	})

	replyCountByQuestionId := getVoteResultByQuestionId.ReplyCount - 1
	qsc.UpdateQuestion(context.Background(), &comment_proto.UpdateQuestionRequest{
		XId:        getVoteResultByQuestionId.XId,
		VoteCount:  &getVoteResultByQuestionId.VoteCount,
		ReplyCount: &replyCountByQuestionId,
		ViewCount:  &getVoteResultByQuestionId.ViewCount,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}
	return res
}

func GetComment(csc comment_proto.CommentServiceClient, commentID string) *comment_proto.CommentResponse {
	res, err := csc.GetComment(context.Background(), &comment_proto.CommentRequest{
		XId: commentID,
	})

	if err != nil {
		return nil
	}
	return res
}

func GetCommentByQuestionId(ctx *gin.Context, csc comment_proto.CommentServiceClient, questionID string) []*comment_proto.Comment {
	b := comment_proto.QuestionIdRequest{
		QuestionId: questionID,
	}

	stream, err := csc.GetCommentByQuestionId(context.Background(), &b)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}

	var comments []*comment_proto.Comment
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return nil
		}

		comments = append(comments, res)
	}
	return comments
}

func GetCommentsByUserId(ctx *gin.Context, csc comment_proto.CommentServiceClient, usc user_proto.UserServiceClient) []*comment_proto.CommentResponse {
	userId := ctx.Param("user_id")
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	stream, err := csc.GetCommentsByUserId(context.Background(), &comment_proto.CommentResquestByUserId{
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

		commentData := &vote_proto.CommentResponse{
			XId:         res.XId,
			UserId:      res.UserId,
			QuestionId:  res.QuestionId,
			ParentId:    res.ParentId,
			Sort:        res.Sort,
			Description: res.Description,
			VoteCount:   res.VoteCount,
			IsSolution:  res.IsSolution,
			IsDeleted:   res.IsDeleted,
			CreatedAt:   res.CreatedAt,
			UpdatedAt:   res.UpdatedAt,
		}
		commentsData = append(commentsData, commentData)
	}

	return commentsData
}

func GetAnswersByUserId(ctx *gin.Context, csc comment_proto.CommentServiceClient, usc user_proto.UserServiceClient) []*comment_proto.CommentResponse {
	userId := ctx.Param("user_id")
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	stream, err := csc.GetAnswersByUserId(context.Background(), &comment_proto.CommentResquestByUserId{
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

		commentData := &vote_proto.CommentResponse{
			XId:         res.XId,
			UserId:      res.UserId,
			QuestionId:  res.QuestionId,
			ParentId:    res.ParentId,
			Sort:        res.Sort,
			Description: res.Description,
			VoteCount:   res.VoteCount,
			IsSolution:  res.IsSolution,
			IsDeleted:   res.IsDeleted,
			CreatedAt:   res.CreatedAt,
			UpdatedAt:   res.UpdatedAt,
		}

		commentsData = append(commentsData, commentData)
	}

	return commentsData
}

func GetCommentsByUserIdWithSolved(ctx *gin.Context, csc comment_proto.CommentServiceClient, usc user_proto.UserServiceClient) []*comment_proto.CommentResponse {
	userId := ctx.Param("user_id")
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	stream, err := csc.GetCommentsByUserIdWithSolved(context.Background(), &comment_proto.CommentResquestByUserId{
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
			XId: res.XId,
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
