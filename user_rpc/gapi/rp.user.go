package gapi

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/consts"
	models "github.com/scm-dev1dev5/mtm-community-forum/user_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (userServer *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	var role = ""
	switch req.GetRole() {
	case 1:
		role = "manager"
	case 2:
		role = "bse"
	case 3:
		role = "leader"
	case 4:
		role = "sub leader"
	case 5:
		role = "senior"
	case 6:
		role = "junior"
	default:
		return nil, status.Errorf(codes.InvalidArgument, "Invalid role value")
	}
	user := &models.CreateUserRequest{
		StaffID:      req.GetStaffId(),
		Name:         req.GetName(),
		Email:        req.GetEmail(),
		Profile:      req.GetProfile(),
		DisplayName:  req.GetDisplayName(),
		Password:     req.GetPassword(),
		DepartmentId: req.GetDepartmentId(),
		TeamId:       req.GetTeamId(),
		Role:         role,
		AboutMe:      req.GetAboutMe(),
		Phone:        req.GetPhone(),
		NotiToken:    req.GetNotiToken(),
		Address:      req.GetAddress(),
		Dob:          req.GetDob().AsTime().Local(),
	}

	newUser, err := userServer.userService.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "email") {
			return nil, status.Errorf(codes.AlreadyExists, "email already exists")
		}
		if strings.Contains(err.Error(), "staff_id") {
			return nil, status.Errorf(codes.AlreadyExists, "staff id already exists")
		}
		if strings.Contains(err.Error(), "displayname") {
			return nil, status.Errorf(codes.AlreadyExists, "displayname already exists")
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserResponse{
		XId:          newUser.Id.Hex(),
		StaffId:      newUser.StaffID,
		Name:         newUser.Name,
		Email:        newUser.Email,
		Profile:      newUser.Profile,
		DisplayName:  newUser.DisplayName,
		Role:         newUser.Role,
		DepartmentId: newUser.DepartmentId,
		TeamId:       newUser.TeamId,
		AboutMe:      newUser.AboutMe,
		Deleted:      wrapperspb.Bool(newUser.Deleted),
		Phone:        newUser.Phone,
		Address:      newUser.Address,
		Dob:          timestamppb.New(user.Dob),
		LastLogin:    timestamppb.New(user.LastLogin),
		LastPost:     timestamppb.New(user.LastPost),
		CreatedAt:    timestamppb.New(newUser.CreateAt),
		UpdatedAt:    timestamppb.New(newUser.UpdatedAt),
	}
	return res, nil
}

