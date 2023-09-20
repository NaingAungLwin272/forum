package view_service

import (
	"context"
	"errors"

	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ViewServiceImpl struct {
	viewCollection *mongo.Collection
	ctx            context.Context
}

// CreateView implements ViewService.
func (vsi *ViewServiceImpl) CreateView(vsi_request *models.CreateViewRequest) (*models.DBView, error) {
	res, err := vsi.viewCollection.InsertOne(vsi.ctx, vsi_request)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("Cannot create view")
		}
		return nil, err
	}

	var newView *models.DBView
	query := bson.M{"_id": res.InsertedID}
	if err = vsi.viewCollection.FindOne(vsi.ctx, query).Decode(&newView); err != nil {
		return nil, err
	}

	return newView, nil
}

// DeleteView implements ViewService.
func (vsi *ViewServiceImpl) DeleteView(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := vsi.viewCollection.DeleteOne(vsi.ctx, query)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with the Id exists ...vsi delete")
	}

	return nil
}

// GetView implements ViewService.
func (vsi *ViewServiceImpl) GetView(id string) (*models.DBView, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var view *models.DBView

	if err := vsi.viewCollection.FindOne(vsi.ctx, query).Decode(&view); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no documents with that Id exists")
		}

		return nil, err
	}

	return view, nil
}

// GetViews implements ViewService.
func (vsi *ViewServiceImpl) GetViews(page int, limit int) ([]*models.DBView, error) {
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
	opt.SetSort(bson.M{"created_at": -1})

	query := bson.M{}

	cursor, err := vsi.viewCollection.Find(vsi.ctx, query, &opt)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(vsi.ctx)

	var views []*models.DBView

	for cursor.Next(vsi.ctx) {
		view := &models.DBView{}
		err := cursor.Decode(view)

		if err != nil {
			return nil, err
		}

		views = append(views, view)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(views) == 0 {
		return []*models.DBView{}, nil
	}

	return views, nil
}

// UpdateView implements ViewService.
func (vsi *ViewServiceImpl) UpdateView(id string, data *models.UpdateView) (*models.DBView, error) {
	doc, err := utils.ToDoc(data)

	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}

	res := vsi.viewCollection.FindOneAndUpdate(vsi.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedView *models.DBView

	if err := res.Decode(&updatedView); err != nil {
		return nil, errors.New("no vote with that Id exsts...")
	}

	return updatedView, nil
}

func (vsi *ViewServiceImpl) GetViewsByUserId(id string) ([]*models.DBView, error) {

	query := bson.D{{Key: "user_id", Value: id}}

	cursor, err := vsi.viewCollection.Find(vsi.ctx, query)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(vsi.ctx)

	var views []*models.DBView

	for cursor.Next(vsi.ctx) {
		view := &models.DBView{}
		err := cursor.Decode(view)
		if err != nil {
			return nil, err
		}

		views = append(views, view)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(views) == 0 {
		return []*models.DBView{}, nil
	}

	return views, err
}

func (vsi *ViewServiceImpl) GetViewsByUserIdQuestionId(id string, question_id string, page int, limit int) ([]*models.DBView, error) {
	query := bson.D{{Key: "user_id", Value: id}, {Key: "question_id", Value: question_id}}

	cursor, err := vsi.viewCollection.Find(vsi.ctx, query)
	var views []*models.DBView

	if err != nil {
		return nil, err
	}

	defer cursor.Close(vsi.ctx)

	for cursor.Next(vsi.ctx) {
		view := &models.DBView{}
		err := cursor.Decode(view)
		if err != nil {
			return nil, err
		}

		views = append(views, view)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(views) == 0 {
		return []*models.DBView{}, nil
	}

	return views, nil
}

func NewViewService(viewCollection *mongo.Collection, ctx context.Context) ViewService {
	return &ViewServiceImpl{viewCollection, ctx}
}
