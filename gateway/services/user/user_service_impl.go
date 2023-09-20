package user_service

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	clients "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	mail_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/mail/pb"
	user_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/models"
	user_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Clients clients.ServiceClient
}

// ChangePasswrd implements UserServiceInterface.
func (userSvc *UserService) ChangePasswrd(ctx *gin.Context) (*user_proto.ChangePasswordResponse, error) {
	data, err := user_proxy.ChangePassword(ctx, userSvc.Clients.Auth)
	return data, err
}

// ResetPassword implements UserServiceInterface.
func (userSvc *UserService) ResetPassword(ctx *gin.Context) (*user_proto.ResetPasswordResponse, error) {
	data, err := user_proxy.ResetPassword(ctx, userSvc.Clients.Auth)
	log.Println(err, "Error is here svc interface")
	return data, err
}

// ForgetPassword implements UserServiceInterface.
func (userSvc *UserService) ForgetPassword(ctx *gin.Context) (*mail_proto.MailResponse, error) {
	data, err := user_proxy.ForgetPassword(ctx, userSvc.Clients.Auth, userSvc.Clients.Mail)
	return data, err
}

// Login implements UserServiceInterface.
func (userSvc *UserService) Login(ctx *gin.Context) (*user_proto.LoginResponse, error) {
	data, err := user_proxy.Login(ctx, userSvc.Clients.Auth, userSvc.Clients.User)
	return data, err
}

// Register implements UserServiceInterface.
func (*UserService) Register(ctx *gin.Context) (*user_proto.RegisterResponse, error) {
	panic("unimplemented")
}

// User Processes
func (userSvc *UserService) Create(ctx *gin.Context) (*user_proto.UserResponse, error) {
	data, err := user_proxy.CreateUser(ctx, userSvc.Clients.User, userSvc.Clients.UserPoint, userSvc.Clients.Mail)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, err
	}
	return data, err
}

func (userSvc *UserService) GetUser(ctx *gin.Context) (*user_proto.UserResponse, error) {
	userID := ctx.Param("user_id")
	data, err := user_proxy.GetUser(userSvc.Clients.User, userID)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, err
	}
	return data, err
}

func (userSvc *UserService) GetUsers(ctx *gin.Context) []*user_proto.User {
	data := user_proxy.GetUserList(ctx, userSvc.Clients.User)
	return data
}

func (userSvc *UserService) UpdateUser(ctx *gin.Context) (*user_proto.UserResponse, error) {
	data, err := user_proxy.UpdateUser(ctx, userSvc.Clients.User)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (userSvc *UserService) DeleteUser(ctx *gin.Context) (*user_proto.DeleteUserResponse, error) {
	data, err := user_proxy.DeleteUser(ctx, userSvc.Clients.User)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (userSvc *UserService) CreateUserWithCsv(ctx *gin.Context) (*user_proto.UserWithCsvResponse, error) {
	data, err := user_proxy.CreateUserWithCsv(ctx, userSvc.Clients.User, userSvc.Clients.UserPoint, userSvc.Clients.Mail)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (userSvc *UserService) FilterUser(ctx *gin.Context) []*user_proto.User {
	data := user_proxy.FilterUser(ctx, userSvc.Clients.User)
	return data
}

func (userSvc *UserService) GetUserSummary(ctx *gin.Context) *models.UserSummary {
	data := user_proxy.GetUserSummary(ctx, userSvc.Clients.BookMark, userSvc.Clients.Vote, userSvc.Clients.Question, userSvc.Clients.Noti, userSvc.Clients.UserBadge, userSvc.Clients.Comment, userSvc.Clients.Mention)

	return data
}

func (userSvc *UserService) GetUserCount(ctx *gin.Context) *user_proto.UserCountResponse {
	data := user_proxy.GetUserCount(ctx, userSvc.Clients.User)

	return data
}

// Team Processes
func (userSvc *UserService) CreateTeam(ctx *gin.Context) (*user_proto.TeamResponse, error) {
	data, err := user_proxy.CreateTeam(ctx, userSvc.Clients.Team)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, err
	}
	return data, err
}

func (userSvc *UserService) GetTeam(ctx *gin.Context) (*user_proto.TeamResponse, error) {
	data, err := user_proxy.GetTeam(ctx, userSvc.Clients.Team)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (userSvc *UserService) GetTeams(ctx *gin.Context) []*user_proto.Team {
	data := user_proxy.GetTeamList(ctx, userSvc.Clients.Team)
	return data
}

func (userSvc *UserService) UpdateTeam(ctx *gin.Context) (*user_proto.TeamResponse, error) {
	data, err := user_proxy.UpdateTeam(ctx, userSvc.Clients.Team)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (userSvc *UserService) DeleteTeam(ctx *gin.Context) (*user_proto.DeleteTeamResponse, error) {
	data, err := user_proxy.DeleteTeam(ctx, userSvc.Clients.Team)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (userSvc *UserService) GetTeamsByDepartmentId(ctx *gin.Context) ([]*user_proto.Team, error) {
	data := user_proxy.GetTeamsByDepartmentId(ctx, userSvc.Clients.Team)
	if data == nil {
		return nil, errors.New("no team with that departmentid doesn`t exists")
	}
	return data, nil
}

func (userSvc *UserService) GetTeamCount(ctx *gin.Context) (*user_proto.TeamCountResponse, error) {
	data, err := user_proxy.GetTeamCount(ctx, userSvc.Clients.Team)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

// Department Processes
func (userSvc *UserService) CreateDepartment(ctx *gin.Context) (*user_proto.DeparmentResponse, error) {
	data, err := user_proxy.CreateDeparment(ctx, userSvc.Clients.Department)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, err
	}
	return data, err
}

func (userSvc *UserService) GetDepartment(ctx *gin.Context) (*user_proto.DeparmentResponse, error) {
	data, err := user_proxy.GetDepartment(ctx, userSvc.Clients.Department)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (userSvc *UserService) GetDepartments(ctx *gin.Context) []*user_proto.Deparment {
	data := user_proxy.GetDepartmentsList(ctx, userSvc.Clients.Department)
	return data
}

func (userSvc *UserService) UpdateDepartment(ctx *gin.Context) (*user_proto.DeparmentResponse, error) {
	data, err := user_proxy.UpdateDepartment(ctx, userSvc.Clients.Department)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (userSvc *UserService) DeleteDepartment(ctx *gin.Context) (*user_proto.DeleteDeparmentResponse, error) {
	data, err := user_proxy.DeleteDepartment(ctx, userSvc.Clients.Department)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func (userSvc *UserService) GetDepartmentCount(ctx *gin.Context) (*user_proto.DepartmentCountResponse, error) {
	data, err := user_proxy.GetDepartmentCount(ctx, userSvc.Clients.Department)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

// GetUserByDisplayName implements UserServiceInterface.
func (userSvc *UserService) GetUserByDisplayName(ctx *gin.Context) (*user_proto.UserNameResponse, error) {
	data, err := user_proxy.GetUserByDisplayName(ctx, userSvc.Clients.User)
	if err, ok := err.(mongo.WriteException); ok && err.WriteErrors[0].Code == 11000 {
		return nil, errors.New("unexpected error")
	}
	return data, err
}

func NewUserService(Clients clients.ServiceClient) UserServiceInterface {
	return &UserService{
		Clients: Clients,
	}
}
