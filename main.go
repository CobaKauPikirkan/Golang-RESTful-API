package main

import (
	_ "github.com/go-sql-driver/mysql"

	"fachrurozirizky/belajar-golang-restful-api/app"
	"fachrurozirizky/belajar-golang-restful-api/controller"
	//"fachrurozirizky/belajar-golang-restful-api/exception"
	"fachrurozirizky/belajar-golang-restful-api/helper"
	"fachrurozirizky/belajar-golang-restful-api/middleware"
	"fachrurozirizky/belajar-golang-restful-api/repository"
	"fachrurozirizky/belajar-golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	//"github.com/julienschmidt/httprouter"
)

func main() {
	db:= app.NewDB()
	validate:= validator.New()

	categoryRepository:= repository.NewCategoryRepository()
	categoryService:= service.NewCategoryService(categoryRepository, db, validate)
	categoryController:= controller.NewCategoryController(categoryService)

	router:= app.NewRouter(categoryController)//semua router dijadikan satu dalam satu function

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),//digunakan untuk authorization API check folder middleware
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}