package internal

import (
	"todos-api/internal/middleware"
	"todos-api/internal/todos"

	"github.com/gorilla/mux"
)

func NewApiRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)

	todosRouter := r.PathPrefix("/todos").Subrouter()
	todosController := todos.NewController()

	todosRouter.HandleFunc("", todosController.Find).Methods("GET")
	todosRouter.HandleFunc("/{id}", todosController.FindOne).Methods("GET")
	todosRouter.HandleFunc("", todosController.Create).Methods("POST")
	todosRouter.HandleFunc("/{id}", todosController.Update).Methods("PATCH")
	todosRouter.HandleFunc("/{id}", todosController.Remove).Methods("DELETE")

	r.Path("/").Handler(todosRouter)

	return r
}
