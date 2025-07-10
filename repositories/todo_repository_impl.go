package repositories

import (
	"ardiman-xyz/go-todo-app/helper"
	"ardiman-xyz/go-todo-app/models/domain"
	"context"
	"database/sql"
	"errors"
)

type TodoRepositoryImpl struct{}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

// Save implements TodoRepository.
func (repository *TodoRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, todo domain.TodoList) domain.TodoList {
		
	SQL :=`INSERT INTO todo (task, status) VALUES (?, ?)`
		result, err := tx.ExecContext(ctx, SQL, todo.Task, todo.Status)
		helper.PanicIfError(err)

		id, err := result.LastInsertId()
	
		helper.PanicIfError(err)

		todo.ID = int(id)
		return todo

}

// FindAll implements TodoRepository.
func (repository *TodoRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.TodoList {
	SQL := ` SELECT id, task, status FROM todo ORDER BY id DESC`

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var todoList []domain.TodoList

	for rows.Next() {
		todo := domain.TodoList{}
		err := rows.Scan(&todo.ID, &todo.Task, &todo.Status)
		helper.PanicIfError(err)
		todoList = append(todoList, todo)
	}

	return todoList
}

// FindById implements TodoRepository.
func (repository *TodoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.TodoList, error) {
	SQL := ` SELECT id, task, status FROM todo WHERE id = ?`

	rows, err := tx.QueryContext(ctx, SQL, todoId)
	helper.PanicIfError(err)

	defer rows.Close()

	todo := domain.TodoList{}
	
	if rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Task, &todo.Status)
		helper.PanicIfError(err)
		return todo, nil
	}else{
		return todo, errors.New("todo is not found")
	}
}

// Update implements TodoRepository.
func (repository *TodoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todo domain.TodoList) domain.TodoList {
	SQL := ` UPDATE todo SET task = ?, status = ? WHERE id = ?`

	_, err := tx.ExecContext(ctx, SQL, todo.Task, todo.Status, todo.ID)
	helper.PanicIfError(err)

	return todo
}

// Delete implements TodoRepository.
func (repository *TodoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todo domain.TodoList) {
	SQL := ` DELETE FROM todo WHERE id = ?`

	_, err := tx.ExecContext(ctx, SQL, todo.ID)
	helper.PanicIfError(err)
	
}