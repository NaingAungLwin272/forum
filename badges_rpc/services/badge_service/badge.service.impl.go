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

type BadgeServiceImpl struct {
	badgeCollection *mongo.Collection
	ctx             context.Context
}

func NewBadgeService(badgeCollection *mongo.Collection, ctx context.Context) BadgeService {
	return &BadgeServiceImpl{badgeCollection, ctx}
}

func (p *BadgeServiceImpl) CreateBadge(badge *models.CreateBadgeRequest) (*models.DBBadge, error) {
	badge.CreateAt = time.Now()
	badge.UpdatedAt = badge.CreateAt

	res, err := p.badgeCollection.InsertOne(p.ctx, badge)

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}

	if _, err := p.badgeCollection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, errors.New(consts.BadgeCreateErr)
	}

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New(consts.BadgeExists)
		}
		return nil, err
	}

	var newBadge *models.DBBadge
	query := bson.M{"_id": res.InsertedID}
	if err = p.badgeCollection.FindOne(p.ctx, query).Decode(&newBadge); err != nil {
		return nil, err
	}

	return newBadge, nil
}

func (p *BadgeServiceImpl) GetBadge(id string) (*models.DBBadge, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var badge *models.DBBadge

	if err := p.badgeCollection.FindOne(p.ctx, query).Decode(&badge); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(consts.NotFoundInDB)
		}

		return nil, err
	}

	return badge, nil
}

func (p *BadgeServiceImpl) GetBadges(page int, limit int) ([]*models.DBBadge, error) {
	// if page == 0 {
	// 	page = 0
	// }

	// if limit == 0 {
	// 	limit = 0
	// }

	// skip := (page - 1) * limit

	// if limit > 0 {
	// 	opt.SetLimit(int64(limit))
	// 	opt.SetSkip(int64(skip))
	// 	opt.SetSort(bson.M{"type": 1, "level": })
	// }
	// opt := options.FindOptions{}
	opts := options.Find().SetSort(bson.D{{Key: "type", Value: 1}, {Key: "level", Value: 1}})

	query := bson.M{}

	cursor, err := p.badgeCollection.Find(p.ctx, query, opts)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var badges []*models.DBBadge

	for cursor.Next(p.ctx) {
		badge := &models.DBBadge{}
		err := cursor.Decode(badge)

		if err != nil {
			return nil, err
		}

		badges = append(badges, badge)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(badges) == 0 {
		return []*models.DBBadge{}, nil
	}

	return badges, nil
}

func (p *BadgeServiceImpl) UpdateBadge(id string, data *models.UpdateBadge) (*models.DBBadge, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.badgeCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedBadge *models.DBBadge
	if err := res.Decode(&updatedBadge); err != nil {
		return nil, errors.New(consts.NotFoundInDB)
	}

	return updatedBadge, nil
}

func (p *BadgeServiceImpl) DeleteBadge(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.badgeCollection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New(consts.NotFoundInDB)
	}

	return nil
}
