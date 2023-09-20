package user_service

import (
	"github.com/gin-gonic/gin"
	mail_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/mail/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/models"
	user_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/pb"
)

type UserServiceInterface interface {
	// Auth Process
	Register(ctx *gin.Context) (*user_proto.RegisterResponse, error)
	Login(ctx *gin.Context) (*user_proto.LoginResponse, error)
	ForgetPassword(ctx *gin.Context) (*mail_proto.MailResponse, error)
	ResetPassword(ctx *gin.Context) (*user_proto.ResetPasswordResponse, error)
	ChangePasswrd(ctx *gin.Context) (*user_proto.ChangePasswordResponse, error)
	// User Processes
	Create(ctx *gin.Context) (*user_proto.UserResponse, error)
	GetUser(ctx *gin.Context) (*user_proto.UserResponse, error)
	GetUsers(ctx *gin.Context) []*user_proto.User
	UpdateUser(ctx *gin.Context) (*user_proto.UserResponse, error)
	DeleteUser(ctx *gin.Context) (*user_proto.DeleteUserResponse, error)
	CreateUserWithCsv(ctx *gin.Context) (*user_proto.UserWithCsvResponse, error)
	GetUserSummary(ctx *gin.Context) *models.UserSummary
	FilterUser(ctx *gin.Context) []*user_proto.User
	GetUserCount(ctx *gin.Context) *user_proto.UserCountResponse
	GetUserByDisplayName(ctx *gin.Context) (*user_proto.UserNameResponse, error)

	// Team Processes
	CreateTeam(ctx *gin.Context) (*user_proto.TeamResponse, error)
	GetTeam(ctx *gin.Context) (*user_proto.TeamResponse, error)
	GetTeams(ctx *gin.Context) []*user_proto.Team
	UpdateTeam(ctx *gin.Context) (*user_proto.TeamResponse, error)
	DeleteTeam(ctx *gin.Context) (*user_proto.DeleteTeamResponse, error)
	GetTeamsByDepartmentId(ctx *gin.Context) ([]*user_proto.Team, error)
	GetTeamCount(ctx *gin.Context) (*user_proto.TeamCountResponse, error)

	// Department Processes
	CreateDepartment(ctx *gin.Context) (*user_proto.DeparmentResponse, error)
	GetDepartment(ctx *gin.Context) (*user_proto.DeparmentResponse, error)
	GetDepartments(ctx *gin.Context) []*user_proto.Deparment
	UpdateDepartment(ctx *gin.Context) (*user_proto.DeparmentResponse, error)
	DeleteDepartment(ctx *gin.Context) (*user_proto.DeleteDeparmentResponse, error)
	GetDepartmentCount(ctx *gin.Context) (*user_proto.DepartmentCountResponse, error)
}
