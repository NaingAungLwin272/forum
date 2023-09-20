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

type TeamServiceImpl struct {
	teamCollection *mongo.Collection
	ctx            context.Context
}

func NewTeamService(TeamCollection *mongo.Collection, ctx context.Context) TeamService {
	return &TeamServiceImpl{TeamCollection, ctx}
}

func (p *TeamServiceImpl) CreateTeam(team *models.CreateTeamRequest) (*models.DBTeam, error) {
	team.CreateAt = time.Now()
	team.UpdatedAt = team.CreateAt

	res, err := p.teamCollection.InsertOne(p.ctx, team)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New(consts.TeamExists)
		}
		return nil, err
	}

	var newTeam *models.DBTeam
	query := bson.M{"_id": res.InsertedID}
	if err = p.teamCollection.FindOne(p.ctx, query).Decode(&newTeam); err != nil {
		return nil, err
	}
	return newTeam, nil
}

func (p *TeamServiceImpl) GetTeam(id string) (*models.DBTeam, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var Team *models.DBTeam

	if err := p.teamCollection.FindOne(p.ctx, query).Decode(&Team); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(consts.TeamNotFound)
		}

		return nil, err
	}

	return Team, nil
}

func (p *TeamServiceImpl) GetTeamList(page int, limit int) []*models.DBTeam {
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

	cursor, err := p.teamCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil
	}

	defer cursor.Close(p.ctx)

	var teams []*models.DBTeam

	for cursor.Next(p.ctx) {
		team := &models.DBTeam{}
		err := cursor.Decode(team)

		if err != nil {
			return []*models.DBTeam{}
		}

		teams = append(teams, team)
	}

	if err := cursor.Err(); err != nil {
		return []*models.DBTeam{}
	}

	if len(teams) == 0 {
		return []*models.DBTeam{}
	}

	return teams
}

func (p *TeamServiceImpl) UpdateTeam(id string, data *models.UpdateTeam) (*models.DBTeam, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.teamCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedTeam *models.DBTeam
	if err := res.Decode(&updatedTeam); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(consts.TeamNotFound)
		} else {
			return nil, errors.New(consts.TeamExists)
		}
	}

	return updatedTeam, nil
}

func (p *TeamServiceImpl) DeleteTeam(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.teamCollection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New(consts.TeamNotFound)
	}

	return nil
}

func (p *TeamServiceImpl) GetTeamByDeparmentId(id string) ([]*models.DBTeam, error) {

	query := bson.D{{Key: "deparment_id", Value: id}}

	var teams []*models.DBTeam

	cursor, err := p.teamCollection.Find(p.ctx, query)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	for cursor.Next(p.ctx) {
		var team *models.DBTeam
		if err := cursor.Decode(&team); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(teams) == 0 {
		return []*models.DBTeam{}, nil
	}

	return teams, err
}

func (p *TeamServiceImpl) GetTeamCount(page int, limit int) int64 {
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

	res, err := p.teamCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}
