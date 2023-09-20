package services

import (
	"context"
	"errors"
	"time"

	"github.com/scm-dev1dev5/mtm-community-forum/category_rpc/consts"
	models "github.com/scm-dev1dev5/mtm-community-forum/category_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/category_rpc/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryServiceImpl struct {
	categoryCollection *mongo.Collection
	ctx                context.Context
}

func NewCategoryService(categoryCollection *mongo.Collection, ctx context.Context) CategoryService {
	return &CategoryServiceImpl{categoryCollection, ctx}
}

func (p *CategoryServiceImpl) CreateCategory(category *models.CreateCategoryRequest) (*models.DBCategory, error) {
	category.CreateAt = time.Now()
	category.UpdatedAt = category.CreateAt

	res, err := p.categoryCollection.InsertOne(p.ctx, category)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New(consts.CategoryExists)
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}

	if _, err := p.categoryCollection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, errors.New(consts.CategoryCreateErr)
	}

	var newCategory *models.DBCategory
	query := bson.M{"_id": res.InsertedID}
	if err = p.categoryCollection.FindOne(p.ctx, query).Decode(&newCategory); err != nil {
		return nil, err
	}

	return newCategory, nil
}

func (p *CategoryServiceImpl) GetCategory(id string) (*models.DBCategory, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var category *models.DBCategory

	if err := p.categoryCollection.FindOne(p.ctx, query).Decode(&category); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(consts.NotFound)
		}

		return nil, err
	}

	return category, nil
}

func (p *CategoryServiceImpl) GetCategories(page int, limit int) ([]*models.DBCategory, error) {
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

	cursor, err := p.categoryCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var categories []*models.DBCategory

	for cursor.Next(p.ctx) {
		category := &models.DBCategory{}
		err := cursor.Decode(category)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return []*models.DBCategory{}, nil
	}

	return categories, nil
}

func (p *CategoryServiceImpl) UpdateCategory(id string, data *models.UpdateCategory) (*models.DBCategory, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.categoryCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedCategory *models.DBCategory
	if err := res.Decode(&updatedCategory); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		} else {
			return nil, err
		}
	}

	return updatedCategory, nil
}

func (p *CategoryServiceImpl) DeleteCategory(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.categoryCollection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New(consts.NotFound)
	}

	return nil
}

func (p *CategoryServiceImpl) GetCategoryByType(page int64, limit int64, categoryType int32) ([]*models.DBCategory, error) {
	query := bson.M{"type": categoryType}

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

	cursor, err := p.categoryCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx)

	var categories []*models.DBCategory
	for cursor.Next(p.ctx) {
		category := &models.DBCategory{}
		err := cursor.Decode(category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (p *CategoryServiceImpl) GetCategoryCount(page int, limit int) (languageCount int64, tagCount int64) {
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

	languageQuery := bson.M{"type": 1}
	tagQuery := bson.M{"type": 2}

	languageCount, err := p.categoryCollection.CountDocuments(context.TODO(), languageQuery)
	if err != nil {
		panic(err)
	}

	tagCount, err = p.categoryCollection.CountDocuments(context.TODO(), tagQuery)
	if err != nil {
		panic(err)
	}

	return languageCount, tagCount
}
