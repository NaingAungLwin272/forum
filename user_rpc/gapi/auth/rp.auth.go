package auth_gapi

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/consts"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

/*
Steps & Logics

 1. Get Email from request data
 2. Find user data with given email
 3. Genarate Token witn user data
 4. Return Token if generate token successfully
*/
func (authServer *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	lengthOfEmail := len(req.GetEmail())
	lengthOfStaffId := len(req.GetStaffId())

	if lengthOfEmail < 1 && lengthOfStaffId < 1 {
		return nil, status.Error(codes.InvalidArgument, "Email or StaffId is required")
	}

	query := bson.M{
		"$or": []bson.M{
			{"email": req.GetEmail()},
			{"staff_id": req.GetStaffId()},
		},
	}
	var user *models.DBUser
	if err := authServer.userCollection.FindOne(ctx, query).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "cannot find user with that email")
		}
	}

	log.Println(user, "User data in login-----")
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword()))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	var tokenDuration time.Duration

	if req.GetIsRememberMe() == true {
		tokenDuration = 720 * time.Hour
	} else {
		tokenDuration = 8 * time.Hour
	}

	token, err := authServer.jwtManager.Generate(user, tokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't generate access token")
	}

	userResponse := &pb.UserResponse{
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
		LastLogin:     timestamppb.New(user.LastLogin),
		LastPost:      timestamppb.New(user.LastPost),
		CreatedAt:     timestamppb.New(user.CreateAt),
		UpdatedAt:     timestamppb.New(user.UpdatedAt),
	}
	res := &pb.LoginResponse{AccessToken: token, User: userResponse}

	log.Println(user.Role, "user.Role")
	log.Println(userResponse.Role, "userResponse.Role")
	return res, nil
}

/*
Steps & Logics
 1. Get token from request data
 2. Check Token is valid or not
 3. Return true if valid
*/
func (authServer *AuthServer) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	lengthOfToken := len(req.GetAccessToken())

	if lengthOfToken < 1 {
		return nil, status.Error(codes.InvalidArgument, "Token is required")
	}
	_, ok, err := authServer.jwtManager.Verify(req.GetAccessToken())

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.VerifyTokenResponse{IsTokenVerified: ok}

	return res, nil
}

/*
Steps & Logics

 1. Create user by given user data
 2. Return true if create successfully
*/
func (authServer *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	responseUnsuccess := &pb.RegisterResponse{
		IsSuccess: false,
	}

	response := &pb.RegisterResponse{
		IsSuccess: true,
	}

	lengthOfName := len(req.GetName())
	lengthOfPassword := len(req.GetPassword())
	lengthOfEmail := len(req.GetEmail())

	if lengthOfPassword < 1 || lengthOfName < 1 || lengthOfEmail < 1 {
		return responseUnsuccess, nil
	}

	hashedPassword, hashederror := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	fmt.Println(hashederror, "hashederror is here")

	user := &models.CreateUserRequest{
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		StaffID:   req.GetStaffId(),
		Password:  string(hashedPassword),
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
	_, err := authServer.userCollection.InsertOne(ctx, user)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return responseUnsuccess, status.Errorf(codes.AlreadyExists, "user with that email already exists")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	indexEmail := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true)}
	if _, err := authServer.userCollection.Indexes().CreateOne(ctx, indexEmail); err != nil {
		return nil, errors.New(consts.UserCreateErr)
	}

	indexStaffID := mongo.IndexModel{Keys: bson.M{"staff_id": 1}, Options: options.Index().SetUnique(true)}
	if _, err := authServer.userCollection.Indexes().CreateOne(ctx, indexStaffID); err != nil {
		return nil, errors.New("could not create index for staffid")
	}

	return response, nil
}

