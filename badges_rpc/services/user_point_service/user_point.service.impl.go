package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/consts"
	models "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/models"
	badgeService "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/services/badge_service"
	userBadgeService "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/services/user_badges_service"
	"github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserPointServiceImpl struct {
	UserPointCollection *mongo.Collection
	userBadgeCollection *mongo.Collection
	commentCollection   *mongo.Collection
	userCollection      *mongo.Collection
	badgeCollection     *mongo.Collection
	ctx                 context.Context
}

func NewUserPointService(
	UserPointCollection *mongo.Collection,
	userBadgeCollection *mongo.Collection,
	badgeCollection *mongo.Collection,
	commentCollection *mongo.Collection,
	userCollection *mongo.Collection,
	ctx context.Context) UserPointService {
	return &UserPointServiceImpl{
		UserPointCollection,
		userBadgeCollection,
		commentCollection,
		userCollection,
		badgeCollection,
		ctx}
}

func (p *UserPointServiceImpl) CreateUserPoint(UserPoint *models.CreateUserPointRequest) (*models.DBUserPoint, error) {
	UserPoint.CreateAt = time.Now()
	UserPoint.UpdatedAt = UserPoint.CreateAt

	res, err := p.UserPointCollection.InsertOne(p.ctx, UserPoint)

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"user_id": 1}, Options: opt}

	if _, err := p.UserPointCollection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, errors.New(consts.CreateUserPointErr)
	}

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New(consts.UserPointExists)
		}
		return nil, err
	}

	var newUserPoint *models.DBUserPoint
	query := bson.M{"_id": res.InsertedID}
	if err = p.UserPointCollection.FindOne(p.ctx, query).Decode(&newUserPoint); err != nil {
		return nil, err
	}

	return newUserPoint, nil
}

func (p *UserPointServiceImpl) GetUserPoint(userId string) (*models.DBUserPoint, error) {

	query := bson.M{"user_id": userId}

	var UserPoint *models.DBUserPoint

	if err := p.UserPointCollection.FindOne(p.ctx, query).Decode(&UserPoint); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(consts.NotFoundInDB)
		}

		return nil, err
	}

	return UserPoint, nil
}

func (p *UserPointServiceImpl) GetUserPoints(page int, limit int) ([]*models.DBUserPoint, error) {
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

	cursor, err := p.UserPointCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var UserPoints []*models.DBUserPoint

	for cursor.Next(p.ctx) {
		UserPoint := &models.DBUserPoint{}
		err := cursor.Decode(UserPoint)

		if err != nil {
			return nil, err
		}

		UserPoints = append(UserPoints, UserPoint)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(UserPoints) == 0 {
		return []*models.DBUserPoint{}, nil
	}

	return UserPoints, nil
}

func (p *UserPointServiceImpl) UpdateUserPoint(userId string, data *models.UpdateUserPoint) (*models.DBUserPoint, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.UserPointCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedUserPoint *models.DBUserPoint
	if err := res.Decode(&updatedUserPoint); err != nil {
		return nil, errors.New(consts.NotFoundInDB)
	}

	return updatedUserPoint, nil
}

func (p *UserPointServiceImpl) DeleteUserPoint(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.UserPointCollection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New(consts.NotFoundInDB)
	}

	return nil
}

