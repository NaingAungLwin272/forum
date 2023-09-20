package models

import (
	"time"

	question_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateQuestionRequest struct {
	UserId      string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title       string    `json:"title,omitempty" bson:"title,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	LanguageIds []string  `json:"language_ids,omitempty" bson:"language_ids,omitempty"`
	TagIds      []string  `json:"tag_ids,omitempty" bson:"tag_ids,omitempty"`
	ViewCount   uint64    `json:"view_count" bson:"view_count"`
	VoteCount   uint64    `json:"vote_count" bson:"vote_count"`
	ReplyCount  uint64    `json:"reply_count" bson:"reply_count"`
	IsDeleted   bool      `json:"is_deleted" bson:"is_deleted"`
	UserIds     []string  `json:"user_ids,omitempty" bson:"user_ids,omitempty"`
	CreateAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	IsMentioned bool      `json:"is_mentioned" bson:"is_mentioned"`
}

type DBQuestion struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId        string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title         string             `json:"title,omitempty" bson:"title,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	LanguageIds   []string           `json:"language_ids,omitempty" bson:"language_ids,omitempty"`
	TagIds        []string           `json:"tag_ids,omitempty" bson:"tag_ids,omitempty"`
	ViewCount     uint64             `json:"view_count" bson:"view_count"`
	VoteCount     uint64             `json:"vote_count" bson:"vote_count"`
	ReplyCount    uint64             `json:"reply_count" bson:"reply_count"`
	SolutionCount uint64             `json:"solution_count" bson:"solution_count"`
	UserIds       []string           `json:"user_ids,omitempty" bson:"user_ids,omitempty"`
	IsDeleted     bool               `json:"is_deleted" bson:"is_deleted"`
	CreateAt      time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateQuestion struct {
	Title         string    `json:"title,omitempty" bson:"title,omitempty"`
	Description   string    `json:"description,omitempty" bson:"description,omitempty"`
	LanguageIds   []string  `json:"language_ids,omitempty" bson:"language_ids,omitempty"`
	TagIds        []string  `json:"tag_ids,omitempty" bson:"tag_ids,omitempty"`
	ViewCount     uint64    `json:"view_count" bson:"view_count"`
	VoteCount     uint64    `json:"vote_count" bson:"vote_count"`
	ReplyCount    uint64    `json:"reply_count" bson:"reply_count"`
	SolutionCount uint64    `json:"solution_count" bson:"solution_count"`
	IsDeleted     bool      `json:"is_deleted" bson:"is_deleted"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type QuestionRequestByPage struct {
	Page  int64 `json:"page" bson:"page"`
	Limit int64 `json:"limit" bson:"limit"`
}

type FilterQuestionRequest struct {
	Page        int64
	Limit       int64
	LanguageIds []string `json:"language_ids,omitempty" bson:"language_id,omitempty"`
	TagIds      []string `json:"tag_ids,omitempty" bson:"email,omitempty"`
	UserId      []string `json:"user_id,omitempty" bson:"name,omitempty"`
	Title       string   `json:"title,omitempty" bson:"department_id,omitempty"`
}

type QuestionDetail struct {
	Id            string                    `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId        string                    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title         string                    `json:"title,omitempty" bson:"title,omitempty"`
	LanguageIds   []string                  `json:"language_ids,omitempty" bson:"language_ids,omitempty"`
	TagIds        []string                  `json:"tag_ids,omitempty" bson:"tag_ids,omitempty"`
	ViewCount     uint64                    `json:"view_count" bson:"view_count"`
	VoteCount     uint64                    `json:"vote_count" bson:"vote_count"`
	ReplyCount    uint64                    `json:"reply_count" bson:"reply_count"`
	SolutionCount uint64                    `json:"solution_count" bson:"solution_count"`
	UserIds       []string                  `json:"user_ids,omitempty" bson:"user_ids, omitempty"`
	CreateAt      time.Time                 `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time                 `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Comments      []*question_proto.Comment `json:"comments,omitempty" bson:"comments,omitempty"`
}

type UpdateUserPoint struct {
	UserId        string    `json:"user_id,omitempty" bson:"user_id,omitempty" binding:"required"`
	ReactionLevel int32     `json:"reaction_level" bson:"reaction_level"`
	QaLevel       int32     `json:"qa_level" bson:"qa_level"`
	QuestionCount int32     `json:"question_count" bson:"question_count"`
	AnswerCount   int32     `json:"answer_count" bson:"answer_count"`
	SolvedCount   int32     `json:"solved_count" bson:"solved_count"`
	CreateAt      time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
