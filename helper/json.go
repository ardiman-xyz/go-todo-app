package helper

import (
	"ardiman-xyz/go-todo-app/models/domain"
	"ardiman-xyz/go-todo-app/models/web"
	"encoding/json"
	"net/http"
)


func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(result)
	PanicIfError(err)
}


func ToTodoResponse(todo domain.TodoList) web.TodoResponse {
	return web.TodoResponse{
		Id:   todo.ID,
		Task: todo.Task,
		Status: todo.Status,
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}


func ToTodoResponseList(todos []domain.TodoList) []web.TodoResponse {
	var todoResponses []web.TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, ToTodoResponse(todo))
	}

	return todoResponses
}
