package todos

import (
	"fmt"

	"github.com/gorilla/mux"
)

func NewRouter() (*mux.Router, error) {
	r := mux.NewRouter()
	controller, err := NewController()

	if err != nil {
		return nil, fmt.Errorf("failed to create instance of the controller, details: %v", err)
	}

	r.HandleFunc("/todos", controller.Find).Methods("GET")
	r.HandleFunc("/todos/{id}", controller.FindOne).Methods("GET")
	r.HandleFunc("/todos", controller.Create).Methods("POST")
	r.HandleFunc("/todos/{id}", controller.Update).Methods("PATCH")
	r.HandleFunc("/todos/{id}", controller.Remove).Methods("DELETE")

	return r, nil
}
