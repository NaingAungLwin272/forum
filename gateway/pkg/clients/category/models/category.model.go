package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateCategoryRequest struct {
	Type      int32     `json:"type,omitempty" bson:"type,omitempty"`
	Name      string    `json:"name" bson:"name" binding:"required"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBCategory struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type      int32              `json:"type,omitempty" bson:"type,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	CreateAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateCategory struct {
	Type      int32     `json:"type,omitempty" bson:"type,omitempty"`
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
