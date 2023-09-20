package gapi

import (
	"context"
	"strings"
	"time"

	"github.com/scm-dev1dev5/mtm-community-forum/category_rpc/consts"
	models "github.com/scm-dev1dev5/mtm-community-forum/category_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/category_rpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (categoryServer *CategoryServer) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category := &models.CreateCategoryRequest{
		Type: req.GetType(),
		Name: req.GetName(),
	}

	newCategory, err := categoryServer.categoryService.CreateCategory(category)

	if err != nil {
		if strings.Contains(err.Error(), "category") {
			return nil, status.Errorf(codes.AlreadyExists, consts.CategoryExists)
		}
		return nil, status.Errorf(codes.Internal, consts.CategoryExists)
	}

	res := &pb.CategoryResponse{
		XId:       newCategory.Id.Hex(),
		Type:      newCategory.Type,
		Name:      newCategory.Name,
		CreatedAt: timestamppb.New(newCategory.CreateAt),
		UpdatedAt: timestamppb.New(newCategory.UpdatedAt),
	}
	return res, nil
}

func (categoryServer *CategoryServer) GetCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.CategoryResponse, error) {
	categorytId := req.GetId()

	category, err := categoryServer.categoryService.GetCategory(categorytId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())

		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CategoryResponse{
		XId:       category.Id.Hex(),
		Type:      category.Type,
		Name:      category.Name,
		CreatedAt: timestamppb.New(category.CreateAt),
		UpdatedAt: timestamppb.New(category.UpdatedAt),
	}
	return res, nil
}

func (categoryServer *CategoryServer) GetCategories(ctx context.Context, req *pb.GetCategoriesRequest) (*pb.CategoryResponseList, error) {
	var page = req.GetPage()
	var limit = req.GetLimit()

	categories, err := categoryServer.categoryService.GetCategories(int(page), int(limit))
	if err != nil {
		return nil, err
	}

	cateList := make([]*pb.Category, 0, len(categories))
	for _, category := range categories {
		response := &pb.Category{
			XId:       category.Id.Hex(),
			Type:      category.Type,
			Name:      category.Name,
			CreatedAt: timestamppb.New(category.CreateAt),
			UpdatedAt: timestamppb.New(category.UpdatedAt),
		}
		cateList = append(cateList, response)
	}

	response := &pb.CategoryResponseList{
		Categories: cateList,
	}

	return response, nil
}

func (categoryServer *CategoryServer) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.CategoryResponse, error) {
	categoryId := req.GetXId()

	category := &models.UpdateCategory{
		Type:      req.GetType(),
		Name:      req.GetName(),
		UpdatedAt: time.Now(),
	}

	updatedCategory, err := categoryServer.categoryService.UpdateCategory(categoryId, category)

	if err != nil {
		if strings.Contains(err.Error(), "categories") {
			return nil, status.Errorf(codes.AlreadyExists, consts.CategoryExists)
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CategoryResponse{
		XId:       updatedCategory.Id.Hex(),
		Type:      updatedCategory.Type,
		Name:      updatedCategory.Name,
		CreatedAt: timestamppb.New(updatedCategory.CreateAt),
		UpdatedAt: timestamppb.New(updatedCategory.UpdatedAt),
	}
	return res, nil
}

func (categoryServer *CategoryServer) DeleteCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.DeleteCategoryResponse, error) {
	categoryId := req.GetId()

	err := categoryServer.categoryService.DeleteCategory(categoryId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteCategoryResponse{
		Success: true,
	}

	return res, nil
}

func (categoryServer *CategoryServer) GetCategoryByType(ctx context.Context, req *pb.GetCategoryByTypeRequest) (*pb.CategoryResponseList, error) {
	categoryType := req.GetType()
	var page = req.GetPage()
	var limit = req.GetLimit()

	categories, err := categoryServer.categoryService.GetCategoryByType(page, limit, categoryType)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	categoryList := make([]*pb.Category, 0, len(categories))
	for _, category := range categories {
		categoryPB := &pb.Category{
			XId:       category.Id.Hex(),
			Type:      category.Type,
			Name:      category.Name,
			CreatedAt: timestamppb.New(category.CreateAt),
			UpdatedAt: timestamppb.New(category.UpdatedAt),
		}
		categoryList = append(categoryList, categoryPB)
	}

	response := &pb.CategoryResponseList{
		Categories: categoryList,
	}

	return response, nil
}

func (categoryServer *CategoryServer) GetCategoryCount(context context.Context, req *pb.GetCategoriesRequest) (*pb.CategoryCountResponse, error) {
	var page = req.GetPage()
	var limit = req.GetLimit()
	language, tag := categoryServer.categoryService.GetCategoryCount(int(page), int(limit))
	res := &pb.CategoryCountResponse{
		LanguageCount: int64(language),
		TagCount:      int64(tag),
	}
	return res, nil
}
