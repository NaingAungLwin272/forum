package user_proxy

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/timestamp"
	badge_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/badge/pb"
	features_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	mail_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/mail/pb"
	notis_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/noti/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/models"
	user_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/pb"
	_ "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Auth Processes
func Register(ctx *gin.Context, asc user_proto.AuthServiceClient) (*user_proto.RegisterResponse, error) {
	registerRequestModel := models.CreateUserRequest{}

	if err := ctx.BindJSON(&registerRequestModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := asc.Register(context.Background(), &user_proto.RegisterRequest{
		Name:     registerRequestModel.Name,
		Email:    registerRequestModel.Email,
		StaffId:  registerRequestModel.StaffID,
		Password: registerRequestModel.Password,
	})

	return res, err
}

func Login(ctx *gin.Context, asc user_proto.AuthServiceClient, usc user_proto.UserServiceClient) (*user_proto.LoginResponse, error) {
	loginRequestModel := models.LoginRequest{}

	// bodyAsByteArray, _ := ioutil.ReadAll(ctx.Request.Body)
	// jsonBody := string(bodyAsByteArray)
	// log.Println(jsonBody,  "data is here")
	if err := ctx.BindJSON(&loginRequestModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := asc.Login(context.Background(), &user_proto.LoginRequest{
		Email:        loginRequestModel.Email,
		Password:     loginRequestModel.Password,
		IsRememberMe: loginRequestModel.IsRememberMe,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
		ctx.AbortWithError(http.StatusNotFound, err)
		return nil, err
	}

	user, _ := usc.GetUser(context.Background(), &user_proto.UserRequest{
		XId: res.User.XId,
	})

	var role int32
	switch user.Role {
	case "manager":
		role = 1
	case "bse":
		role = 2
	case "leader":
		role = 3
	case "sub leader":
		role = 4
	case "senior":
		role = 5
	case "junior":
		role = 6
	default:
		ctx.AbortWithError(http.StatusBadRequest, status.Errorf(codes.InvalidArgument, "Invalid role value"))
		return nil, status.Errorf(codes.InvalidArgument, "Invalid role value")
	}

	currentLoginTime := time.Now()
	dobTime, err := convertProtoTimestampToTime(user.Dob)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	_, er := usc.UpdateUser(context.Background(), &user_proto.UpdateUserRequest{
		XId:           res.User.XId,
		Name:          &user.Name,
		Email:         &user.Email,
		Profile:       &user.Profile,
		Phone:         &user.Phone,
		DisplayName:   &user.DisplayName,
		Role:          &role,
		DepartmentId:  &user.DepartmentId,
		TeamId:        &user.TeamId,
		AboutMe:       &user.AboutMe,
		Address:       &user.Address,
		MailSubscribe: &user.MailSubscribe,
		Dob:           &timestamp.Timestamp{Seconds: dobTime.Unix(), Nanos: int32(dobTime.Nanosecond())},
		LastPost:      user.LastPost,
		LastLogin:     &timestamp.Timestamp{Seconds: currentLoginTime.Unix(), Nanos: int32(currentLoginTime.Nanosecond())},
	})

	if er != nil {
		fmt.Println("Error updating user last login time:", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	return res, err
}

func convertProtoTimestampToTime(ts *timestamppb.Timestamp) (time.Time, error) {
	return time.Unix(ts.Seconds, int64(ts.Nanos)).UTC(), nil
}

func ResetPassword(ctx *gin.Context, asc user_proto.AuthServiceClient) (*user_proto.ResetPasswordResponse, error) {
	resetPasswordRequestModel := models.ResetPasswordRequest{}
	if err := ctx.BindJSON(&resetPasswordRequestModel); err != nil {
		log.Println(err, "Error is here in proxy")
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := asc.ResetPassword(context.Background(), &user_proto.ResetPasswordRequest{
		Email:    resetPasswordRequestModel.Email,
		Token:    resetPasswordRequestModel.Token,
		Password: resetPasswordRequestModel.Password,
	})

	return res, err
}

func ChangePassword(ctx *gin.Context, asc user_proto.AuthServiceClient) (*user_proto.ChangePasswordResponse, error) {
	changePasswordRequestModel := models.ChangePasswordRequest{}
	if err := ctx.BindJSON(&changePasswordRequestModel); err != nil {
		log.Println(err, "Error is here in proxy")
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := asc.ChangePassword(context.Background(), &user_proto.ChangePasswordRequest{
		UserId:          changePasswordRequestModel.UserId,
		Password:        changePasswordRequestModel.Password,
		NewPassword:     changePasswordRequestModel.NewPassword,
		ConfrimPassword: changePasswordRequestModel.ConfirmPassword,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	return res, err
}

func ForgetPassword(ctx *gin.Context, asc user_proto.AuthServiceClient, msc mail_proto.MailServiceClient) (*mail_proto.MailResponse, error) {

	forgetPasswordRequestModel := models.ForgetPasswordRequest{}

	if err := ctx.BindJSON(&forgetPasswordRequestModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	token, err := asc.ForgetPassword(context.Background(), &user_proto.ForgetPasswordRequest{
		Email:  forgetPasswordRequestModel.Email,
		Origin: &forgetPasswordRequestModel.Origin,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return nil, err
	}

	req := &mail_proto.ForgetMailRequest{
		Email:  forgetPasswordRequestModel.Email,
		Token:  token.Token,
		Origin: &token.Origin,
		Name:   token.Name,
	}

	response, err := msc.ForgetPasswordMail(context.Background(), req)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}

	return response, nil
}

// User Processes
func CreateUser(ctx *gin.Context, usc user_proto.UserServiceClient, bsc badge_proto.UserPointServiceClient, masc mail_proto.MailServiceClient) (*user_proto.UserResponse, error) {
	userModel := &models.CreateUserRequest{}

	if err := ctx.BindJSON(&userModel); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}
	var dobTimestamp *timestamp.Timestamp

	if userModel.Dob != nil {
		dobTimestamp = timestamppb.New(*userModel.Dob)
	}

	res, err := usc.CreateUser(context.Background(), &user_proto.CreateUserRequest{
		Name:         userModel.Name,
		StaffId:      userModel.StaffID,
		Email:        userModel.Email,
		Profile:      userModel.Profile,
		DisplayName:  userModel.DisplayName,
		Password:     "123456",
		TeamId:       userModel.TeamId,
		DepartmentId: userModel.DepartmentId,
		Role:         int32(userModel.Role),
		AboutMe:      userModel.AboutMe,
		Phone:        userModel.Phone,
		NotiToken:    userModel.NotiToken,
		Address:      userModel.Address,
		Dob:          dobTimestamp,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return nil, err
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	if res != nil {
		bsc.CreateUserPoint(context.Background(), &badge_proto.CreateUserPointRequest{
			UserId:        res.XId,
			ReactionLevel: 0,
			QaLevel:       0,
			QuestionCount: 0,
			AnswerCount:   0,
			SolvedCount:   0,
		})
	}
	if res.Email != "" {
		displayName := res.DisplayName
		a, err := masc.SendMail(context.Background(), &mail_proto.MailRequest{
			Email:   res.Email,
			Type:    7,
			Subject: fmt.Sprintf("Welcome to Our MTM Community Forum! - %s", displayName),
		})
		fmt.Print(a, "a", err, "err")
	}

	return res, err
}

func GetUserList(ctx *gin.Context, usc user_proto.UserServiceClient) []*user_proto.User {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	stream, err := usc.GetUsers(context.Background(), &user_proto.GetUsersRequest{
		Page:  &convertedPage,
		Limit: &convertedLimit,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var users []*user_proto.User
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}

		users = append(users, res)
	}

	return users
}

func GetUser(usc user_proto.UserServiceClient, userID string) (*user_proto.UserResponse, error) {

	res, err := usc.GetUser(context.Background(), &user_proto.UserRequest{
		XId: userID,
	})

	return res, err
}

func UpdateUser(ctx *gin.Context, usc user_proto.UserServiceClient) (*user_proto.UserResponse, error) {
	userModel := models.UpdateUser{}
	userId := ctx.Param("user_id")

	if err := ctx.BindJSON(&userModel); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var dobTimestamp *timestamp.Timestamp

	if userModel.Dob != nil {
		dobTimestamp = timestamppb.New(*userModel.Dob)
	}
	res, err := usc.UpdateUser(context.Background(), &user_proto.UpdateUserRequest{
		XId:           userId,
		StaffId:       &userModel.StaffID,
		Name:          &userModel.Name,
		Email:         &userModel.Email,
		Profile:       &userModel.Profile,
		Phone:         &userModel.Phone,
		DisplayName:   &userModel.DisplayName,
		Role:          &userModel.Role,
		DepartmentId:  &userModel.DepartmentId,
		TeamId:        &userModel.TeamId,
		AboutMe:       &userModel.AboutMe,
		Address:       &userModel.Address,
		MailSubscribe: &userModel.MailSubscribe,
		NotiToken:     &userModel.NotiToken,
		LastPost:      timestamppb.New(userModel.LastPost),
		LastLogin:     timestamppb.New(userModel.LastLogin),
		Dob:           dobTimestamp,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	return res, err
}

func DeleteUser(ctx *gin.Context, usc user_proto.UserServiceClient) (*user_proto.DeleteUserResponse, error) {
	userId := ctx.Param("user_id")

	res, err := usc.DeleteUser(context.Background(), &user_proto.UserRequest{
		XId: userId,
	})

	if err != nil {
		if strings.Contains(err.Error(), "connection") {
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
	}

	return res, err
}

func CreateUserWithCsv(ctx *gin.Context, usc user_proto.UserServiceClient, bsc badge_proto.UserPointServiceClient, masc mail_proto.MailServiceClient) (*user_proto.UserWithCsvResponse, error) {

	var userRequests []*user_proto.CreateUserRequest

	userModel := []models.CreateUserRequest{}

	if err := ctx.BindJSON(&userModel); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}

	for _, req := range userModel {
		user := &user_proto.CreateUserRequest{
			Name:          req.Name,
			StaffId:       req.StaffID,
			Email:         req.Email,
			Profile:       req.Profile,
			DisplayName:   req.DisplayName,
			Password:      "123456",
			MailSubscribe: true,
			TeamId:        req.TeamId,
			DepartmentId:  req.DepartmentId,
			Role:          int32(req.Role),
			Phone:         req.Phone,
			Address:       req.Address,
			AboutMe:       req.AboutMe,
		}

		userRequests = append(userRequests, user)
	}

	request := &user_proto.CreateUserWithCsvRequest{
		Users: userRequests,
	}

	res, err := usc.CreateUsersWithCsv(context.Background(), request)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	} else {
		for _, userData := range res.Users {
			bsc.CreateUserPoint(context.Background(), &badge_proto.CreateUserPointRequest{
				UserId:        userData.XId,
				ReactionLevel: 0,
				QaLevel:       0,
				QuestionCount: 0,
				AnswerCount:   0,
				SolvedCount:   0,
			})

			if userData.Email != "" {
				displayName := userData.DisplayName
				masc.SendMail(context.Background(), &mail_proto.MailRequest{
					Email:   userData.Email,
					Type:    7,
					Subject: fmt.Sprintf("LogIn Information - %s", displayName),
				})
			}
		}
	}

	return res, err
}

func FilterUser(ctx *gin.Context, usc user_proto.UserServiceClient) []*user_proto.User {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)
	userModel := models.FilterUserRequest{}

	if err := ctx.BindJSON(&userModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	stream, err := usc.FilterUser(context.Background(), &user_proto.FilterUserRequest{
		Page:         &convertedPage,
		Limit:        &convertedLimit,
		Name:         &userModel.Name,
		Email:        &userModel.Email,
		TeamId:       userModel.TeamId,
		DepartmentId: userModel.DepartmentId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var users []*user_proto.User
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}

		users = append(users, res)
	}

	return users
}

func GetUserSummary(ctx *gin.Context, fsc features_proto.BookmarkServiceClient, vsc features_proto.VoteServiceClient, qsc features_proto.QuestionServiceClient, nsc notis_proto.NotiServiceClient, bsc badge_proto.UserBadgeServiceClient, asc features_proto.CommentServiceClient, msc features_proto.MentionServiceClient) *models.UserSummary {
	userId := ctx.Param("user_id")

	bookmarkResult, _ := fsc.GetBookmarkCount(
		context.Background(),
		&features_proto.BookmarkRequestByUserId{
			UserId: userId,
		},
	)

	voteResult, _ := vsc.GetVoteCount(
		context.Background(),
		&features_proto.VoteRequestByUserId{
			UserId: userId,
		},
	)

	questionResult, _ := qsc.GetQuestionCount(
		context.Background(),
		&features_proto.QuestionResquestByUserId{
			UserId: userId,
		},
	)

	notificationResult, _ := nsc.GetNotiForUserSummary(
		context.Background(),
		&notis_proto.NotiRequestByUserId{
			UserId: userId,
		},
	)

	badgeResult, _ := bsc.GetBadgeCount(
		context.Background(),
		&badge_proto.BadgeRequestByUserId{
			UserId: userId,
		},
	)

	commentResult, _ := asc.GetCommentCount(
		context.Background(),
		&features_proto.CommentResquestByUserId{
			UserId: userId,
		},
	)

	solvedResult, _ := asc.GetCommentCountBySolved(
		context.Background(),
		&features_proto.CommentResquestByUserId{
			UserId: userId,
		},
	)

	mentionResult, _ := msc.GetMentionCount(
		context.Background(),
		&features_proto.MentionRequestByUserId{
			UserId: userId,
		},
	)

	user := &models.UserSummary{
		Questions:     int(questionResult.Count),
		Answers:       int(commentResult.Count),
		Votes:         int(voteResult.Count),
		Solved:        int(solvedResult.Count),
		Bookmarks:     int(bookmarkResult.Count),
		Badges:        int(badgeResult.Count),
		Notifications: int(notificationResult.Count),
		Messages:      int(0),
		Mentions:      int(mentionResult.Count),
	}

	return user
}

func GetUserCount(ctx *gin.Context, usc user_proto.UserServiceClient) *user_proto.UserCountResponse {
	userModel := models.FilterUserRequest{}

	if err := ctx.BindJSON(&userModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	res, _ := usc.GetUserCount(context.Background(), &user_proto.FilterUserRequest{
		Name:         &userModel.Name,
		Email:        &userModel.Email,
		TeamId:       userModel.TeamId,
		DepartmentId: userModel.DepartmentId,
	})

	return res
}

// Team Processes
func CreateTeam(ctx *gin.Context, usc user_proto.TeamServiceClient) (*user_proto.TeamResponse, error) {
	teamModel := models.CreateTeamRequest{}

	if err := ctx.BindJSON(&teamModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := usc.CreateTeam(context.Background(), &user_proto.CreateTeamRequest{
		Name:         teamModel.Name,
		DepartmentId: teamModel.DeparmentId,
	})

	return res, err
}

func GetTeamList(ctx *gin.Context, usc user_proto.TeamServiceClient) []*user_proto.Team {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	stream, err := usc.GetTeams(context.Background(), &user_proto.GetTeamsRequest{
		Page:  &convertedPage,
		Limit: &convertedLimit,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var teams []*user_proto.Team
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}

		teams = append(teams, res)
	}

	return teams
}

func GetTeam(ctx *gin.Context, usc user_proto.TeamServiceClient) (*user_proto.TeamResponse, error) {
	teamId := ctx.Param("team_id")

	res, err := usc.GetTeam(context.Background(), &user_proto.TeamRequest{
		XId: teamId,
	})

	return res, err
}

func GetTeamCount(ctx *gin.Context, usc user_proto.TeamServiceClient) (*user_proto.TeamCountResponse, error) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)
	res, err := usc.GetTeamCount(context.Background(), &user_proto.GetTeamsRequest{
		Page:  &convertedPage,
		Limit: &convertedLimit,
	})
	return res, err
}

func UpdateTeam(ctx *gin.Context, usc user_proto.TeamServiceClient) (*user_proto.TeamResponse, error) {
	userModel := user_proto.TeamUpdateRequest{}
	teamId := ctx.Param("team_id")

	if err := ctx.BindJSON(&userModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	res, err := usc.UpdateTeam(context.Background(), &user_proto.TeamUpdateRequest{
		XId:          teamId,
		Name:         userModel.Name,
		DepartmentId: userModel.DepartmentId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	return res, err
}

func DeleteTeam(ctx *gin.Context, usc user_proto.TeamServiceClient) (*user_proto.DeleteTeamResponse, error) {
	teamId := ctx.Param("team_id")

	res, err := usc.DeleteTeam(context.Background(), &user_proto.TeamRequest{
		XId: teamId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	return res, err
}

func GetTeamsByDepartmentId(ctx *gin.Context, usc user_proto.TeamServiceClient) []*user_proto.Team {
	departmentId := ctx.Param("department_id")

	stream, err := usc.GetTeamByDeparmentId(context.Background(), &user_proto.TeamRequestByDeparmentId{
		DepartmentId: departmentId,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil
	}

	var teams []*user_proto.Team
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return nil
		}

		teams = append(teams, res)
	}

	return teams

}

// Department Processes
func CreateDeparment(ctx *gin.Context, usc user_proto.DeparmentServiceClient) (*user_proto.DeparmentResponse, error) {
	departmentModel := models.CreateDeparmentRequest{}

	if err := ctx.BindJSON(&departmentModel); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	res, err := usc.CreateDeparment(context.Background(), &user_proto.CreateDeparmentRequest{
		Name: departmentModel.Name,
	})

	return res, err
}

func GetDepartmentsList(ctx *gin.Context, usc user_proto.DeparmentServiceClient) []*user_proto.Deparment {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)

	stream, err := usc.GetDeparments(context.Background(), &user_proto.GetDeparmentsRequest{
		Page:  &convertedPage,
		Limit: &convertedLimit,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	var departments []*user_proto.Deparment
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}

		departments = append(departments, res)
	}

	return departments
}

func GetDepartment(ctx *gin.Context, usc user_proto.DeparmentServiceClient) (*user_proto.DeparmentResponse, error) {
	departmentId := ctx.Param("department_id")

	res, err := usc.GetDeparment(context.Background(), &user_proto.DeparmentRequest{
		XId: departmentId,
	})

	return res, err
}

func UpdateDepartment(ctx *gin.Context, usc user_proto.DeparmentServiceClient) (*user_proto.DeparmentResponse, error) {
	departmentModel := user_proto.DeparmentUpdateRequest{}
	departmentId := ctx.Param("department_id")

	if err := ctx.BindJSON(&departmentModel); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	res, err := usc.UpdateDeparment(context.Background(), &user_proto.DeparmentUpdateRequest{
		XId:  departmentId,
		Name: departmentModel.Name,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	return res, err
}

func DeleteDepartment(ctx *gin.Context, usc user_proto.DeparmentServiceClient) (*user_proto.DeleteDeparmentResponse, error) {
	teamId := ctx.Param("department_id")

	res, err := usc.DeleteDeparment(context.Background(), &user_proto.DeparmentRequest{
		XId: teamId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	return res, err
}

func GetDepartmentCount(ctx *gin.Context, usc user_proto.DeparmentServiceClient) (*user_proto.DepartmentCountResponse, error) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	convertedPage, _ := strconv.ParseInt(page, 0, 64)
	convertedLimit, _ := strconv.ParseInt(limit, 0, 64)
	res, err := usc.GetDepartmentCount(context.Background(), &user_proto.GetDeparmentsRequest{
		Page:  &convertedPage,
		Limit: &convertedLimit,
	})
	return res, err
}

func GetUserByDisplayName(ctx *gin.Context, usc user_proto.UserServiceClient) (*user_proto.UserNameResponse, error) {
	displayName := ctx.Param("displayname")
	res, err := usc.GetUserByDisplayName(context.Background(), &user_proto.UserNameRequest{
		Displayname: displayName,
	})
	return res, err
}
