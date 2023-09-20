package models

type LoginRequest struct {
	Email        string `json:"email" bson:"email" binding:"required"`
	Password     string `json:"password" bson:"password" binding:"required"`
	IsRememberMe bool   `json:"is_remember_me,omitempty" bson:"is_remember_me,omitempty"`
}

type ForgetPasswordRequest struct {
	Email  string `json:"email" bson:"email" binding:"required"`
	Origin string `json:"origin,omitempty" bson:"origin,omitempty"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Token    string `json:"token" bson:"token" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type ChangePasswordRequest struct {
	UserId          string `json:"user_id" bson:"user_id" binding:"required"`
	Password        string `json:"password" bson:"password" binding:"required"`
	NewPassword     string `json:"new_password" bson:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" bson:"confirm_password" binding:"required"`
}
