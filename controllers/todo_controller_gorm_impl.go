package controllers

import (
	"ardiman-xyz/go-todo-app/helper"
	"ardiman-xyz/go-todo-app/models/web"
	"ardiman-xyz/go-todo-app/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TodoControllerGormImpl struct {
	TodoService services.TodoServiceGorm
}

func NewTodoControllerGorm(todoService services.TodoServiceGorm) TodoControllerGorm {
	return &TodoControllerGormImpl{
		TodoService: todoService,
	}
}

func (c *TodoControllerGormImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoResponses := c.TodoService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponses,
	}
	
	helper.WriteToResponseBody(writer, webResponse)
}

func (c *TodoControllerGormImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoCreateRequest := web.TodoCreateRequest{}
	helper.ReadFromRequestBody(request, &todoCreateRequest)
	
	todoResponse := c.TodoService.Create(request.Context(), todoCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}
	
	helper.WriteToResponseBody(writer, webResponse)
}

func (c *TodoControllerGormImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	c.TodoService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *TodoControllerGormImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	todoResponse := c.TodoService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data: todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *TodoControllerGormImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoUpdateRequest := web.TodoUpdateRequest{}
	helper.ReadFromRequestBody(request, &todoUpdateRequest)

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)
	todoUpdateRequest.Id = id

	todoResponse := c.TodoService.Update(request.Context(), todoUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *TodoControllerGormImpl) UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	type StatusUpdateRequest struct {
		Status bool `json:"status"`
	}

	statusUpdate := StatusUpdateRequest{}
	helper.ReadFromRequestBody(request, &statusUpdate)

	// Ambil todo yang ada
	todoResponse := c.TodoService.FindById(request.Context(), id)

	// Update hanya status
	todoUpdateRequest := web.TodoUpdateRequest{
		Id:     id,
		Task:   todoResponse.Task,
		Status: statusUpdate.Status,
	}

	updatedTodo := c.TodoService.Update(request.Context(), todoUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   updatedTodo,
	}

	helper.WriteToResponseBody(writer, webResponse)
}