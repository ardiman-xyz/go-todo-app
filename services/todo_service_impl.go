package services

import (
	"ardiman-xyz/go-todo-app/exception"
	"ardiman-xyz/go-todo-app/helper"
	"ardiman-xyz/go-todo-app/models/domain"
	"ardiman-xyz/go-todo-app/models/web"
	"ardiman-xyz/go-todo-app/repositories"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type TodoServiceImpl struct {
	TodoRepository repositories.TodoRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewTodoService(todoRepository repositories.TodoRepository, db *sql.DB, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		DB:             db,
		Validate:       validate,
	}
}

// Create implements TodoService.
func (service *TodoServiceImpl) Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo := domain.TodoList{
		Task: 	  request.Task,
		Status:  false,
	}

	todo = service.TodoRepository.Save(ctx, tx, todo)


	return helper.ToTodoResponse(todo)

}

// Delete implements TodoService.
func (service *TodoServiceImpl) Delete(ctx context.Context, todoId int) {
		tx, err := service.DB.Begin()
		helper.PanicIfError(err)
		defer helper.CommitOrRollback(tx)

		todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		service.TodoRepository.Delete(ctx, tx, todo)
}

// FindAll implements TodoService.
func (service *TodoServiceImpl) FindAll(ctx context.Context) []web.TodoResponse {
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todos := service.TodoRepository.FindAll(ctx, tx)

	return helper.ToTodoResponseList(todos)
}

// FindById implements TodoService.
func (service *TodoServiceImpl) FindById(ctx context.Context, todoId int) web.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo,err := service.TodoRepository.FindById(ctx, tx, todoId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToTodoResponse(todo)
}

// Update implements TodoService.
func (service *TodoServiceImpl) Update(ctx context.Context, request web.TodoUpdateRequest) web.TodoResponse {
	
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo := domain.TodoList{
		ID:     request.Id,
		Task:   request.Task,
		Status: request.Status,	
	}

	service.TodoRepository.Update(ctx, tx, todo)

	return helper.ToTodoResponse(todo)

}

