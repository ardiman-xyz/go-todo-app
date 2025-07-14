package services

import (
	"ardiman-xyz/go-todo-app/helper"
	"ardiman-xyz/go-todo-app/models/entity"
	"ardiman-xyz/go-todo-app/models/web"
	gormRepo "ardiman-xyz/go-todo-app/repositories/gorm"
	"context"

	"github.com/go-playground/validator"
)

type TodoServiceGormImpl struct {
	TodoRepository gormRepo.TodoRepositoryGorm
	Validate       *validator.Validate
}

func NewTodoServiceGorm(todoRepository gormRepo.TodoRepositoryGorm, validate *validator.Validate) TodoServiceGorm {
	return &TodoServiceGormImpl{
		TodoRepository: todoRepository,
		Validate:       validate,
	}
}

func (s *TodoServiceGormImpl) Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	todo := &entity.Todo{
		Task:   request.Task,
		Status: false,
	}

	err = s.TodoRepository.Save(ctx, todo)
	helper.PanicIfError(err)

	return web.TodoResponse{
		Id:     int(todo.ID),
		Task:   todo.Task,
		Status: todo.Status,
	}
}

func (s *TodoServiceGormImpl) Update(ctx context.Context, request web.TodoUpdateRequest) web.TodoResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	// Find existing todo first
	todo, err := s.TodoRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	// Update fields
	todo.Task = request.Task
	todo.Status = request.Status

	err = s.TodoRepository.Update(ctx, todo)
	helper.PanicIfError(err)

	return web.TodoResponse{
		Id:     int(todo.ID),
		Task:   todo.Task,
		Status: todo.Status,
	}
}

func (s *TodoServiceGormImpl) Delete(ctx context.Context, todoId int) {
	// Check if exists
	_, err := s.TodoRepository.FindById(ctx, todoId)
	helper.PanicIfError(err)

	err = s.TodoRepository.Delete(ctx, todoId)
	helper.PanicIfError(err)
}

func (s *TodoServiceGormImpl) FindById(ctx context.Context, todoId int) web.TodoResponse {
	todo, err := s.TodoRepository.FindById(ctx, todoId)
	helper.PanicIfError(err)

	return web.TodoResponse{
		Id:     int(todo.ID),
		Task:   todo.Task,
		Status: todo.Status,
	}
}

func (s *TodoServiceGormImpl) FindAll(ctx context.Context) []web.TodoResponse {
	todos, err := s.TodoRepository.FindAll(ctx)
	helper.PanicIfError(err)

	var todoResponses []web.TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, web.TodoResponse{
			Id:     int(todo.ID),
			Task:   todo.Task,
			Status: todo.Status,
		})
	}

	return todoResponses
}