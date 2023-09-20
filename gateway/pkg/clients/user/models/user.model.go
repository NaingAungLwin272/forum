package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	Client int32
	Admin  int32
}

type CreateUserRequest struct {
	Name          string     `json:"name" bson:"name" binding:"required"`
	StaffID       string     `json:"staff_id" bson:"staff_id"`
	Email         string     `json:"email" bson:"email" binding:"required"`
	Profile       string     `json:"profile" bson:"profile"`
	DisplayName   string     `json:"display_name" bson:"display_name" binding:"required"`
	Role          int32      `json:"role" bson:"role" binding:"required"`
	Password      string     `json:"password" bson:"password"`
	TeamId        string     `json:"team_id" bson:"team_id" binding:"required"`
	Address       string     `json:"address" bson:"address"`
	Phone         string     `json:"phone" bson:"phone"`
	Dob           *time.Time `json:"dob" bson:"dob"`
	DepartmentId  string     `json:"department_id" bson:"department_id" binding:"required"`
	Deleted       bool       `json:"deleted" bson:"deleted"`
	AboutMe       string     `json:"about_me,omitempty" bson:"about_me,omitempty"`
	MailSubscribe bool       `json:"mail_subscribe,omitempty" bson:"mail_subscribe,omitempty"`
	NotiToken     string     `json:"noti_token" bson:"noti_token"`
	LastLogin     time.Time  `json:"last_login,omitempty" bson:"last_login,omitempty"`
	LastPost      time.Time  `json:"last_post,omitempty" bson:"last_post,omitempty"`
	CreateAt      time.Time  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
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
	Deleted      bool               `json:"delete,omitempty" bson:"delete,omitempty"`
	AboutMe      string             `json:"about_me,omitempty" bson:"about_me,omitempty"`
	Phone        string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address      string             `json:"address,omitempty" bson:"address,omitempty"`
	Dob          time.Time          `json:"dob,omitempty" bson:"dob,omitempty"`
	LastLogin    time.Time          `json:"last_login,omitempty" bson:"last_login,omitempty"`
	LastPost     time.Time          `json:"last_post,omitempty" bson:"last_post,omitempty"`
	CreateAt     time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateUser struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	StaffID       string             `json:"staff_id,omitempty" bson:"staff_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	Profile       string             `json:"profile,omitempty" bson:"profile,omitempty"`
	DisplayName   string             `json:"display_name,omitempty" bson:"display_name,omitempty"`
	Password      string             `json:"password,omitempty" bson:"password,omitempty"`
	Role          int32              `json:"role" bson:"role" binding:"required"`
	DepartmentId  string             `json:"department_id,omitempty" bson:"department_id,omitempty"`
	TeamId        string             `json:"team_id,omitempty" bson:"team_id,omitempty"`
	Deleted       bool               `json:"deleted,omitempty" bson:"deleted,omitempty"`
	AboutMe       string             `json:"about_me,omitempty" bson:"about_me,omitempty"`
	Phone         string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address       string             `json:"address,omitempty" bson:"address,omitempty"`
	Dob           *time.Time         `json:"dob" bson:"dob"`
	MailSubscribe bool               `json:"mail_subscribe,omitempty" bson:"mail_subscribe,omitempty"`
	NotiToken     string             `json:"noti_token,omitempty" bson:"noti_token,omitempty"`
	LastLogin     time.Time          `json:"last_login,omitempty" bson:"last_login,omitempty"`
	LastPost      time.Time          `json:"last_post,omitempty" bson:"last_post,omitempty"`
	CreateAt      time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DeleteUser struct {
	Deleted   bool      `json:"deleted,omitempty" bson:"deleted,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type FilterUserRequest struct {
	DisplayName  string   `json:"displayname,omitempty" bson:"displayname,omitempty"`
	Name         string   `json:"name,omitempty" bson:"name,omitempty"`
	Email        string   `json:"email,omitempty" bson:"email,omitempty"`
	DepartmentId []string `json:"department_id,omitempty" bson:"department_id,omitempty"`
	TeamId       []string `json:"team_id,omitempty" bson:"team_id,omitempty"`
}

type UserSummary struct {
	Questions     int `json:"questions" bson:"questions"`
	Answers       int `json:"answers" bson:"answers"`
	Votes         int `json:"votes" bson:"votes"`
	Solved        int `json:"solved" bson:"solved"`
	Bookmarks     int `json:"bookmarks" bson:"bookmarks"`
	Badges        int `json:"badges" bson:"badges"`
	Notifications int `json:"notifications" bson:"notifications"`
	Messages      int `json:"messages" bson:"messages"`
	Mentions      int `json:"mentions" bson:"mentions"`
}
