package vote_service

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

type VoteServiceImpl struct {
	voteCollection *mongo.Collection
	ctx            context.Context
}

// CreateVote implements VoteService.
func (vsi *VoteServiceImpl) CreateVote(vsi_request *models.CreateVoteRequest) (*models.DBVote, error) {
	res, err := vsi.voteCollection.InsertOne(vsi.ctx, vsi_request)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("Cannot create vote")
		}
		return nil, err
	}

	var newVote *models.DBVote
	query := bson.M{"_id": res.InsertedID}
	if err = vsi.voteCollection.FindOne(vsi.ctx, query).Decode(&newVote); err != nil {
		return nil, err
	}

	return newVote, nil
}

// DeleteVote implements VoteService.
func (vsi *VoteServiceImpl) DeleteVote(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := vsi.voteCollection.DeleteOne(vsi.ctx, query)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with the Id exists ...vsi delete")
	}

	return nil
}

// GetVote implements VoteService.
func (vsi *VoteServiceImpl) GetVote(id string) (*models.DBVote, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var vote *models.DBVote

	if err := vsi.voteCollection.FindOne(vsi.ctx, query).Decode(&vote); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no documents with that Id exists")
		}

		return nil, err
	}

	return vote, nil
}

// GetVotes implements VoteService.
func (vsi *VoteServiceImpl) GetVotes(page int, limit int) ([]*models.DBVote, error) {
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

	cursor, err := vsi.voteCollection.Find(vsi.ctx, query, &opt)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(vsi.ctx)

	var votes []*models.DBVote

	for cursor.Next(vsi.ctx) {
		vote := &models.DBVote{}
		err := cursor.Decode(vote)

		if err != nil {
			return nil, err
		}

		votes = append(votes, vote)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(votes) == 0 {
		return []*models.DBVote{}, nil
	}

	return votes, nil
}

// UpdateVote implements VoteService.
func (vsi *VoteServiceImpl) UpdateVote(id string, data *models.UpdateVote) (*models.DBVote, error) {
	doc, err := utils.ToDoc(data)

	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}

	res := vsi.voteCollection.FindOneAndUpdate(vsi.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedVote *models.DBVote

	if err := res.Decode(&updatedVote); err != nil {
		return nil, errors.New("no vote with that Id exsts...")
	}

	return updatedVote, nil
}

func (vsi *VoteServiceImpl) GetVotesByUserId(id string, page int, limit int) ([]*models.DBVote, error) {

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

	cursor, err := vsi.voteCollection.Find(vsi.ctx, query, &opt)
	var votes []*models.DBVote
	if err != nil {
		return nil, err
	}

	defer cursor.Close(vsi.ctx)

	for cursor.Next(vsi.ctx) {
		vote := &models.DBVote{}
		err := cursor.Decode(vote)
		if err != nil {
			return nil, err
		}

		votes = append(votes, vote)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(votes) == 0 {
		return []*models.DBVote{}, nil
	}

	return votes, err
}

func (vsi *VoteServiceImpl) GetVoteCount(id string) (count int64) {
	query := bson.D{{Key: "user_id", Value: id}}

	res, err := vsi.voteCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}

func (vsi *VoteServiceImpl) GetVotesByUserIdQuestionId(id string, question_id string, page int, limit int) ([]*models.DBVote, error) {
	query := bson.D{{Key: "user_id", Value: id}, {Key: "question_id", Value: question_id}}

	cursor, err := vsi.voteCollection.Find(vsi.ctx, query)
	var votes []*models.DBVote

	if err != nil {
		return nil, err
	}

	defer cursor.Close(vsi.ctx)

	for cursor.Next(vsi.ctx) {
		vote := &models.DBVote{}
		err := cursor.Decode(vote)
		if err != nil {
			return nil, err
		}

		votes = append(votes, vote)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(votes) == 0 {
		return []*models.DBVote{}, nil
	}

	return votes, err
}

func NewVoteService(voteCollection *mongo.Collection, ctx context.Context) VoteService {
	return &VoteServiceImpl{voteCollection, ctx}
}
