package bookmark_service

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

type BookmarkServiceImpl struct {
	bookmarkCollection *mongo.Collection
	ctx                context.Context
}

// CreateBookmark implements BookmarkService.
func (csi *BookmarkServiceImpl) CreateBookmark(bookmark_request *models.CreateBookmarkRequest) (*models.DBBookmark, error) {

	res, err := csi.bookmarkCollection.InsertOne(csi.ctx, bookmark_request)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("Cannot create bookmark")
		}
		return nil, err
	}

	var newBookmark *models.DBBookmark
	query := bson.M{"_id": res.InsertedID}
	if err = csi.bookmarkCollection.FindOne(csi.ctx, query).Decode(&newBookmark); err != nil {
		return nil, err
	}

	return newBookmark, nil
}

// DeleteBookmark implements BookmarkService.
func (csi *BookmarkServiceImpl) DeleteBookmark(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := csi.bookmarkCollection.DeleteOne(csi.ctx, query)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with the Id exists ...csi delete")
	}

	return nil
}

// GetBookmark implements BookmarkService.
func (csi *BookmarkServiceImpl) GetBookmark(id string) (*models.DBBookmark, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var bookmark *models.DBBookmark

	if err := csi.bookmarkCollection.FindOne(csi.ctx, query).Decode(&bookmark); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no documents with that Id exists")
		}

		return nil, err
	}

	return bookmark, nil
}

// GetBookmarks implements BookmarkService.
func (csi *BookmarkServiceImpl) GetBookmarks(page int, limit int) ([]*models.DBBookmark, error) {
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

	cursor, err := csi.bookmarkCollection.Find(csi.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(csi.ctx)

	var bookmarks []*models.DBBookmark

	for cursor.Next(csi.ctx) {
		bookmark := &models.DBBookmark{}
		err := cursor.Decode(bookmark)

		if err != nil {
			return nil, err
		}

		bookmarks = append(bookmarks, bookmark)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(bookmarks) == 0 {
		return []*models.DBBookmark{}, nil
	}

	return bookmarks, nil
}

// UpdateBookmark implements BookmarkService.
func (csi *BookmarkServiceImpl) UpdateBookmark(id string, data *models.UpdateBookmark) (*models.DBBookmark, error) {
	doc, err := utils.ToDoc(data)

	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}

	res := csi.bookmarkCollection.FindOneAndUpdate(csi.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedBookmark *models.DBBookmark

	if err := res.Decode(&updatedBookmark); err != nil {
		return nil, errors.New("no bookmark with that Id exists")
	}

	return updatedBookmark, nil
}

func (csi *BookmarkServiceImpl) GetBookmarksByUserId(id string, page int, limit int) ([]*models.DBBookmark, error) {
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

	query := bson.D{{Key: "user_id", Value: id}}

	cursor, err := csi.bookmarkCollection.Find(csi.ctx, query, &opt)
	var bookmarks []*models.DBBookmark

	if err != nil {
		return nil, err
	}

	defer cursor.Close(csi.ctx)

	for cursor.Next(csi.ctx) {
		question := &models.DBBookmark{}
		err := cursor.Decode(question)
		if err != nil {
			return nil, err
		}

		bookmarks = append(bookmarks, question)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(bookmarks) == 0 {
		return []*models.DBBookmark{}, nil
	}

	return bookmarks, err
}

func (bsi *BookmarkServiceImpl) GetBookmarkCount(id string) (count int64) {
	query := bson.D{{Key: "user_id", Value: id}}
	res, err := bsi.bookmarkCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}

func (csi *BookmarkServiceImpl) GetBookmarksByUserIdQuestionId(id string, question_id string, page int, limit int) ([]*models.DBBookmark, error) {
	query := bson.D{{Key: "user_id", Value: id}, {Key: "question_id", Value: question_id}}

	cursor, err := csi.bookmarkCollection.Find(csi.ctx, query)
	var bookmarks []*models.DBBookmark

	if err != nil {
		return nil, err
	}

	defer cursor.Close(csi.ctx)

	for cursor.Next(csi.ctx) {
		bookmark := &models.DBBookmark{}
		err := cursor.Decode(bookmark)
		if err != nil {
			return nil, err
		}

		bookmarks = append(bookmarks, bookmark)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(bookmarks) == 0 {
		return []*models.DBBookmark{}, nil
	}

	return bookmarks, err
}

func NewBookmarkService(bookmarkCollection *mongo.Collection, ctx context.Context) BookmarkService {
	return &BookmarkServiceImpl{bookmarkCollection, ctx}
}
