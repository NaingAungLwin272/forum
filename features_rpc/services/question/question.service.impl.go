package question_service

import (
	"context"
	"errors"
	"time"

	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QuestionServiceImpl struct {
	questionCollection *mongo.Collection
	ctx                context.Context
}

// DeleteQuestion implements QuestionService.
func (qsi *QuestionServiceImpl) DeleteQuestion(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := qsi.questionCollection.DeleteOne(qsi.ctx, query)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with the Id exists...qsi delete")
	}

	return nil
}

// GetQuestion implements QuestionService.
func (qsi *QuestionServiceImpl) GetQuestion(id string) (*models.DBQuestion, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var question *models.DBQuestion

	if err := qsi.questionCollection.FindOne(qsi.ctx, query).Decode(&question); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("question not found")
		}

		return nil, err
	}

	return question, nil
}

// UpdateQuestion implements QuestionService.
func (qsi *QuestionServiceImpl) UpdateQuestion(id string, data *models.UpdateQuestion) (*models.DBQuestion, error) {
	doc, err := utils.ToDoc(data)

	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}

	res := qsi.questionCollection.FindOneAndUpdate(qsi.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedQuestion *models.DBQuestion
	if err := res.Decode(&updatedQuestion); err != nil {
		return nil, errors.New("no question with the Id exists...")
	}

	return updatedQuestion, nil
}

// CreateQuestion implements QuestionService.
func (qs *QuestionServiceImpl) CreateQuestion(question *models.CreateQuestionRequest) (*models.DBQuestion, error) {
	question.CreateAt = time.Now()
	question.UpdatedAt = question.CreateAt
	question.ViewCount = 0
	question.VoteCount = 0
	question.ReplyCount = 0
	question.SolutionCount = 0
	// if (!question.IsDeleted) {
	// 	question.IsDeleted = false
	// }

	res, err := qs.questionCollection.InsertOne(qs.ctx, question)
	/*
		if err != nil {
		   		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
		   			return nil, errors.New("user with that title already exists")
		   		}
		   		return nil, err
		   	}

		   	opt := options.Index()
		   	opt.SetUnique(true)

		   	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

		   	if _, err := p.userCollection.Indexes().CreateOne(p.ctx, index); err != nil {
		   		return nil, errors.New("could not create index for email")
		   	}
	*/

	var newQuestion *models.DBQuestion

	query := bson.M{"_id": res.InsertedID}
	if err = qs.questionCollection.FindOne(qs.ctx, query).Decode(&newQuestion); err != nil {
		return nil, err
	}

	return newQuestion, nil
}

