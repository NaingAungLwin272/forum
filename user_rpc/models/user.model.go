package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserRequest struct {
	Name          string    `json:"name" bson:"name" binding:"required"`
	StaffID       string    `json:"staff_id" bson:"staff_id" binding:"required"`
	Email         string    `json:"email" bson:"email" binding:"required"`
	Profile       string    `json:"profile" bson:"profile" binding:"required"`
	DisplayName   string    `json:"displayname" bson:"displayname" binding:"required"`
	Role          string    `json:"role" bson:"role" binding:"required"`
	Password      string    `json:"password" bson:"password" binding:"required"`
	TeamId        string    `json:"team_id" bson:"team_id" binding:"required"`
	Address       string    `json:"address" bson:"address"`
	Phone         string    `json:"phone" bson:"phone"`
	Dob           time.Time `json:"dob" bson:"dob"`
	DepartmentId  string    `json:"department_id" bson:"department_id" binding:"required"`
	Deleted       bool      `json:"is_deleted" bson:"is_deleted"`
	AboutMe       string    `json:"about_me,omitempty" bson:"about_me,omitempty"`
	MailSubscribe bool      `json:"is_mail_subscribed,omitempty" bson:"is_mail_subscribed,omitempty"`
	NotiToken     string    `json:"noti_token" bson:"noti_token"`
	LastLogin     time.Time `json:"last_login,omitempty" bson:"last_login,omitempty"`
	LastPost      time.Time `json:"last_post,omitempty" bson:"last_post,omitempty"`
	CreateAt      time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBUser struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	StaffID       string             `json:"staff_id,omitempty" bson:"staff_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	Profile       string             `json:"profile,omitempty" bson:"profile,omitempty"`
	DisplayName   string             `json:"displayname,omitempty" bson:"displayname,omitempty"`
	Password      string             `json:"password,omitempty" bson:"password,omitempty"`
	Role          string             `json:"role,omitempty" bson:"role,omitempty"`
	DepartmentId  string             `json:"department_id,omitempty" bson:"department_id,omitempty"`
	TeamId        string             `json:"team_id,omitempty" bson:"team_id,omitempty"`
	Deleted       bool               `json:"is_deleted,omitempty" bson:"is_deleted,omitempty"`
	AboutMe       string             `json:"about_me,omitempty" bson:"about_me,omitempty"`
	Phone         string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address       string             `json:"address,omitempty" bson:"address,omitempty"`
	Dob           time.Time          `json:"dob,omitempty" bson:"dob,omitempty"`
	MailSubscribe bool               `json:"is_mail_subscribed" bson:"is_mail_subscribed"`
	NotiToken     string             `json:"noti_token,omitempty" bson:"noti_token,omitempty"`
	LastLogin     time.Time          `json:"last_login,omitempty" bson:"last_login,omitempty"`
	LastPost      time.Time          `json:"last_post,omitempty" bson:"last_post,omitempty"`
	CreateAt      time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateUser struct {
	Name          string    `json:"name,omitempty" bson:"name,omitempty"`
	StaffID       string    `json:"staff_id,omitempty" bson:"staff_id,omitempty"`
	Email         string    `json:"email,omitempty" bson:"email,omitempty"`
	Profile       string    `json:"profile,omitempty" bson:"profile,omitempty"`
	CreateAt      time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	DisplayName   string    `json:"displayname,omitempty" bson:"displayname,omitempty"`
	Password      string    `json:"password,omitempty" bson:"password,omitempty"`
	Role          string    `json:"role,omitempty" bson:"role,omitempty"`
	DepartmentId  string    `json:"department_id,omitempty" bson:"department_id,omitempty"`
	TeamId        string    `json:"team_id,omitempty" bson:"team_id,omitempty"`
	AboutMe       string    `json:"about_me" bson:"about_me"`
	Deleted       bool      `json:"is_deleted,omitempty" bson:"is_deleted,omitempty"`
	Phone         string    `json:"phone" bson:"phone"`
	Address       string    `json:"address" bson:"address"`
	Dob           time.Time `json:"dob" bson:"dob"`
	MailSubscribe bool      `json:"is_mail_subscribed" bson:"is_mail_subscribed"`
	NotiToken     string    `json:"noti_token" bson:"noti_token"`
	LastLogin     time.Time `json:"last_login,omitempty" bson:"last_login,omitempty"`
	LastPost      time.Time `json:"last_post,omitempty" bson:"last_post,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DeleteUser struct {
	Deleted   bool      `json:"is_deleted,omitempty" bson:"is_deleted,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type FilterUserRequest struct {
	Page         int64
	Limit        int64
	DisplayName  string   `json:"displayname,omitempty" bson:"displayname,omitempty"`
	Email        string   `json:"email,omitempty" bson:"email,omitempty"`
	Name         string   `json:"name,omitempty" bson:"name,omitempty"`
	DepartmentId []string `json:"department_id,omitempty" bson:"department_id,omitempty"`
	TeamId       []string `json:"team_id,omitempty" bson:"team_id,omitempty"`
}

type UserSummary struct {
	Questions     int32 `json:"questions,omitempty" bson:"questions,omitempty"`
	Answers       int32 `json:"answers,omitempty" bson:"answers,omitempty"`
	Votes         int32 `json:"votes,omitempty" bson:"votes,omitempty"`
	Solved        int32 `json:"solved,omitempty" bson:"solved,omitempty"`
	Bookmarks     int32 `json:"bookmarks,omitempty" bson:"bookmarks,omitempty"`
	Badges        int32 `json:"badges,omitempty" bson:"badges,omitempty"`
	Notifications int32 `json:"notifications,omitempty" bson:"notifications,omitempty"`
	Messages      int32 `json:"messages,omitempty" bson:"messages,omitempty"`
}