/*
Steps & Logics

 1. Find user data with given email
 2. Generate token if user data exists ,
 3. Return token if generates successfully
*/
func (authServer *AuthServer) ForgetPassword(ctx context.Context, req *pb.ForgetPasswordRequest) (*pb.ForgetPasswordResponse, error) {
	lengthOfEmail := len(req.GetEmail())
	if lengthOfEmail < 1 {
		return nil, status.Error(codes.InvalidArgument, "Email is required")
	}
	query := bson.M{"email": req.GetEmail()}
	var user *models.DBUser
	if err := authServer.userCollection.FindOne(ctx, query).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "cannot find user with that email")
		}
	}

	tokenDuration := 5 * time.Minute

	token, err := authServer.forgetPassword.Generate(user, tokenDuration)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't generate access token")
	}

	res := &pb.ForgetPasswordResponse{
		Token:  token,
		Name:   user.Name,
		Origin: req.GetOrigin(),
	}

	return res, err
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

/*
Steps & Logics

 1. Find user data with given email
 2. Get user "_id" if exists ,
 3. Update password by finding user data with id
 4. Return true if update successfully
*/
func (authServer *AuthServer) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	responseUnsuccess := &pb.ResetPasswordResponse{
		IsSuccess: false,
	}

	// MAKE SURE DATA IS INCLUDED IN PAYLOAD
	lengthOfPassword := len(req.GetPassword())
	lengthOfEmail := len(req.GetEmail())
	if lengthOfPassword < 1 || lengthOfEmail < 1 {
		return responseUnsuccess, status.Error(codes.InvalidArgument, "Email and Password are required")
	}

	// VERIFY TOKEN
	userClaims, ok, err := authServer.forgetPassword.Verify(req.GetToken())
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, err.Error())
	}

	// CHECK EMAIL FROM TOKEN IS EQUAL TO EMAIL FROM USERCLAIMS
	if userClaims.Email != req.GetEmail() {
		return nil, status.Error(codes.PermissionDenied, "Invalid Email")
	}

	// HASHED GIVEN NEW PASSWORD
	hashed, err := HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error hashing new password: %v", err)
	}

	queryFind := bson.M{"email": req.GetEmail()}
	var user *models.DBUser
	// FIND USER WITH EMAIL
	if err := authServer.userCollection.FindOne(ctx, queryFind).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "cannot find user with that email: %v", err)
		}
	}
	obId, _ := primitive.ObjectIDFromHex(user.Id.Hex())
	query := bson.D{{Key: "_id", Value: obId}, {Key: "is_deleted", Value: bson.D{{Key: "$ne", Value: true}}}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: hashed}}}}
	//	UPDATE USER WITH USER ID
	_, errorOfUpdate := authServer.userCollection.UpdateOne(ctx, query, update)

	if errorOfUpdate != nil {
		return nil, status.Errorf(codes.NotFound, "cannot update: %v", errorOfUpdate)
	}
	res := &pb.ResetPasswordResponse{
		IsSuccess: ok,
	}

	return res, nil
}

func (authServer *AuthServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	responseUnsuccess := &pb.ChangePasswordResponse{
		IsSuccess: false,
	}

	lengthOfUserID := len(req.GetUserId())
	lengthOfCurrentPassword := len(req.GetPassword())
	lengthOfNewPassword := len(req.GetNewPassword())
	lengthOfConfirmPassword := len(req.GetConfrimPassword())
	if lengthOfUserID < 1 || lengthOfCurrentPassword < 1 || lengthOfNewPassword < 1 || lengthOfConfirmPassword < 1 {
		return responseUnsuccess, status.Error(codes.InvalidArgument, "UserID, Current Password, New Password, and Confirm Password are required")
	}

	// FIND USER BY ID
	userID, err := primitive.ObjectIDFromHex(req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid UserID")
	}

	var user *models.DBUser
	queryFind := bson.M{"_id": userID}
	if err := authServer.userCollection.FindOne(ctx, queryFind).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "User not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Database error: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword()))
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "Invalid current password")
	}

	hashedNewPassword, err := HashPassword(req.GetNewPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error hashing new password: %v", err)
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: hashedNewPassword}}}}
	filter := bson.D{{Key: "_id", Value: userID}, {Key: "is_deleted", Value: bson.D{{Key: "$ne", Value: true}}}}
	_, err = authServer.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error updating password: %v", err)
	}

	res := &pb.ChangePasswordResponse{
		IsSuccess: true,
	}

	return res, nil
}
