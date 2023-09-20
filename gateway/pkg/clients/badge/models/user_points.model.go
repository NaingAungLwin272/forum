package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserPointRequest struct {
	UserId        string    `json:"user_id" bson:"user_id" binding:"required"`
	ReactionLevel int32     `json:"reaction_level" bson:"reaction_level"`
	QaLevel       int32     `json:"qa_level,omitempty" bson:"qa_level"`
	QuestionCount int32     `json:"question_count" bson:"question_count"`
	AnswerCount   int32     `json:"answer_count" bson:"answer_count"`
	SolvedCount   int32     `json:"solved_count" bson:"solved_count"`
	CreateAt      time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBUserPoint struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId        string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	ReactionLevel int32              `json:"reaction_level" bson:"reaction_level"`
	QaLevel       int32              `json:"qa_level" bson:"qa_level"`
	QuestionCount int32              `json:"question_count" bson:"question_count"`
	AnswerCount   int32              `json:"answer_count" bson:"answer_count"`
	SolvedCount   int32              `json:"solved_count" bson:"solved_count"`
	CreateAt      time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateUserPoint struct {
	UserId        string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	ReactionLevel int32     `json:"reaction_level" bson:"reaction_level"`
	QaLevel       int32     `json:"qa_level" bson:"qa_level"`
	QuestionCount int32     `json:"question_count" bson:"question_count"`
	AnswerCount   int32     `json:"answer_count" bson:"answer_count"`
	SolvedCount   int32     `json:"solved_count" bson:"solved_count"`
	CreateAt      time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type PointRequestByPage struct {
	Page  int64
	Limit int64
}
