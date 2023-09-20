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
	UserId        string    `json:"user_id,omitempty" bson:"user_id,omitempty" binding:"required"`
	ReactionLevel int32     `json:"reaction_level" bson:"reaction_level"`
	QaLevel       int32     `json:"qa_level" bson:"qa_level"`
	QuestionCount int32     `json:"question_count" bson:"question_count"`
	AnswerCount   int32     `json:"answer_count" bson:"answer_count"`
	SolvedCount   int32     `json:"solved_count" bson:"solved_count"`
	CreateAt      time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBUser struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	StaffID      string             `json:"staff_id,omitempty" bson:"staff_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Profile      string             `json:"profile,omitempty" bson:"profile,omitempty"`
	DisplayName  string             `json:"displayname,omitempty" bson:"displayname,omitempty"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	Role         string             `json:"role,omitempty" bson:"role,omitempty"`
	DepartmentId string             `json:"department_id,omitempty" bson:"department_id,omitempty"`
	TeamId       string             `json:"team_id,omitempty" bson:"team_id,omitempty"`
	Deleted      bool               `json:"deleted,omitempty" bson:"deleted,omitempty"`
	AboutMe      string             `json:"about_me,omitempty" bson:"about_me,omitempty"`
	Phone        string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address      string             `json:"address,omitempty" bson:"address,omitempty"`
	Dob          time.Time          `json:"dob,omitempty" bson:"dob,omitempty"`
	LastLogin    time.Time          `json:"last_login,omitempty" bson:"last_login,omitempty"`
	LastPost     time.Time          `json:"last_post,omitempty" bson:"last_post,omitempty"`
	CreateAt     time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
