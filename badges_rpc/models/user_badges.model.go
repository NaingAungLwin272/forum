package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBUserBadge struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User_Id   string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Badge_Id  string             `json:"badge_id,omitempty" bson:"badge_id,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CreateUserBadgeRequest struct {
	User_Id   string    `json:"user_id" bson:"user_id" binding:"required"`
	Badge_Id  string    `json:"badge_id" bson:"badge_id"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateUserBadge struct {
	User_Id   string    `json:"user_id" bson:"user_id" binding:"required"`
	Badge_Id  string    `json:"badge_id" bson:"badge_id" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
