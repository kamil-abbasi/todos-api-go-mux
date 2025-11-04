package main

import (
	"fmt"
	"net/http"
	"todos-api/internal/todos"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	todosRouter, err := todos.NewRouter()

	if err != nil {
		fmt.Printf("failed to create instance of todos router, details: %v", err)
		panic("Error!")
	}

	r.PathPrefix("/").Handler(todosRouter)

	http.ListenAndServe(":3000", r)
}
