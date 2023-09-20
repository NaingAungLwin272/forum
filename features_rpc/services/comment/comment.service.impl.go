package comment_service

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

type CommentServiceImpl struct {
	commentCollection *mongo.Collection
	ctx               context.Context
}

// CreateComment implements CommentService.
func (csi *CommentServiceImpl) CreateComment(comment_request *models.CreateCommentRequest) (*models.DBComment, error) {
	comment_request.Vote_Count = 0
	res, err := csi.commentCollection.InsertOne(csi.ctx, comment_request)
	if err != nil {
		return nil, err
	}
	var newComment *models.DBComment

	query := bson.M{"_id": res.InsertedID}

	if err = csi.commentCollection.FindOne(csi.ctx, query).Decode(&newComment); err != nil {
		return nil, err
	}

	return newComment, nil
}

// DeleteComment implements CommentService.
func (csi *CommentServiceImpl) DeleteComment(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	delete := bson.D{{Key: "$set", Value: bson.D{{Key: "is_deleted", Value: true}}}}

	res := csi.commentCollection.FindOneAndUpdate(csi.ctx, query, delete, options.FindOneAndUpdate().SetReturnDocument(1))
	var deletedComment *models.DBComment
	if err := res.Decode(&deletedComment); err != nil {
		return errors.New("no comment with that Id exists")
	}

	return nil
}

// GetComment implements CommentService.
func (csi *CommentServiceImpl) GetComment(id string) (*models.DBComment, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var comment *models.DBComment

	if err := csi.commentCollection.FindOne(csi.ctx, query).Decode(&comment); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no documents with that Id exists")
		}

		return nil, err
	}

	return comment, nil
}

// GetComments implements CommentService.
func (csi *CommentServiceImpl) GetComments(page int, limit int) ([]*models.DBComment, error) {
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

	cursor, err := csi.commentCollection.Find(csi.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(csi.ctx)

	var comments []*models.DBComment

	for cursor.Next(csi.ctx) {
		comment := &models.DBComment{}
		err := cursor.Decode(comment)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return []*models.DBComment{}, nil
	}

	return comments, nil
}

// UpdateComment implements CommentService.
func (csi *CommentServiceImpl) UpdateComment(id string, data *models.UpdateComment) (*models.DBComment, error) {

	doc, err := utils.ToDoc(data)

	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}

	res := csi.commentCollection.FindOneAndUpdate(csi.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedComment *models.DBComment

	if err := res.Decode(&updatedComment); err != nil {
		return nil, errors.New("no comment with that Id exists")
	}

	return updatedComment, nil
}

func (csi *CommentServiceImpl) GetCommentByQuestionId(id string) []*models.DBComment {
	opt := options.FindOptions{}
	query := bson.D{{Key: "question_id", Value: id}, {Key: "is_deleted", Value: false}}

	cursor, err := csi.commentCollection.Find(csi.ctx, query, &opt)
	if err != nil {
		return nil
	}

	defer cursor.Close(csi.ctx)

	var comments []*models.DBComment

	for cursor.Next(csi.ctx) {
		comment := &models.DBComment{}
		err := cursor.Decode(comment)

		if err != nil {
			return nil
		}

		comments = append(comments, comment)
	}

	if err := cursor.Err(); err != nil {
		return nil
	}

	if len(comments) == 0 {
		return nil
	}

	return comments
}

func (csi *CommentServiceImpl) GetCommentsByUserId(id string, page int, limit int) ([]*models.DBComment, error) {

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

	cursor, err := csi.commentCollection.Find(csi.ctx, query, &opt)
	var comments []*models.DBComment

	if err != nil {
		return nil, err
	}

	defer cursor.Close(csi.ctx)

	for cursor.Next(csi.ctx) {
		bookmark := &models.DBComment{}
		err := cursor.Decode(bookmark)
		if err != nil {
			return nil, err
		}

		comments = append(comments, bookmark)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return []*models.DBComment{}, nil
	}

	return comments, err
}

func (csi *CommentServiceImpl) GetAnswersByUserId(id string, page int, limit int) ([]*models.DBComment, error) {

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

	query := bson.D{
		{Key: "user_id", Value: id},
		{Key: "$or", Value: bson.A{
			bson.D{
				{Key: "sort", Value: bson.D{{Key: "$ne", Value: 1}}},
			},
			bson.D{
				{Key: "sort", Value: 1},
				{Key: "parent_id", Value: bson.D{{Key: "$exists", Value: true}}},
			},
		}},
	}

	cursor, err := csi.commentCollection.Find(csi.ctx, query, &opt)
	var comments []*models.DBComment

	if err != nil {
		return nil, err
	}

	defer cursor.Close(csi.ctx)

	for cursor.Next(csi.ctx) {
		bookmark := &models.DBComment{}
		err := cursor.Decode(bookmark)
		if err != nil {
			return nil, err
		}

		comments = append(comments, bookmark)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return []*models.DBComment{}, nil
	}

	return comments, err
}

func (csi *CommentServiceImpl) GetCommentsByUserIdWithSolved(id string, page int, limit int) ([]*models.DBComment, error) {

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

	query := bson.D{{Key: "user_id", Value: id}, {Key: "is_solution", Value: true}}

	cursor, err := csi.commentCollection.Find(csi.ctx, query, &opt)
	var comments []*models.DBComment

	if err != nil {
		return nil, err
	}

	defer cursor.Close(csi.ctx)

	for cursor.Next(csi.ctx) {
		bookmark := &models.DBComment{}
		err := cursor.Decode(bookmark)
		if err != nil {
			return nil, err
		}

		comments = append(comments, bookmark)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return []*models.DBComment{}, nil
	}

	return comments, err
}

func (csi *CommentServiceImpl) GetCommentCount(id string) (count int64) {
	query := bson.D{
		{Key: "user_id", Value: id},
		{Key: "$or", Value: bson.A{
			bson.D{
				{Key: "sort", Value: bson.D{{Key: "$ne", Value: 1}}},
			},
			bson.D{
				{Key: "sort", Value: 1},
				{Key: "parent_id", Value: bson.D{{Key: "$exists", Value: true}}},
			},
		}},
	}

	res, err := csi.commentCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}

func (csi *CommentServiceImpl) GetCommentCountBySolved(id string) (count int64) {
	query := bson.D{{Key: "user_id", Value: id}, {Key: "is_solution", Value: true}}

	res, err := csi.commentCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}

func (csi *CommentServiceImpl) GetCommentCountByQuestionIdSolved(id string) (count int64) {
	query := bson.D{{Key: "question_id", Value: id}, {Key: "is_solution", Value: true}}

	res, err := csi.commentCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}

func NewCommentService(commentCollection *mongo.Collection, ctx context.Context) CommentService {
	return &CommentServiceImpl{commentCollection, ctx}
}
