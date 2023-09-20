package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTeamRequest struct {
	Name        string    `json:"name" bson:"name" binding:"required"`
	DeparmentId string    `json:"department_id" bson:"department_id" binding:"required"`
	CreateAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBTeam struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	DeparmentId string             `json:"department_id,omitempty" bson:"department_id,omitempty"`
	CreateAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateTeam struct {
	Name        string    `json:"name,omitempty" bson:"name,omitempty"`
	DeparmentId string    `json:"department_id,omitempty" bson:"department_id,omitempty"`
	CreateAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type TeamCount struct {
	Count int64
}
