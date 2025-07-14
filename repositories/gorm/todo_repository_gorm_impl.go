package gorm

import (
	"ardiman-xyz/go-todo-app/models/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type TodoRepositoryGormImpl struct {
	DB *gorm.DB
}

func NewTodoRepositoryGorm(db *gorm.DB) TodoRepositoryGorm {
	return &TodoRepositoryGormImpl{DB: db}
}

func (r *TodoRepositoryGormImpl) Save(ctx context.Context, todo *entity.Todo) error {
	return r.DB.WithContext(ctx).Create(todo).Error
}

func (r *TodoRepositoryGormImpl) Update(ctx context.Context, todo *entity.Todo) error {
	return r.DB.WithContext(ctx).Save(todo).Error
}

func (r *TodoRepositoryGormImpl) Delete(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Delete(&entity.Todo{}, id).Error
}

func (r *TodoRepositoryGormImpl) FindById(ctx context.Context, id int) (*entity.Todo, error) {
	var todo entity.Todo
	err := r.DB.WithContext(ctx).First(&todo, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepositoryGormImpl) FindAll(ctx context.Context) ([]entity.Todo, error) {
	var todos []entity.Todo
	err := r.DB.WithContext(ctx).Order("id DESC").Find(&todos).Error
	return todos, err
}