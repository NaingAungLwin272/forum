package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateNotiRequest struct {
	UserId      string    `json:"user_id" bson:"user_id" binding:"required"`
	Type        int       `json:"type" bson:"type" binding:"required"`
	Name        string    `json:"name" bson:"name" binding:"required"`
	Description string    `json:"description" bson:"description" binding:"required"`
	Link        string    `json:"link,omitempty" bson:"link,omitempty"`
	Status      *bool     `json:"status,omitempty" bson:"status,omitempty"`
	CreateAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBNoti struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId      string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Type        int                `json:"type,omitempty" bson:"type,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Link        string             `json:"link,omitempty" bson:"link,omitempty"`
	Status      *bool              `json:"status,omitempty" bson:"status,omitempty"`
	CreateAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateNoti struct {
	Name        string    `json:"name,omitempty" bson:"name,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Link        string    `json:"link,omitempty" bson:"link,omitempty"`
	Status      bool      `json:"status,omitempty" bson:"status,omitempty"`
	CreateAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
