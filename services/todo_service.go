package services

import (
	"ardiman-xyz/go-todo-app/models/web"
	"context"
)

type TodoService interface {
	Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse
	Update(ctx context.Context, request web.TodoUpdateRequest) web.TodoResponse
	Delete(ctx context.Context, todoId int)
	FindById(ctx context.Context, todoId int) web.TodoResponse
	FindAll(ctx context.Context) []web.TodoResponse
}