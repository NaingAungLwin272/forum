package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bookmark_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/features/bookmark"
	comment_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/features/comment"
	mention_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/features/mention"
	question_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/features/question"
	view_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/features/view"
	vote_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/features/vote"
)

type VoteController struct {
	VoteServiceInterface vote_service.VoteServiceInterface
}

type BookmarkController struct {
	BookmarkServiceInterface bookmark_service.BookmarkServiceInterface
}

type MentionController struct {
	MentionServiceInterface mention_service.MentionServiceInterface
}

type ViewController struct {
	ViewServiceInterface view_service.ViewServiceInterfae
}

type QuestionController struct {
	QuestionServiceInterface question_service.QuestionServiceInterface
}

type CommentController struct {
	CommentServiceInterface comment_service.CommentServiceInterface
}

// Vote Processes
func (controller *VoteController) CreateVote(ctx *gin.Context) {
	data, err := controller.VoteServiceInterface.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *VoteController) GetVotesByUserId(ctx *gin.Context) {
	data := controller.VoteServiceInterface.GetVotesByUserId(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (controller *VoteController) DeleteVote(ctx *gin.Context) {
	data, err := controller.VoteServiceInterface.DeleteVote(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, data)
	}
}

func (controller *VoteController) GetVotesByUserIdQuestionId(ctx *gin.Context) {
	data, err := controller.VoteServiceInterface.GetVotesByUserIdQuestionId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

// Bookmark Processes
func (controller *BookmarkController) CreateBookmark(ctx *gin.Context) {
	data, err := controller.BookmarkServiceInterface.CreateBookmark(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *BookmarkController) GetBookmarkByUserId(ctx *gin.Context) {
	data := controller.BookmarkServiceInterface.GetBookmarksByUserId(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, &data)
}

func (controller *BookmarkController) DeleteBookmark(ctx *gin.Context) {
	data, err := controller.BookmarkServiceInterface.DeleteBookmark(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, data)
	}

}

func (controller *BookmarkController) GetBookmarksByUserIdQuestionId(ctx *gin.Context) {
	data, err := controller.BookmarkServiceInterface.GetBookmarksByUserIdQuestionId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

// Mention Processes
func (controller *MentionController) CreateMention(ctx *gin.Context) {
	data, err := controller.MentionServiceInterface.CreateMention(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *MentionController) GetMentionsByUserId(ctx *gin.Context) {
	data := controller.MentionServiceInterface.GetMentionsByUserId(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// View Processess
func (controller *ViewController) CreateView(ctx *gin.Context) {
	data, err := controller.ViewServiceInterface.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *ViewController) GetViewsByUserId(ctx *gin.Context) {
	data, err := controller.ViewServiceInterface.GetViewsByUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

// Questions Processes
func (controller *QuestionController) CreateQuestion(ctx *gin.Context) {
	data, err := controller.QuestionServiceInterface.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *QuestionController) GetQuestions(ctx *gin.Context) {
	data := controller.QuestionServiceInterface.GetQuestions(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (controller *QuestionController) GetQuestionById(ctx *gin.Context) {
	data, err := controller.QuestionServiceInterface.GetQuestionById(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *QuestionController) GetQuestionsByUserId(ctx *gin.Context) {
	data := controller.QuestionServiceInterface.GetQuestionsByUserId(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (controller *QuestionController) GetQuestionCountAll(ctx *gin.Context) {
	data, err := controller.QuestionServiceInterface.GetQuestionCountAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &data)
	}
}

func (controller *QuestionController) FilterQuestion(ctx *gin.Context) {
	data := controller.QuestionServiceInterface.FilterQuestion(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (controller *QuestionController) GetFilteredQuestionCount(ctx *gin.Context) {
	data, err := controller.QuestionServiceInterface.GetFilteredQuestionCount(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &data)
	}
}

// Comment Processes
func (controller *CommentController) GetComment(ctx *gin.Context) {
	data, err := controller.CommentServiceInterface.GetComment(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, data)
	}
}

func (controller *CommentController) CreateComment(ctx *gin.Context) {

	data, err := controller.CommentServiceInterface.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *CommentController) UpdateComment(ctx *gin.Context) {
	data, err := controller.CommentServiceInterface.UpdateComment(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, data)
	}
}

func (controller *CommentController) DeleteComment(ctx *gin.Context) {
	data, err := controller.CommentServiceInterface.DeleteComment(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, data)
	}

}

func (controller *CommentController) GetCommentByQuestionId(ctx *gin.Context) {
	data, err := controller.CommentServiceInterface.GetCommentByQuestionId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, data)
	}
}

func (controller *CommentController) GetCommentsByUserId(ctx *gin.Context) {
	data := controller.CommentServiceInterface.GetCommentByUserId(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (controller *CommentController) GetAnswersByUserId(ctx *gin.Context) {
	data := controller.CommentServiceInterface.GetAnswersByUserId(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (controller *CommentController) GetCommentsByUserIdWithSolved(ctx *gin.Context) {
	data := controller.CommentServiceInterface.GetCommentsByUserIdWithSolved(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func NewVoteController(VoteServiceInterface vote_service.VoteServiceInterface) *VoteController {
	return &VoteController{
		VoteServiceInterface: VoteServiceInterface,
	}
}

func NewBookmarkController(BookmarkServiceInterface bookmark_service.BookmarkServiceInterface) *BookmarkController {
	return &BookmarkController{
		BookmarkServiceInterface: BookmarkServiceInterface,
	}
}

func NewMentionController(MentionServiceInterface mention_service.MentionServiceInterface) *MentionController {
	return &MentionController{
		MentionServiceInterface: MentionServiceInterface,
	}
}

func NewViewController(ViewServiceInterface view_service.ViewServiceInterfae) *ViewController {
	return &ViewController{
		ViewServiceInterface: ViewServiceInterface,
	}
}

func NewQuestionController(QuestionServiceInterface question_service.QuestionServiceInterface) *QuestionController {
	return &QuestionController{
		QuestionServiceInterface: QuestionServiceInterface,
	}
}

func NewCommentController(CommentServiceInterface comment_service.CommentServiceInterface) *CommentController {
	return &CommentController{
		CommentServiceInterface: CommentServiceInterface,
	}
}
