package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go_redis_1/config"
	"go_redis_1/controller"
	"go_redis_1/helper"
	"go_redis_1/repository"
	"go_redis_1/service"
	"log"
	"net/http"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDB()
	rdb := config.ConnectRedis()
	validate := validator.New()
	//fmt.Println(db)
	fmt.Println(rdb)

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, rdb, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := config.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
