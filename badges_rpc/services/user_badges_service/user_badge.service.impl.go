package services

import (
	"context"
	"errors"
	"time"

	"github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/consts"
	models "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserBadgeServiceImpl struct {
	userBadgeCollection *mongo.Collection
	ctx                 context.Context
}

func NewUserBadgeService(userBadgeCollection *mongo.Collection, ctx context.Context) UserBadgeService {
	return &UserBadgeServiceImpl{userBadgeCollection, ctx}
}

func (p *UserBadgeServiceImpl) CreateUserBadge(userBadge *models.CreateUserBadgeRequest) (*models.DBUserBadge, error) {
	userBadge.CreatedAt = time.Now()
	userBadge.UpdatedAt = userBadge.CreatedAt

	res, err := p.userBadgeCollection.InsertOne(p.ctx, userBadge)

	opt := options.Index()
	opt.SetUnique(true)

	// if err != nil {
	// 	if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
	// 		return nil, errors.New(consts.UserBadgeExists)
	// 	}
	// 	return nil, err
	// }

	var newUserBadge *models.DBUserBadge
	query := bson.M{"_id": res.InsertedID}
	if err = p.userBadgeCollection.FindOne(p.ctx, query).Decode(&newUserBadge); err != nil {
		return nil, err
	}

	return newUserBadge, nil
}

func (p *UserBadgeServiceImpl) GetUserBadge(userId string, badgeId string) (*models.DBUserBadge, error) {

	query := bson.M{
		"user_id":  userId,
		"badge_id": badgeId,
	}

	var userBadge *models.DBUserBadge

	if err := p.userBadgeCollection.FindOne(p.ctx, query).Decode(&userBadge); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(consts.NotFoundInDB)
		}

		return nil, err
	}

	return userBadge, nil
}

func (p *UserBadgeServiceImpl) GetUserBadges(page int, limit int) ([]*models.DBUserBadge, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))
	opt.SetSort(bson.M{"_id": 1})

	query := bson.M{}

	cursor, err := p.userBadgeCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var userBadges []*models.DBUserBadge

	for cursor.Next(p.ctx) {
		userBadge := &models.DBUserBadge{}
		err := cursor.Decode(userBadge)

		if err != nil {
			return nil, err
		}

		userBadges = append(userBadges, userBadge)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(userBadges) == 0 {
		return []*models.DBUserBadge{}, nil
	}

	return userBadges, nil
}

func (p *UserBadgeServiceImpl) GetUserBadgesOfUser(userId string) ([]*models.DBUserBadge, error) {

	opt := options.FindOptions{}
	opt.SetSort(bson.M{"_id": 1})

	query := bson.M{"user_id": userId}

	cursor, err := p.userBadgeCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var userBadges []*models.DBUserBadge

	for cursor.Next(p.ctx) {
		userBadge := &models.DBUserBadge{}
		err := cursor.Decode(userBadge)

		if err != nil {
			return nil, err
		}

		userBadges = append(userBadges, userBadge)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(userBadges) == 0 {
		return []*models.DBUserBadge{}, nil
	}

	return userBadges, nil
}

func (p *UserBadgeServiceImpl) UpdateUserBadge(userId string, data *models.UpdateUserBadge) (*models.DBUserBadge, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.userBadgeCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))
	var updatedUserBadge *models.DBUserBadge
	if err := res.Decode(&updatedUserBadge); err != nil {
		return nil, errors.New(consts.NotFoundInDB)
	}

	return updatedUserBadge, nil
}

func (p *UserBadgeServiceImpl) DeleteUserBadge(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.userBadgeCollection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New(consts.NotFoundInDB)
	}

	return nil
}

func (p *UserBadgeServiceImpl) GetBadgeCount(id string) (count int64) {
	query := bson.D{{Key: "user_id", Value: id}}

	res, err := p.userBadgeCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}
