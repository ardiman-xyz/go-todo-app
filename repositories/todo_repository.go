package repositories

import (
	"ardiman-xyz/go-todo-app/models/domain"
	"context"
	"database/sql"
)

type TodoRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todo domain.TodoList) domain.TodoList
	Update(ctx context.Context, tx *sql.Tx, todo domain.TodoList) domain.TodoList
	Delete(ctx context.Context, tx *sql.Tx, todo domain.TodoList)
	FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.TodoList, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.TodoList
}