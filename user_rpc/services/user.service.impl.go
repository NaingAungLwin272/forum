package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	bcrypt "golang.org/x/crypto/bcrypt"

	auth "github.com/scm-dev1dev5/mtm-community-forum/user_rpc/auth"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/consts"
	models "github.com/scm-dev1dev5/mtm-community-forum/user_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
	jwtManager     *auth.JWTManager
}

func (p *UserServiceImpl) UploadImage(ureq *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {
	// Generate a unique file name for the image
	uniqueFileName := generateUniqueFileName()

	// Create a new file with the unique file name
	filePath := "uploads/" + uniqueFileName
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Write the file content to the newly created file
	_, err = file.Write(ureq.FileContent)
	if err != nil {
		return nil, err
	}

	// Return the URL of the uploaded file
	fileURL := fmt.Sprintf("/uploads/%s", uniqueFileName)
	return &pb.FileUploadResponse{FileUrl: fileURL}, nil
}

// Helper function to generate a unique file name
func generateUniqueFileName() string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%d", timestamp)
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context, jwtManager *auth.JWTManager) UserService {
	return &UserServiceImpl{userCollection, ctx, jwtManager}
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func (p *UserServiceImpl) CreateUser(user *models.CreateUserRequest) (*models.DBUser, error) {
	user.CreateAt = time.Now()
	user.UpdatedAt = user.CreateAt
	user.Deleted = false
	user.MailSubscribe = true
	hashed, _ := HashPassword(user.Password)
	user.Password = hashed

	res, err := p.userCollection.InsertOne(p.ctx, user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, err
		}
		return nil, err
	}

	indexEmail := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true)}
	if _, err := p.userCollection.Indexes().CreateOne(p.ctx, indexEmail); err != nil {
		return nil, errors.New(consts.UserCreateErr)
	}

	indexDisplayName := mongo.IndexModel{Keys: bson.M{"displayname": 1}, Options: options.Index().SetUnique(true)}
	if _, err := p.userCollection.Indexes().CreateOne(p.ctx, indexDisplayName); err != nil {
		return nil, errors.New("could not create index for displayname")
	}

	indexStaffID := mongo.IndexModel{Keys: bson.M{"staff_id": 1}, Options: options.Index().SetUnique(true)}
	if _, err := p.userCollection.Indexes().CreateOne(p.ctx, indexStaffID); err != nil {
		return nil, errors.New("could not create index for staffid")
	}

	var newUser *models.DBUser
	query := bson.M{"_id": res.InsertedID}
	if err = p.userCollection.FindOne(p.ctx, query).Decode(&newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (p *UserServiceImpl) GetUser(id string) (*models.DBUser, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var user *models.DBUser

	if err := p.userCollection.FindOne(p.ctx, query).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(consts.UserNotFound)
		}

		return nil, err
	}

	return user, nil
}

func (p *UserServiceImpl) GetUsers(page int, limit int) ([]*models.DBUser, error) {
	if page == 0 {
		page = 0
	}

	if limit == 0 {
		limit = 0
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	if limit > 0 {
		opt.SetLimit(int64(limit))
		opt.SetSkip(int64(skip))
		opt.SetSort(bson.M{"created_at": -1})
	}

	query := bson.M{"is_deleted": false}
	cursor, err := p.userCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx)

	var users []*models.DBUser
	for cursor.Next(p.ctx) {
		user := &models.DBUser{}
		err := cursor.Decode(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return []*models.DBUser{}, nil
	}

	return users, nil
}

func (p *UserServiceImpl) UpdateUser(id string, data *models.UpdateUser, ctx context.Context) (*models.DBUser, error) {
	// claims, err := utils.CheckAuth(ctx, p.jwtManager)
	// if err != nil {
	// 	return nil, err
	// }

	// if (claims.Id != id && claims.Role != "manager") {
	// 	return nil, errors.New(consts.UnAuthorizedUser)
	// }

	// if (claims.Role != "manager") {
	// 	data.Role = ""
	// 	data.StaffID = ""
	// }

	// log.Println("id", id)
	if data.Password != "" {
		hashed, err := HashPassword(data.Password)
		if err != nil {
			return nil, err
		}
		data.Password = hashed
	} else {
		data.Password = ""
	}

	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}, {Key: "deleted", Value: bson.D{{Key: "$ne", Value: true}}}}
	update := bson.D{{Key: "$set", Value: doc}}

	res := p.userCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedUser *models.DBUser
	if err := res.Decode(&updatedUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		} else {
			return nil, err
		}
	}

	return updatedUser, nil
}

func (p *UserServiceImpl) DeleteUser(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	delete := bson.D{{Key: "$set", Value: bson.D{{Key: "is_deleted", Value: true}}}}

	res := p.userCollection.FindOneAndUpdate(p.ctx, query, delete, options.FindOneAndUpdate().SetReturnDocument(1))
	var deletedUser *models.DBUser
	if err := res.Decode(&deletedUser); err != nil {
		return errors.New(consts.UserNotFound)
	}

	return nil
}

