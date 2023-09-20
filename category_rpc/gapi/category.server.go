package gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/category_rpc/pb"
	services "github.com/scm-dev1dev5/mtm-community-forum/category_rpc/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryServer struct {
	pb.UnimplementedCategoryServiceServer
	categoryCollection *mongo.Collection
	categoryService    services.CategoryService
}

func NewGrpcCategoryServer(categoryCollection *mongo.Collection, categoryService services.CategoryService) (*CategoryServer, error) {
	categoryServer := &CategoryServer{
		categoryCollection: categoryCollection,
		categoryService:    categoryService,
	}

	return categoryServer, nil
}
