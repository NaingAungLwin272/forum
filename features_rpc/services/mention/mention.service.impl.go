package mention_service

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

type MentionServiceImpl struct {
	mentionCollection *mongo.Collection
	ctx               context.Context
}

// CreateMention implements MentionService.
func (msi *MentionServiceImpl) CreateMention(msi_request *models.CreateMentionRequest) (*models.DBMention, error) {
	// Check if a mention with the same user_id, question_id, and comment_id exists
	existingMention, err := msi.findExistingMention(msi_request.User_Id, msi_request.Question_Id, msi_request.Comment_Id)
	if err != nil {
		return nil, err
	}

	// If an existing mention is found, return it instead of creating a new one
	if existingMention != nil {
		return existingMention, nil
	}

	// No existing mention found, proceed to create the new mention
	res, err := msi.mentionCollection.InsertOne(msi.ctx, msi_request)
	if err != nil {
		return nil, err
	}

	// Retrieve the newly created mention from the database
	var newMention *models.DBMention
	query := bson.M{"_id": res.InsertedID}
	if err = msi.mentionCollection.FindOne(msi.ctx, query).Decode(&newMention); err != nil {
		return nil, err
	}

	return newMention, nil
}

// Helper function to find an existing mention with the same user_id, question_id, and comment_id
func (msi *MentionServiceImpl) findExistingMention(user_id, question_id, comment_id string) (*models.DBMention, error) {
	query := bson.M{
		"user_id":     user_id,
		"question_id": question_id,
		"comment_id":  comment_id,
	}

	var existingMention *models.DBMention
	err := msi.mentionCollection.FindOne(msi.ctx, query).Decode(&existingMention)

	if err != nil && err != mongo.ErrNoDocuments {
		// Error occurred during the query (excluding "not found" errors)
		return nil, err
	}

	if existingMention != nil {
		// Found an existing mention with the same user_id, question_id, and comment_id
		return existingMention, nil
	}

	// No existing mention found
	return nil, nil
}

// DeleteMention implements MentionService.
func (msi *MentionServiceImpl) DeleteMention(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := msi.mentionCollection.DeleteOne(msi.ctx, query)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with the Id exists ...msi delete")
	}

	return nil
}

// GetMention implements MentionService.
func (msi *MentionServiceImpl) GetMention(id string) (*models.DBMention, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var mention *models.DBMention

	if err := msi.mentionCollection.FindOne(msi.ctx, query).Decode(&mention); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no documents with that Id exists")
		}

		return nil, err
	}

	return mention, nil
}

// GetMentions implements MentionService.
func (msi *MentionServiceImpl) GetMentions(page int, limit int) ([]*models.DBMention, error) {
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

	cursor, err := msi.mentionCollection.Find(msi.ctx, query, &opt)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(msi.ctx)

	var mentions []*models.DBMention

	for cursor.Next(msi.ctx) {
		mention := &models.DBMention{}
		err := cursor.Decode(mention)

		if err != nil {
			return nil, err
		}

		mentions = append(mentions, mention)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(mentions) == 0 {
		return []*models.DBMention{}, nil
	}

	return mentions, nil
}

// UpdateMention implements MentionService.
func (msi *MentionServiceImpl) UpdateMention(id string, data *models.UpdateMention) (*models.DBMention, error) {
	doc, err := utils.ToDoc(data)

	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}

	res := msi.mentionCollection.FindOneAndUpdate(msi.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedMention *models.DBMention

	if err := res.Decode(&updatedMention); err != nil {
		return nil, errors.New("no mention with that Id exsts...")
	}

	return updatedMention, nil
}

func (msi *MentionServiceImpl) GetMentionsByUserId(id string, page int, limit int) []*models.DBMention {

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
	var mentions []*models.DBMention
	cursor, err := msi.mentionCollection.Find(msi.ctx, query, &opt)

	if err != nil {
		return nil
	}

	defer cursor.Close(msi.ctx)

	for cursor.Next(msi.ctx) {
		mention := &models.DBMention{}
		err := cursor.Decode(mention)
		if err != nil {
			return nil
		}

		mentions = append(mentions, mention)
	}

	if err := cursor.Err(); err != nil {
		return nil
	}

	if len(mentions) == 0 {
		return []*models.DBMention{}
	}

	return mentions
}

func (msi *MentionServiceImpl) GetMentionCount(id string) (count int64) {
	query := bson.D{{Key: "user_id", Value: id}}

	res, err := msi.mentionCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}

func NewMentionService(mentionCollection *mongo.Collection, ctx context.Context) MentionService {
	return &MentionServiceImpl{mentionCollection, ctx}
}
