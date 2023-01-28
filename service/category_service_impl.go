package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"go_redis_1/exception"
	"go_redis_1/helper"
	"go_redis_1/model/domain"
	"go_redis_1/model/web"
	"go_redis_1/repository"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *gorm.DB
	RDB                *redis.Client
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *gorm.DB, rdb *redis.Client, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		RDB:                rdb,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx := service.DB
	//tx, err := service.DB.Begin()
	//helper.PanicIfError(err)
	//defer helper.CommitOrRollback(tx)

	rdb := service.RDB

	category, err := service.CategoryRepository.FindById(ctx, tx, rdb, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx := service.DB
	//tx, err := service.DB.Begin()
	//helper.PanicIfError(err)
	//defer helper.CommitOrRollback(tx)

	rdb := service.RDB

	categories := service.CategoryRepository.FindAll(ctx, tx, rdb)

	return helper.ToCategoryResponses(categories)
}
