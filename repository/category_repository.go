package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go_redis_1/model/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *gorm.DB, category domain.Category) domain.Category
	FindById(ctx context.Context, tx *gorm.DB, rdb *redis.Client, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *gorm.DB, rdb *redis.Client) []domain.Category
}
