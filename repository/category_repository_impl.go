package repository

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go_redis_1/helper"
	"go_redis_1/model/domain"
	"gorm.io/gorm"
	"strconv"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, category domain.Category) domain.Category {
	//SQL := "insert into category(name) values (?)"
	//result, err := tx.ExecContext(ctx, SQL, category.Name)

	//Category{}
	user := domain.Category{Name: category.Name}
	tx.Create(&user)

	//helper.PanicIfError(err)

	id := user.Id
	//helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, rdb *redis.Client, categoryId int) (domain.Category, error) {
	/*
	   Redis Pipelines Example
	*/
	//var category domain.Category
	//
	//pipe := rdb.Pipeline()
	//
	//result1 := pipe.Get(ctx, "category_name_"+strconv.Itoa(categoryId))
	//_ = pipe.Set(ctx, "category_name_"+strconv.Itoa(categoryId), `{"Id":47,"name":"Jakarta"}`, 0).Err()
	//result2 := pipe.Get(ctx, "category_name_"+strconv.Itoa(categoryId))
	//
	//_, err := pipe.Exec(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// The value is available only after Exec is called.
	//fmt.Println(result1.Val())
	//fmt.Println(result2.Val())
	//
	//return category, nil

	/*
	   Redis Transactions using TxPipeline
	*/
	//var category domain.Category
	//
	//pipe := rdb.TxPipeline()
	//
	//result1 := pipe.Get(ctx, "category_name_"+strconv.Itoa(categoryId))
	//_ = pipe.Set(ctx, "category_name_"+strconv.Itoa(categoryId), `{"Id":46,"name":"Ijul"}`, 0).Err()
	//result2 := pipe.Get(ctx, "category_name_"+strconv.Itoa(categoryId))
	//
	//_, err := pipe.Exec(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// The value is available only after Exec is called.
	//fmt.Println(result1.Val())
	//fmt.Println(result2.Val())
	//
	//return category, nil

	/*
		   Redis Get & Set By ID
			Without TTL
	*/
	var category domain.Category

	resultRedis, errGetRedis := rdb.Get(ctx, "category_name_"+strconv.Itoa(categoryId)).Result()

	if errGetRedis != nil {
		tx.First(&category, categoryId)
		resRedis, _ := json.Marshal(category)

		errSetRedis := rdb.Set(ctx, "category_name_"+strconv.Itoa(categoryId), resRedis, 0).Err()
		if errSetRedis != nil {
			helper.PanicIfError(errSetRedis)
		}

		return category, nil
	}

	json.Unmarshal([]byte(resultRedis), &category)
	return category, nil

	/*
		   Redis Get & Set By ID
			With TTL
	*/
	//var category domain.Category
	//
	//resultRedis, errGetRedis := rdb.Get(ctx, "category_name_"+strconv.Itoa(categoryId)).Result()
	//
	//if errGetRedis != nil {
	//	tx.First(&category, categoryId)
	//	resRedis, _ := json.Marshal(category)
	//	ttl := time.Duration(10) * time.Second // set time to live is 10 second
	//
	//	errSetRedis := rdb.Set(ctx, "category_name_"+strconv.Itoa(categoryId), resRedis, ttl).Err()
	//	if errSetRedis != nil {
	//		helper.PanicIfError(errSetRedis)
	//	}
	//
	//	return category, nil
	//}
	//
	//json.Unmarshal([]byte(resultRedis), &category)
	//return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB, rdb *redis.Client) []domain.Category {
	/*
		   Redis Get & Set All
			Without TTL
	*/
	var category []domain.Category

	getRedis, errGetRedis := rdb.Get(ctx, "get_all_categories").Result()

	if errGetRedis != nil {
		tx.Find(&category)
		resRedis, _ := json.Marshal(category)

		errSetRedis := rdb.Set(ctx, "get_all_categories", resRedis, 0).Err()
		if errSetRedis != nil {
			helper.PanicIfError(errSetRedis)
		}

		return category
	}

	json.Unmarshal([]byte(getRedis), &category)
	return category

	/*
		   Redis Get & Set All
			With TTL
	*/
	//var category []domain.Category
	//
	//getRedis, errGetRedis := rdb.Get(ctx, "get_all_categories").Result()
	//
	//if errGetRedis != nil {
	//	tx.Find(&category)
	//	resRedis, _ := json.Marshal(category)
	//	ttl := time.Duration(10) * time.Second // set time to live is 10 second
	//
	//	errSetRedis := rdb.Set(ctx, "get_all_categories", resRedis, ttl).Err()
	//	if errSetRedis != nil {
	//		helper.PanicIfError(errSetRedis)
	//	}
	//
	//	return category
	//}
	//
	//json.Unmarshal([]byte(getRedis), &category)
	//return category
}
