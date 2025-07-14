package gorm

import (
	"ardiman-xyz/go-todo-app/models/entity"
	"context"
)

type TodoRepositoryGorm interface {
	Save(ctx context.Context, todo *entity.Todo) error
	Update(ctx context.Context, todo *entity.Todo) error
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (*entity.Todo, error)
	FindAll(ctx context.Context) ([]entity.Todo, error)
}