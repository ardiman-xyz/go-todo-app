package controllers

import (
	"ardiman-xyz/go-todo-app/helper"
	"ardiman-xyz/go-todo-app/models/web"
	"ardiman-xyz/go-todo-app/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TodoControllerImpl struct {
	TodoService services.TodoService
}


func NewTodoController(todoService services.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}

func (c *TodoControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoResponses := c.TodoService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponses,
	}
	
	helper.WriteToResponseBody(writer, webResponse)
}

func (c *TodoControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
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

// Delete implements TodoController.
func (c *TodoControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

// FindById implements TodoController.
func (c *TodoControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

// Update implements TodoController.
func (c *TodoControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
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

// UpdateStatus khusus untuk update status todo
func (c *TodoControllerImpl) UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
