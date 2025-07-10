package repositories

import (
	"ardiman-xyz/go-todo-app/models/domain"
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)


func TestTodoRepositoryImpl_Save_Simple(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Test data
	todo := domain.TodoList{
		Task:   "Test Task",
		Status: false,
	}

	mock.ExpectBegin()
	tx, _ := db.Begin()

	// Mock expectations
	mock.ExpectExec(`INSERT INTO todo \(task, status\) VALUES \(\?, \?\)`).
		WithArgs(todo.Task, todo.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	repo := NewTodoRepository()
	result := repo.Save(context.Background(), tx, todo)

	// Commit transaction
	tx.Commit()

	// Assertions
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Task", result.Task)
	assert.Equal(t, false, result.Status)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTodoRepositoryImpl_FindAll(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock transaction begin
	mock.ExpectBegin()

	// Begin transaction
	tx, err := db.Begin()
	assert.NoError(t, err)

	// Mock data
	rows := sqlmock.NewRows([]string{"id", "task", "status"}).
		AddRow(1, "Task 1", false).
		AddRow(2, "Task 2", true)

	// Mock expectations
	mock.ExpectQuery(`SELECT id, task, status FROM todo ORDER BY id DESC`).
		WillReturnRows(rows)

	// Execute
	repo := NewTodoRepository()
	result := repo.FindAll(context.Background(), tx)

	// Assertions
	assert.Len(t, result, 2)
	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "Task 1", result[0].Task)
	assert.Equal(t, false, result[0].Status)
	assert.Equal(t, 2, result[1].ID)
	assert.Equal(t, "Task 2", result[1].Task)
	assert.Equal(t, true, result[1].Status)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTodoRepositoryImpl_FindById_Success(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock transaction begin
	mock.ExpectBegin()

	// Begin transaction
	tx, err := db.Begin()
	assert.NoError(t, err)

	// Mock data
	rows := sqlmock.NewRows([]string{"id", "task", "status"}).
		AddRow(1, "Test Task", false)

	mock.ExpectQuery(`SELECT id, task, status FROM todo WHERE id = \?`).
		WithArgs(1).
		WillReturnRows(rows)

	// Execute
	repo := NewTodoRepository()
	result, err := repo.FindById(context.Background(), tx, 1)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Task", result.Task)
	assert.Equal(t, false, result.Status)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTodoRepositoryImpl_Update(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock transaction begin
	mock.ExpectBegin()

	// Begin transaction
	tx, err := db.Begin()
	assert.NoError(t, err)

	// Test data
	todo := domain.TodoList{
		ID:     1,
		Task:   "Updated Task",
		Status: true,
	}

	// Mock expectations
	mock.ExpectExec(`UPDATE todo SET task = \?, status = \? WHERE id = \?`).
		WithArgs(todo.Task, todo.Status, todo.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Execute
	repo := NewTodoRepository()
	result := repo.Update(context.Background(), tx, todo)

	// Assertions
	assert.Equal(t, todo.ID, result.ID)
	assert.Equal(t, todo.Task, result.Task)
	assert.Equal(t, todo.Status, result.Status)
	assert.NoError(t, mock.ExpectationsWereMet())
}


func TestTodoRepositoryImpl_Delete(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock transaction begin
	mock.ExpectBegin()

	// Begin transaction
	tx, err := db.Begin()
	assert.NoError(t, err)

	// Test data
	todo := domain.TodoList{
		ID: 1,
	}

	// Mock expectations
	mock.ExpectExec(`DELETE FROM todo WHERE id = \?`).
		WithArgs(todo.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Execute
	repo := NewTodoRepository()
	repo.Delete(context.Background(), tx, todo)

	// Assertions
	assert.NoError(t, mock.ExpectationsWereMet())
}
