package controller

import (
	"fachrurozirizky/belajar-golang-restful-api/helper"
	"fachrurozirizky/belajar-golang-restful-api/model/web"
	"fachrurozirizky/belajar-golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//diguanakan untuk read request body terlebih dahulu
	categoryCreateRequest:= web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse:=controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}
	//digunakan untuk write response body , function ada di file json.go
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryUpdateRequest:= web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse:=controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}
	//digunakan untuk write response body , function ada di file json.go
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) { 
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}
	//digunakan untuk write response body , function ada di file json.go
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse:=controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}
	//digunakan untuk write response body , function ada di file json.go
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryResponses:= controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponses,
	}
	//digunakan untuk write response body , function ada di file json.go
	helper.WriteToResponseBody(writer, webResponse)
}