package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateBadgeRequest struct {
	Name        string    `json:"name" bson:"name" binding:"required"`
	Description string    `json:"description" bson:"description" binding:"required"`
	Type        int32     `json:"type,omitempty" bson:"type,omitempty"`
	Level       int32     `json:"level" bson:"level" binding:"required"`
	CreateAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBBadge struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Type        int32              `json:"type,omitempty" bson:"type,omitempty"`
	Level       int32              `json:"level,omitempty" bson:"level,omitempty"`
	CreateAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateBadge struct {
	Name        string    `json:"name,omitempty" bson:"name,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Type        int32     `json:"type,omitempty" bson:"type,omitempty"`
	Level       int32     `json:"level,omitempty" bson:"level,omitempty"`
	CreateAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
