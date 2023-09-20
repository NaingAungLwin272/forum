package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateViewRequest struct {
	User_Id     string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Question_Id string    `json:"question_id,omitempty" bson:"question_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBView struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User_Id     string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Question_Id string             `json:"question_id,omitempty" bson:"question_id,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateView struct {
	User_Id     string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Question_Id string    `json:"question_id,omitempty" bson:"question_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
