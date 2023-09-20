package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBComment struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId      string             `json:"userId,omitempty" bson:"userId,omitempty"`
	QuestionId  string             `json:"questionId,omitempty" bson:"questionId,omitempty"`
	Sort        int32              `json:"sort,omitempty" bson:"sort,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	VoteCount   int32              `json:"voteCount,omitempty" bson:"voteCount,omitempty"`
	ViewCount   int32              `json:"viewCount,omitempty" bson:"viewCount,omitempty"`
	Solution    bool               `json:"solution,omitempty" bson:"solution,omitempty"`
	Deleted     bool               `json:"deleted,omitempty" bson:"deleted,omitempty"`
	CreateAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
