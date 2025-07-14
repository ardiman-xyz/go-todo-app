package main

import (
	"ardiman-xyz/go-todo-app/app"
	"ardiman-xyz/go-todo-app/controllers"
	"ardiman-xyz/go-todo-app/helper"
	"ardiman-xyz/go-todo-app/repositories"
	"ardiman-xyz/go-todo-app/services"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func enableCORSw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func mainenc() {

	db := app.NewDB()
	validate := validator.New()

	todoRepository := repositories.NewTodoRepository()
	todoService := services.NewTodoService(todoRepository, db, validate)
	todoController := controllers.NewTodoController(todoService)

	router := httprouter.New()


	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Write([]byte("Hello world"))
	})

	router.GET("/api/todos", todoController.FindAll)
	router.POST("/api/todos", todoController.Create)
	router.PUT("/api/todos/:todoId", todoController.Update)
	router.PATCH("/api/todos/:todoId/status", todoController.UpdateStatus)
	router.DELETE("/api/todos/:todoId", todoController.Delete)
	router.GET("/api/todos/:todoId", todoController.FindById)


	 server := http.Server{
		Addr:    "localhost:3001",
		Handler: enableCORS(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}


