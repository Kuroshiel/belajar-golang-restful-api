package service

import (
	"belajar-golang-restful-api/model/web"
	"context"
)

// Category Service Golang RESTful API

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindByAll(ctx context.Context) []web.CategoryResponse
}