// GetQuestions implements QuestionService.
func (p *QuestionServiceImpl) GetQuestions(page int, limit int, order string, sort string) ([]*models.DBQuestion, error) {
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

	if order != "" {
		if sort != "ascend" {
			switch order {
			case "vote":
				opt.SetSort(bson.M{"vote_count": 1})
			case "view":
				opt.SetSort(bson.M{"view_count": 1})
			case "reply":
				opt.SetSort(bson.M{"reply_count": 1})
			case "solution":
				opt.SetSort(bson.M{"solution_count": 1})
			}

		}
		if sort != "descend" {
			switch order {
			case "vote":
				opt.SetSort(bson.M{"vote_count": -1})
			case "view":
				opt.SetSort(bson.M{"view_count": -1})
			case "reply":
				opt.SetSort(bson.M{"reply_count": -1})
			case "solution":
				opt.SetSort(bson.M{"solution_count": -1})
			}
		}
		if sort == "undefined" || sort == "null" {
			opt.SetSort(bson.M{"created_at": -1})
		}
	}

	query := bson.M{"is_deleted": false}
	cursor, err := p.questionCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var questions []*models.DBQuestion

	for cursor.Next(p.ctx) {
		question := &models.DBQuestion{}
		err := cursor.Decode(question)

		if err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(questions) == 0 {
		return []*models.DBQuestion{}, nil
	}

	return questions, nil
}

func (p *QuestionServiceImpl) GetQuestionCount(id string) (count int64) {
	query := bson.D{{Key: "user_id", Value: id}, {Key: "is_deleted", Value: false}}

	res, err := p.questionCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}

func (p *QuestionServiceImpl) GetQuestionsByUserId(id string, page int, limit int, order string, sort string) []*models.DBQuestion {
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

	if order != "" {
		if sort != "ascend" {
			switch order {
			case "vote":
				opt.SetSort(bson.M{"vote_count": 1})
			case "view":
				opt.SetSort(bson.M{"view_count": 1})
			case "reply":
				opt.SetSort(bson.M{"reply_count": 1})
			case "solution":
				opt.SetSort(bson.M{"solution_count": 1})
			}

		}
		if sort != "descend" {
			switch order {
			case "vote":
				opt.SetSort(bson.M{"vote_count": -1})
			case "view":
				opt.SetSort(bson.M{"view_count": -1})
			case "reply":
				opt.SetSort(bson.M{"reply_count": -1})
			case "solution":
				opt.SetSort(bson.M{"solution_count": -1})
			}
		}
		if sort == "undefined" || sort == "null" {
			opt.SetSort(bson.M{"created_at": -1})
		}
	}

	query := bson.D{{Key: "user_id", Value: id}, {Key: "is_deleted", Value: false}}

	cursor, err := p.questionCollection.Find(p.ctx, query, &opt)
	var questions []*models.DBQuestion

	if err != nil {
		return []*models.DBQuestion{}
	}

	defer cursor.Close(p.ctx)

	for cursor.Next(p.ctx) {
		question := &models.DBQuestion{}
		err := cursor.Decode(question)
		if err != nil {
			return nil
		}

		questions = append(questions, question)
	}

	if err := cursor.Err(); err != nil {
		return []*models.DBQuestion{}
	}

	return questions
}

func (p *QuestionServiceImpl) FilterQuestion(question *models.FilterQuestionRequest) ([]*models.DBQuestion, error) {
	if question.Page == 0 {
		question.Page = 0
	}

	if question.Limit == 0 {
		question.Limit = 0
	}

	skip := (question.Page - 1) * question.Limit

	opt := options.FindOptions{}
	if question.Limit > 0 {
		opt.SetLimit(int64(question.Limit))
		opt.SetSkip(int64(skip))
		opt.SetSort(bson.M{"created_at": -1})
	}

	if question.Order != "" {
		if question.Sort != "ascend" {
			switch question.Order {
			case "vote":
				opt.SetSort(bson.M{"vote_count": 1})
			case "view":
				opt.SetSort(bson.M{"view_count": 1})
			case "reply":
				opt.SetSort(bson.M{"reply_count": 1})
			case "solution":
				opt.SetSort(bson.M{"solution_count": 1})
			}
		}
		if question.Sort != "descend" {
			switch question.Order {
			case "vote":
				opt.SetSort(bson.M{"vote_count": -1})
			case "view":
				opt.SetSort(bson.M{"view_count": -1})
			case "reply":
				opt.SetSort(bson.M{"reply_count": -1})
			case "solution":
				opt.SetSort(bson.M{"solution_count": -1})
			}
		}
		if question.Sort == "undefined" || question.Sort == "null" {
			opt.SetSort(bson.M{"created_at": -1})
		}
	}

	if question.LanguageIds == nil && question.TagIds == nil && question.UserId == nil && question.Title == "" {
		query := bson.M{"is_deleted": false}
		var allQuestions []*models.DBQuestion
		cursor, err := p.questionCollection.Find(p.ctx, query, &opt)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(p.ctx)

		for cursor.Next(p.ctx) {
			question := &models.DBQuestion{}
			err := cursor.Decode(question)
			if err != nil {
				return nil, err
			}
			allQuestions = append(allQuestions, question)
		}

		if err := cursor.Err(); err != nil {
			return nil, err
		}

		return allQuestions, nil
	}
	filter := bson.D{}

	query := []bson.M{}

	if question.LanguageIds != nil {
		query = append(query, bson.M{"language_ids": bson.M{"$in": question.LanguageIds}})
	}

	if question.TagIds != nil {
		query = append(query, bson.M{"tag_ids": bson.M{"$in": question.TagIds}})
	}

	if question.UserId != nil {
		query = append(query, bson.M{"user_id": bson.M{"$in": question.UserId}})
	}

	if question.Title != "" {
		query = append(query, bson.M{"title": bson.M{"$regex": question.Title, "$options": "i"}})
	}

	filter = append(filter, bson.E{Key: "$and", Value: query})

	var questions []*models.DBQuestion
	cursor, err := p.questionCollection.Find(p.ctx, filter, &opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx)

	for cursor.Next(p.ctx) {
		question := &models.DBQuestion{}
		err := cursor.Decode(question)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(questions) == 0 {
		return []*models.DBQuestion{}, nil
	}

	return questions, nil
}

func (p *QuestionServiceImpl) GetQuestionCountAll() (count int64) {

	query := bson.D{{Key: "is_deleted", Value: false}}

	res, err := p.questionCollection.CountDocuments(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	return res
}

func (p *QuestionServiceImpl) GetFilteredQuestionCount(question *models.FilterQuestionRequest) (count int64) {
	countFilter := bson.D{}

	if question.LanguageIds == nil && question.TagIds == nil && question.UserId == nil && question.Title == "" {
		countFilter = append(countFilter, bson.E{Key: "is_deleted", Value: false})
	} else {
		countQuery := []bson.M{}

		if question.LanguageIds != nil {
			countQuery = append(countQuery, bson.M{"language_ids": bson.M{"$in": question.LanguageIds}})
		}

		if question.TagIds != nil {
			countQuery = append(countQuery, bson.M{"tag_ids": bson.M{"$in": question.TagIds}})
		}

		if question.UserId != nil {
			countQuery = append(countQuery, bson.M{"user_id": bson.M{"$in": question.UserId}})
		}

		if question.Title != "" {
			countQuery = append(countQuery, bson.M{"title": bson.M{"$regex": question.Title, "$options": "i"}})
		}

		countFilter = append(countFilter, bson.E{Key: "$and", Value: countQuery})
	}

	totalDocumentCount, countErr := p.questionCollection.CountDocuments(p.ctx, countFilter)
	if countErr != nil {
		panic(countErr)
	}

	return totalDocumentCount
}
func NewQuestionService(questionCollection *mongo.Collection, ctx context.Context) QuestionService {
	return &QuestionServiceImpl{questionCollection, ctx}
}
