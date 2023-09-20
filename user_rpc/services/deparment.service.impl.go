package services

import (
	"context"
	"errors"
	"time"

	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/consts"
	models "github.com/scm-dev1dev5/mtm-community-forum/user_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeparmentServiceImpl struct {
	deparmentCollection *mongo.Collection
	ctx                 context.Context
}

func NewDeparmentService(deparmentCollection *mongo.Collection, ctx context.Context) DeparmentService {
	return &DeparmentServiceImpl{deparmentCollection, ctx}
}

func (p *DeparmentServiceImpl) CreateDeparment(deparment *models.CreateDeparmentRequest) (*models.DBDeparment, error) {
	deparment.CreateAt = time.Now()
	deparment.UpdatedAt = deparment.CreateAt

	res, err := p.deparmentCollection.InsertOne(p.ctx, deparment)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New(consts.DepartmentExists)
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}

	if _, err := p.deparmentCollection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, err
	}

	var newDeparment *models.DBDeparment
	query := bson.M{"_id": res.InsertedID}
	if err = p.deparmentCollection.FindOne(p.ctx, query).Decode(&newDeparment); err != nil {
		return nil, err
	}
	return newDeparment, nil
}

func (p *DeparmentServiceImpl) GetDeparment(id string) (*models.DBDeparment, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var deparment *models.DBDeparment

	if err := p.deparmentCollection.FindOne(p.ctx, query).Decode(&deparment); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(consts.DepartmentNotFound)
		}

		return nil, err
	}

	return deparment, nil
}

func (p *DeparmentServiceImpl) GetDeparmentList(page int, limit int) []*models.DBDeparment {
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

	query := bson.M{}

	cursor, err := p.deparmentCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil
	}

	defer cursor.Close(p.ctx)

	var deparments []*models.DBDeparment

	for cursor.Next(p.ctx) {
		deparment := &models.DBDeparment{}
		err := cursor.Decode(deparment)

		if err != nil {
			return nil
		}

		deparments = append(deparments, deparment)
	}

	if err := cursor.Err(); err != nil {
		return []*models.DBDeparment{}
	}

	if len(deparments) == 0 {
		return []*models.DBDeparment{}
	}

	return deparments
}

func (p *DeparmentServiceImpl) UpdateDeparment(id string, data *models.UpdateDeparment) (*models.DBDeparment, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.deparmentCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedDeparment *models.DBDeparment
	if err := res.Decode(&updatedDeparment); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		} else {
			return nil, err
		}
	}

	return updatedDeparment, nil
}

func (p *DeparmentServiceImpl) DeleteDeparment(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.deparmentCollection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New(consts.DepartmentNotFound)
	}

	return nil
}

func (p *DeparmentServiceImpl) GetDepartmentCount(page int, limit int) (count int64) {
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

	query := bson.M{}

	res, err := p.deparmentCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}
