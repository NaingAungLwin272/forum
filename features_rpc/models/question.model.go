package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateQuestionRequest struct {
	UserId        string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title         string    `json:"title,omitempty" bson:"title,omitempty"`
	Description   string    `json:"description,omitempty" bson:"description,omitempty"`
	LanguageIds   []string  `json:"language_ids,omitempty" bson:"language_ids,omitempty"`
	TagIds        []string  `json:"tag_ids,omitempty" bson:"tag_ids,omitempty"`
	ViewCount     uint64    `json:"view_count" bson:"view_count"`
	VoteCount     uint64    `json:"vote_count" bson:"vote_count"`
	ReplyCount    uint64    `json:"reply_count" bson:"reply_count"`
	SolutionCount uint64    `json:"solution_count" bson:"solution_count"`
	UserIds       []string  `json:"user_ids,omitempty" bson:"user_ids,omitempty"`
	IsDeleted     bool      `json:"is_deleted" bson:"is_deleted"`
	CreateAt      time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
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
	UserIds       []string  `json:"user_ids,omitempty" bson:"user_ids,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type FilterQuestionRequest struct {
	Page        int64
	Limit       int64
	LanguageIds []string `json:"language_ids,omitempty" bson:"language_id,omitempty"`
	TagIds      []string `json:"tag_ids,omitempty" bson:"email,omitempty"`
	UserId      []string `json:"user_id,omitempty" bson:"name,omitempty"`
	Title       string   `json:"title,omitempty" bson:"department_id,omitempty"`
	Order       string   `json:"order,omitempty" bson:"order,omitempty"`
	Sort        string   `json:"sort,omitempty" bson:"sort,omitempty"`
}
