package services

import (
	"context"
	"errors"
	"time"

	models "github.com/scm-dev1dev5/mtm-community-forum/noti_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/noti_rpc/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NotiServiceImpl struct {
	notiCollection *mongo.Collection
	ctx            context.Context
}

func NewNotiService(notiCollection *mongo.Collection, ctx context.Context) NotiService {
	return &NotiServiceImpl{notiCollection, ctx}
}

// CreateNoti implements NotiService.
func (p *NotiServiceImpl) CreateNoti(noti *models.CreateNotiRequest) (*models.DBNoti, error) {
	dbNoti := &models.DBNoti{
		UserId:      noti.UserId,
		Type:        noti.Type,
		Name:        noti.Name,
		Description: noti.Description,
		Link:        noti.Link,
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
	}
	if noti.UserId == "" {
		return nil, errors.New("userId can't be empty")
	} else if noti.Type == 0 {
		return nil, errors.New("type can't be empty")
	} else if noti.Type <= 0 || noti.Type > 6 {
		return nil, errors.New("type should be 1 to 6")
	}

	if noti.Status != nil {
		dbNoti.Status = noti.Status
	} else {
		defaultStatus := false
		dbNoti.Status = &defaultStatus
	}
	res, err := p.notiCollection.InsertOne(p.ctx, dbNoti)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("noti already exist")
		}
		return nil, err
	}

	var newNoti *models.DBNoti
	query := bson.M{"_id": res.InsertedID}
	if err = p.notiCollection.FindOne(p.ctx, query).Decode(&newNoti); err != nil {
		return nil, err
	}

	return newNoti, nil
}

// GetNoti implements NotiService.
func (p *NotiServiceImpl) GetNoti(id string) (*models.DBNoti, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var noti *models.DBNoti

	if err := p.notiCollection.FindOne(p.ctx, query).Decode(&noti); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no notification with that Id exists")
		}

		return nil, err
	}

	return noti, nil
}

// GetNotis implements NotiService.
func (p *NotiServiceImpl) GetNotis(page int, limit int) ([]*models.DBNoti, error) {
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

	cursor, err := p.notiCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var notis []*models.DBNoti

	for cursor.Next(p.ctx) {
		noti := &models.DBNoti{}
		err := cursor.Decode(noti)

		if err != nil {
			return nil, err
		}
		statusValue := cursor.Current.Lookup("status").Boolean()
		noti.Status = &statusValue
		notis = append(notis, noti)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(notis) == 0 {
		return []*models.DBNoti{}, nil
	}

	return notis, nil
}

// DeleteNoti implements NotiService.
func (p *NotiServiceImpl) DeleteNoti(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.notiCollection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}

// UpdateNoti implements NotiService.
func (p *NotiServiceImpl) UpdateNoti(id string, data *models.UpdateNoti) (*models.DBNoti, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.notiCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedNoti *models.DBNoti
	if err := res.Decode(&updatedNoti); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no notification with that Id exists")
		}
	}

	return updatedNoti, nil
}

// GetNotiByUserId implements NotiService.
func (p *NotiServiceImpl) GetNotiByUserId(id string, page int, limit int) []*models.DBNoti {
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
	query := bson.M{"user_id": id}

	cursor, err := p.notiCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil
	}

	defer cursor.Close(p.ctx)

	var notis []*models.DBNoti

	for cursor.Next(p.ctx) {
		noti := &models.DBNoti{}
		err := cursor.Decode(noti)

		if err != nil {
			return []*models.DBNoti{}
		}

		notis = append(notis, noti)
	}

	if err := cursor.Err(); err != nil {
		return []*models.DBNoti{}
	}

	if len(notis) == 0 {
		return []*models.DBNoti{}
	}

	return notis
}

func (p *NotiServiceImpl) GetNotiCount(id string) (count int64) {
	query := bson.D{
		{Key: "user_id", Value: id},
		{Key: "status", Value: true},
	}

	res, err := p.notiCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}

func (p *NotiServiceImpl) MarkAllNotiAsRead(userID string) error {
	filter := bson.D{
		{Key: "user_id", Value: userID},
		{Key: "status", Value: true},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "status", Value: false}}},
	}

	_, err := p.notiCollection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (p *NotiServiceImpl) GetNotiForUserSummary(id string) (count int64) {
	query := bson.D{
		{Key: "user_id", Value: id},
	}

	res, err := p.notiCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}
