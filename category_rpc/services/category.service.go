package services

import models "github.com/scm-dev1dev5/mtm-community-forum/category_rpc/models"

type CategoryService interface {
	CreateCategory(*models.CreateCategoryRequest) (*models.DBCategory, error)
	GetCategory(string) (*models.DBCategory, error)
	GetCategories(page int, limit int) ([]*models.DBCategory, error)
	UpdateCategory(string, *models.UpdateCategory) (*models.DBCategory, error)
	DeleteCategory(string) error
	GetCategoryByType(page int64, limit int64, categoryType int32) ([]*models.DBCategory, error)
	GetCategoryCount(page int, limit int) (languageCount int64,tagCount int64)
}