func (p *UserServiceImpl) FilterUser(user *models.FilterUserRequest) ([]*models.DBUser, error) {
	if user.Page == 0 {
		user.Page = 0
	}

	if user.Limit == 0 {
		user.Limit = 0
	}

	skip := (user.Page - 1) * user.Limit

	opt := options.FindOptions{}
	if user.Limit > 0 {
		opt.SetLimit(int64(user.Limit))
		opt.SetSkip(int64(skip))
		opt.SetSort(bson.M{"created_at": -1})
	}

	if user.DepartmentId == nil && user.TeamId == nil && user.Name == "" && user.Email == "" {
		query := bson.M{"is_deleted": false}
		var allUsers []*models.DBUser
		cursor, err := p.userCollection.Find(p.ctx, query, &opt)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(p.ctx)

		for cursor.Next(p.ctx) {
			user := &models.DBUser{}
			err := cursor.Decode(user)
			if err != nil {
				return nil, err
			}
			allUsers = append(allUsers, user)
		}

		if err := cursor.Err(); err != nil {
			return nil, err
		}

		return allUsers, nil
	}
	filter := bson.D{}

	query := []bson.M{}

	if user.DepartmentId != nil {
		query = append(query, bson.M{"department_id": bson.M{"$in": user.DepartmentId}})
	}

	if user.TeamId != nil {
		query = append(query, bson.M{"team_id": bson.M{"$in": user.TeamId}})
	}

	if user.Name != "" {
		query = append(query, bson.M{"name": bson.M{"$regex": user.Name, "$options": "i"}})
	}

	if user.Email != "" {
		query = append(query, bson.M{"email": bson.M{"$regex": user.Email, "$options": "i"}})
	}

	filter = append(filter, bson.E{Key: "$and", Value: query})

	var users []*models.DBUser
	cursor, err := p.userCollection.Find(p.ctx, filter, &opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx)

	for cursor.Next(p.ctx) {
		user := &models.DBUser{}
		err := cursor.Decode(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return []*models.DBUser{}, nil
	}

	return users, nil
}

func (p *UserServiceImpl) CreateUsersWithCsv(users []*models.CreateUserRequest) ([]*models.DBUser, error) {
	var dbUsers []*models.DBUser
	var tmpUserData bool = false
	for _, existingUser := range users {
		existingQuery := bson.M{
			"$or": []bson.M{
				{"email": existingUser.Email},
				{"displayname": existingUser.DisplayName},
				{"staff_id": existingUser.StaffID},
			},
		}
		existingCount, err := p.userCollection.CountDocuments(p.ctx, existingQuery, nil)
		if err != nil {
			return nil, err
		}

		if existingCount > 0 {
			tmpUserData = true
			break
		}
	}

	if tmpUserData {
		return nil, errors.New(consts.UserAlreadyExists)
	}

	for _, user := range users {
		user.CreateAt = time.Now()
		user.UpdatedAt = user.CreateAt
		user.Deleted = false
		hashed, _ := HashPassword(user.Password)
		user.Password = hashed

		res, err := p.userCollection.InsertOne(p.ctx, user)
		if err != nil {
			if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
				return nil, err
			}
			return nil, err
		}

		indexEmail := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true)}
		_, err = p.userCollection.Indexes().CreateOne(p.ctx, indexEmail)
		if err != nil {
			return nil, errors.New(consts.UserCreateErr)
		}

		indexDisplayName := mongo.IndexModel{Keys: bson.M{"displayname": 1}, Options: options.Index().SetUnique(true)}
		_, err = p.userCollection.Indexes().CreateOne(p.ctx, indexDisplayName)
		if err != nil {
			return nil, errors.New("could not create index for displayname")
		}

		indexStaffID := mongo.IndexModel{Keys: bson.M{"staff_id": 1}, Options: options.Index().SetUnique(true)}
		_, err = p.userCollection.Indexes().CreateOne(p.ctx, indexStaffID)
		if err != nil {
			return nil, errors.New("could not create index for staffid")
		}
		var newUser *models.DBUser
		query := bson.M{"_id": res.InsertedID}
		if err = p.userCollection.FindOne(p.ctx, query).Decode(&newUser); err != nil {
			return nil, err
		}

		dbUsers = append(dbUsers, newUser)
	}

	return dbUsers, nil
}

func (p *UserServiceImpl) GetUserCount(user *models.FilterUserRequest) (count int64) {

	if user.DepartmentId == nil && user.TeamId == nil && user.Name == "" && user.Email == "" {
		query := bson.M{"is_deleted": false}

		res, err := p.userCollection.CountDocuments(context.TODO(), query)
		if err != nil {
			panic(err)
		}
		return res
	}
	filter := bson.D{}

	query := []bson.M{}

	if user.DepartmentId != nil {
		query = append(query, bson.M{"department_id": bson.M{"$in": user.DepartmentId}})
	}

	if user.TeamId != nil {
		query = append(query, bson.M{"team_id": bson.M{"$in": user.TeamId}})
	}

	if user.Name != "" {
		query = append(query, bson.M{"name": bson.M{"$regex": user.Name, "$options": "i"}})
	}

	if user.Email != "" {
		query = append(query, bson.M{"email": bson.M{"$regex": user.Email, "$options": "i"}})
	}

	filter = append(filter, bson.E{Key: "$and", Value: query})

	res, err := p.userCollection.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	return res
}

func (p *UserServiceImpl) GetUserByDisplayName(displayName string) (*models.DBUser, error) {
	query := bson.M{"displayname": displayName}
	var user *models.DBUser

	if err := p.userCollection.FindOne(p.ctx, query).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(consts.UserNotFound)
		}
		return nil, err
	}

	return user, nil
}
