package config

import (
	"github.com/julienschmidt/httprouter"
	"go_redis_1/controller"
	"go_redis_1/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)

	router.PanicHandler = exception.ErrorHandler

	return router
}