func (userServer *UserServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	userId := req.GetXId()

	user, err := userServer.userService.GetUser(userId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user.Deleted {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	res := &pb.UserResponse{
		XId:           user.Id.Hex(),
		StaffId:       user.StaffID,
		Name:          user.Name,
		Email:         user.Email,
		Profile:       user.Profile,
		DisplayName:   user.DisplayName,
		Role:          user.Role,
		DepartmentId:  user.DepartmentId,
		TeamId:        user.TeamId,
		Deleted:       &wrapperspb.BoolValue{Value: user.Deleted},
		AboutMe:       user.AboutMe,
		Address:       user.Address,
		Phone:         user.Phone,
		Dob:           timestamppb.New(user.Dob),
		MailSubscribe: user.MailSubscribe,
		NotiToken:     user.NotiToken,
		LastLogin:     timestamppb.New(user.LastLogin),
		LastPost:      timestamppb.New(user.LastPost),
		CreatedAt:     timestamppb.New(user.CreateAt),
		UpdatedAt:     timestamppb.New(user.UpdatedAt),
	}
	return res, nil
}

func (userServer *UserServer) GetUserByDisplayName(ctx context.Context, req *pb.UserNameRequest) (*pb.UserNameResponse, error) {
	userName := req.GetDisplayname()
	name, err := userServer.userService.GetUserByDisplayName(userName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	res := &pb.UserNameResponse{
		XId: name.Id.Hex(),
	}
	return res, nil
}

func (userServer *UserServer) GetUsers(req *pb.GetUsersRequest, stream pb.UserService_GetUsersServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()

	users, err := userServer.userService.GetUsers(int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, user := range users {
		stream.Send(&pb.User{
			XId:          user.Id.Hex(),
			StaffId:      user.StaffID,
			Name:         user.Name,
			Email:        user.Email,
			Profile:      user.Profile,
			DisplayName:  user.DisplayName,
			Role:         user.Role,
			DepartmentId: user.DepartmentId,
			TeamId:       user.TeamId,
			Deleted:      wrapperspb.Bool(user.Deleted),
			AboutMe:      user.AboutMe,
			Phone:        user.Phone,
			Address:      user.Address,
			Dob:          timestamppb.New(user.Dob),
			NotiToken:    user.NotiToken,
			LastLogin:    timestamppb.New(user.LastLogin),
			LastPost:     timestamppb.New(user.LastPost),
			CreatedAt:    timestamppb.New(user.CreateAt),
			UpdatedAt:    timestamppb.New(user.UpdatedAt),
		})

	}

	return nil
}

func (userServer *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	fmt.Println(req.GetNotiToken(), "notiToken.......")
	var role = ""
	switch req.GetRole() {
	case 1:
		role = "manager"
	case 2:
		role = "bse"
	case 3:
		role = "leader"
	case 4:
		role = "sub leader"
	case 5:
		role = "senior"
	case 6:
		role = "junior"
	default:
		return nil, status.Errorf(codes.InvalidArgument, "Invalid role value")
	}
	userId := req.GetXId()
	user := &models.UpdateUser{
		Name:          req.GetName(),
		StaffID:       req.GetStaffId(),
		Email:         req.GetEmail(),
		Profile:       req.GetProfile(),
		DisplayName:   req.GetDisplayName(),
		Password:      req.GetPassword(),
		TeamId:        req.GetTeamId(),
		DepartmentId:  req.GetDepartmentId(),
		Role:          role,
		Phone:         req.GetPhone(),
		Address:       req.GetAddress(),
		Dob:           req.GetDob().AsTime(),
		LastPost:      req.LastPost.AsTime(),
		LastLogin:     req.LastLogin.AsTime(),
		MailSubscribe: req.GetMailSubscribe(),
		NotiToken:     req.GetNotiToken(),
		AboutMe:       req.GetAboutMe(),
		UpdatedAt:     time.Now(),
	}

	updatedUser, err := userServer.userService.UpdateUser(userId, user, ctx)

	if err != nil {
		if strings.Contains(err.Error(), "email") {
			return nil, status.Errorf(codes.AlreadyExists, "email already exists")
		}
		if strings.Contains(err.Error(), "staff_id") {
			return nil, status.Errorf(codes.AlreadyExists, "staff id already exists")
		}
		if strings.Contains(err.Error(), "displayname") {
			return nil, status.Errorf(codes.AlreadyExists, "displayname already exists")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if updatedUser.Deleted {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	res := &pb.UserResponse{
		XId:          updatedUser.Id.Hex(),
		StaffId:      updatedUser.StaffID,
		Name:         updatedUser.Name,
		Email:        updatedUser.Email,
		Profile:      updatedUser.Profile,
		DisplayName:  updatedUser.DisplayName,
		DepartmentId: updatedUser.DepartmentId,
		TeamId:       updatedUser.TeamId,
		Role:         updatedUser.Role,
		Phone:        updatedUser.Phone,
		Address:      updatedUser.Address,
		Dob:          timestamppb.New(user.Dob),
		AboutMe:      updatedUser.AboutMe,
		NotiToken:    updatedUser.NotiToken,
		Deleted:      wrapperspb.Bool(updatedUser.Deleted),
		CreatedAt:    timestamppb.New(updatedUser.CreateAt),
		UpdatedAt:    timestamppb.New(updatedUser.UpdatedAt),
	}
	return res, nil
}

func (userServer *UserServer) DeleteUser(ctx context.Context, req *pb.UserRequest) (*pb.DeleteUserResponse, error) {
	userId := req.GetXId()

	user, err := userServer.userService.GetUser(userId)
	if user.Deleted {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}
	delete_err := userServer.userService.DeleteUser(userId)
	if delete_err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteUserResponse{
		Success: true,
	}

	return res, nil
}

func (userServer *UserServer) FilterUser(req *pb.FilterUserRequest, stream pb.UserService_FilterUserServer) error {
	user := &models.FilterUserRequest{
		Page:         req.GetPage(),
		Limit:        req.GetLimit(),
		DisplayName:  req.GetDisplayName(),
		Email:        req.GetEmail(),
		DepartmentId: req.GetDepartmentId(),
		TeamId:       req.GetTeamId(),
		Name:         req.GetName(),
	}

	filteredUser, err := userServer.userService.FilterUser(user)
	if err != nil {
		return nil
	}
	for _, filterUser := range filteredUser {
		// if filterUser.Deleted {
		// 	return nil
		// }
		stream.Send(&pb.User{
			XId:          filterUser.Id.Hex(),
			Name:         filterUser.Name,
			StaffId:      filterUser.StaffID,
			Email:        filterUser.Email,
			Profile:      filterUser.Profile,
			DisplayName:  filterUser.DisplayName,
			TeamId:       filterUser.TeamId,
			DepartmentId: filterUser.DepartmentId,
			Role:         filterUser.Role,
			Phone:        filterUser.Phone,
			Address:      filterUser.Address,
			Deleted:      wrapperspb.Bool(filterUser.Deleted),
			LastLogin:    timestamppb.New(filterUser.LastLogin),
			LastPost:     timestamppb.New(filterUser.LastPost),
			CreatedAt:    timestamppb.New(filterUser.CreateAt),
			UpdatedAt:    timestamppb.New(filterUser.UpdatedAt),
		})
	}
	return nil

}

func (userServer UserServer) CreateUsersWithCsv(ctx context.Context, req *pb.CreateUserWithCsvRequest) (*pb.UserWithCsvResponse, error) {
	nilSlice := []*models.CreateUserRequest{}
	users := req.GetUsers()
	existingEmails := make(map[string]bool)
	existingStaffId := make(map[string]bool)
	existingDisplayName := make(map[string]bool)

	for _, user := range users {
		email := user.GetEmail()
		staffId := user.GetStaffId()
		displayName := user.GetDisplayName()

		if existingEmails[email] {
			return nil, status.Errorf(codes.AlreadyExists, "email already exists")
		}

		if existingStaffId[staffId] {
			return nil, status.Errorf(codes.AlreadyExists, "staff id already exists")
		}

		if existingDisplayName[displayName] {
			return nil, status.Errorf(codes.AlreadyExists, "displayname already exists")
		}

		existingEmails[email] = true
		existingStaffId[staffId] = true
		existingDisplayName[displayName] = true

		var role = ""
		switch user.GetRole() {
		case 1:
			role = "manager"
		case 2:
			role = "bse"
		case 3:
			role = "leader"
		case 4:
			role = "sub leader"
		case 5:
			role = "senior"
		case 6:
			role = "junior"
		default:
			return nil, status.Errorf(codes.InvalidArgument, "Invalid role value")
		}
		nilSlice = append(nilSlice, &models.CreateUserRequest{
			Name:          user.GetName(),
			StaffID:       user.GetStaffId(),
			Email:         user.GetEmail(),
			Profile:       user.GetProfile(),
			DisplayName:   user.GetDisplayName(),
			DepartmentId:  user.GetDepartmentId(),
			Password:      user.GetPassword(),
			MailSubscribe: user.GetMailSubscribe(),
			TeamId:        user.GetTeamId(),
			Role:          role,
			Phone:         user.GetPhone(),
			Address:       user.GetAddress(),
			AboutMe:       user.GetAboutMe(),
			Dob:           user.GetDob().AsTime().Local(),
		})
	}

	newUsers, err := userServer.userService.CreateUsersWithCsv(nilSlice)

	if err != nil {
		if strings.Contains(err.Error(), "email") {
			return nil, status.Errorf(codes.AlreadyExists, consts.UserEmailExists)
		}
		if strings.Contains(err.Error(), "staff_id") {
			return nil, status.Errorf(codes.AlreadyExists, consts.UserStaffIdExists)
		}
		if strings.Contains(err.Error(), "displayname") {
			return nil, status.Errorf(codes.AlreadyExists, consts.UserDisplayNameExists)
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	userResponses := []*pb.UserResponse{}
	for _, user := range newUsers {
		res := &pb.UserResponse{
			XId:          user.Id.Hex(),
			StaffId:      user.StaffID,
			Name:         user.Name,
			Email:        user.Email,
			Profile:      user.Profile,
			DisplayName:  user.DisplayName,
			DepartmentId: user.DepartmentId,
			Role:         user.Role,
			TeamId:       user.TeamId,
			AboutMe:      user.AboutMe,
			Phone:        user.Phone,
			Address:      user.Address,
			Deleted:      wrapperspb.Bool(user.Deleted),
			CreatedAt:    timestamppb.New(user.CreateAt),
			UpdatedAt:    timestamppb.New(user.UpdatedAt),
		}
		userResponses = append(userResponses, res)
	}
	result := &pb.UserWithCsvResponse{
		Users: userResponses,
	}
	return result, nil
}

func (userServer *UserServer) UploadImage(ctx context.Context, req *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {
	response, err := userServer.userService.UploadImage(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to upload image: %v", err)
	}
	return response, nil
}

func (userServer *UserServer) GetUserCount(context context.Context, req *pb.FilterUserRequest) (*pb.UserCountResponse, error) {
	user := &models.FilterUserRequest{
		DisplayName:  req.GetDisplayName(),
		Email:        req.GetEmail(),
		DepartmentId: req.GetDepartmentId(),
		TeamId:       req.GetTeamId(),
		Name:         req.GetName(),
	}

	users := userServer.userService.GetUserCount(user)
	res := &pb.UserCountResponse{
		Count: int64(users),
	}
	return res, nil
}

// func (userServer *UserServer) GetUserNotiCount(context context.Context, req *pb.UserSummaryRequest) (*pb.UserNotiCountResponse, error) {
// 	userId := req.GetUserId()
// 	fmt.Print(userId)
// 	user := userServer.userService.GetUserNotiCount(userId)
// 	fmt.Print(user)
// 	res := &pb.UserNotiCountResponse{
// 		Count: int64(user),
// 	}
// 	fmt.Print("res", res)
// 	return res, nil
// }
