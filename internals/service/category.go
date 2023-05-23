package service

import (
	"context"

	"github.com/gabrielDpadua21/fc-grpc-essentialst/internals/database"
	"github.com/gabrielDpadua21/fc-grpc-essentialst/internals/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, input *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(input.Name, input.Description)
	if err != nil {
		return nil, err
	}

	newCategory := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	categoryResponse := &pb.CategoryResponse{
		Category: newCategory,
	}

	return categoryResponse, nil
}