func (p *UserPointServiceImpl) GetUserComments(userId string) ([]*models.DBComment, error) {

	opt := options.FindOptions{}
	opt.SetSort(bson.M{"_id": -1})

	query := bson.M{"user_id": userId}

	cursor, err := p.commentCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var userComments []*models.DBComment

	for cursor.Next(p.ctx) {
		userComment := &models.DBComment{}
		err := cursor.Decode(userComment)

		if err != nil {
			return nil, err
		}

		userComments = append(userComments, userComment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(userComments) == 0 {
		return []*models.DBComment{}, nil
	}

	return userComments, nil
}

func (p *UserPointServiceImpl) DetermineLevel(
	userComments []*models.DBComment,
	UserPoint *models.DBUserPoint,
	badges []*models.DBBadge) (int32, int32) {

	var tmpReactLvl int32
	var tmpQaLvl int32
	var tmpCount uint64
	if len(userComments) > 0 {
		for _, comment := range userComments {
			// get the biggest vote count from the array
			if comment.Vote_Count > tmpCount {
				tmpCount = comment.Vote_Count
				if comment.Vote_Count >= 5 && comment.Vote_Count < 20 {
					tmpReactLvl = consts.ReactLvlList[0]
				} else if comment.Vote_Count >= 20 && comment.Vote_Count < 40 {
					tmpReactLvl = consts.ReactLvlList[1]
				} else if comment.Vote_Count >= 40 && comment.Vote_Count < 60 {
					tmpReactLvl = consts.ReactLvlList[2]
				} else if comment.Vote_Count >= 60 && comment.Vote_Count < 100 {
					tmpReactLvl = consts.ReactLvlList[3]
				} else if comment.Vote_Count >= 150 {
					tmpReactLvl = consts.ReactLvlList[4]
				}
			}
		}
	}

	activityPoint := UserPoint.AnswerCount + UserPoint.QuestionCount
	if UserPoint.SolvedCount >= 10 && activityPoint >= 150 {
		tmpQaLvl = consts.MedalList[4]
	} else if UserPoint.SolvedCount >= 5 {
		if activityPoint >= 60 {
			tmpQaLvl = consts.MedalList[3]
		} else {
			tmpQaLvl = consts.MedalList[2]
		}
	} else if UserPoint.SolvedCount >= 3 {
		if activityPoint >= 40 {
			tmpQaLvl = consts.MedalList[2]
		} else {
			tmpQaLvl = consts.MedalList[1]
		}
	} else if UserPoint.SolvedCount >= 1 {
		if activityPoint >= 20 {
			tmpQaLvl = consts.MedalList[1]
		} else {
			tmpQaLvl = consts.MedalList[0]
		}
	} else if UserPoint.SolvedCount == 0 {
		if activityPoint >= 5 {
			tmpQaLvl = consts.MedalList[0]
		}
	}
	return tmpReactLvl, tmpQaLvl
}

func (p *UserPointServiceImpl) EvaluatePoints() error {
	badgeService := badgeService.NewBadgeService(p.badgeCollection, p.ctx)

	query := bson.M{"is_deleted": false}
	cursor, err := p.userCollection.Find(p.ctx, query)
	if err != nil {
		return nil
	}
	defer cursor.Close(p.ctx)

	var users []*models.DBUser
	for cursor.Next(p.ctx) {
		user := &models.DBUser{}
		err := cursor.Decode(user)
		if err != nil {
			return nil
		}
		users = append(users, user)
	}

	page := 0
	limit := 0
	var badges []*models.DBBadge
	badges, err = badgeService.GetBadges(page, limit)
	if err != nil {
		return nil
	}

	for _, userData := range users {
		tmpQaBadgeId := []string{}
		tmpReactBadgeId := []string{}
		exampleUserPoint, _ := p.GetUserPoint(userData.Id.Hex())

		exampleUserComment, _ := p.GetUserComments(userData.Id.Hex())

		tmpReactLvl, tmpQaLvl := p.DetermineLevel(exampleUserComment, exampleUserPoint, badges)

		if tmpQaLvl != 0 || tmpReactLvl != 0 {
			for _, item := range badges {
				if item.Type == 1 && item.Level > exampleUserPoint.QaLevel && item.Level <= tmpQaLvl {
					fmt.Println(item, "item", userData.Id, "id Type 1")
					tmpQaBadgeId = append(tmpQaBadgeId, item.Id.Hex())
				}
				if item.Type == 2 && item.Level > exampleUserPoint.ReactionLevel && item.Level <= tmpReactLvl {
					fmt.Println(item, "item", userData.Id, "id Type 2")
					tmpReactBadgeId = append(tmpReactBadgeId, item.Id.Hex())
				}
			}
		}

		if len(tmpQaBadgeId) > 0 {
			for _, qaBadgeId := range tmpQaBadgeId {
				if err := p.setUserGadge(userData.Id.Hex(), qaBadgeId); err != nil {
					return nil
				}
			}
		}
		if len(tmpReactBadgeId) > 0 {
			for _, reactBadgeId := range tmpReactBadgeId {
				if err := p.setUserGadge(userData.Id.Hex(), reactBadgeId); err != nil {
					return nil
				}
			}
		}

		pointData := &models.UpdateUserPoint{
			ReactionLevel: tmpReactLvl,
			QaLevel:       tmpQaLvl,
			QuestionCount: exampleUserPoint.QuestionCount,
			AnswerCount:   exampleUserPoint.AnswerCount,
			SolvedCount:   exampleUserPoint.SolvedCount,
			UpdatedAt:     time.Now(),
		}

		// update new UserPoint with new level id
		p.UpdateUserPoint(userData.Id.Hex(), pointData)

	}
	return nil
}

func (p *UserPointServiceImpl) setUserGadge(userId string, badgeId string) error {
	userBadgeService := userBadgeService.NewUserBadgeService(p.userBadgeCollection, p.ctx)
	userBadge, _ := userBadgeService.GetUserBadge(userId, badgeId)
	if userBadge == nil {
		data := &models.CreateUserBadgeRequest{
			User_Id:  userId,
			Badge_Id: badgeId,
		}
		// if userBadge is not exist in db create new userBadge
		userBadgeService.CreateUserBadge(data)
	}
	return nil
}
