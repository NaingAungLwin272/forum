package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBComment struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User_Id     string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Question_Id string             `json:"question_id,omitempty" bson:"question_id,omitempty"`
	Parent_Id   string             `json:"parent_id,omitempty" bson:"parent_id,omitempty"`
	Sort        uint64             `json:"sort,omitempty" bson:"sort,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Vote_Count  uint64             `json:"vote_count,omitempty" bson:"vote_count,omitempty"`
	Is_Solution bool               `json:"is_solution" bson:"is_solution"`
	Is_Deleted  bool               `json:"is_deleted" bson:"is_deleted"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
