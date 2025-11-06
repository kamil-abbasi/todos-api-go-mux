package todos

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TodosController struct {
	todosService TodosService
}

func NewController() TodosController {
	service, err := NewService()

	if err != nil {
		panic(fmt.Sprintf("failed to create instance of the service, details: %v", err))
	}

	return TodosController{
		todosService: service,
	}
}

func (controller *TodosController) Find(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	todos, err := controller.todosService.Find()

	if err != nil {
		http.Error(w, "failed to find todos", http.StatusInternalServerError)
		return
	}

	if len(todos) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}

	encoder.Encode(todos)
}

func (controller *TodosController) FindOne(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)

	todo, err := controller.todosService.FindOne(params["id"])

	if err != nil {
		http.Error(w, "failed to find one todo", http.StatusInternalServerError)
		return
	}

	if todo.Id == "" {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	encoder.Encode(todo)
}

func (controller *TodosController) Create(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	var createDto TodoCreateDto

	decoder.Decode(&createDto.Name)

	todo, err := controller.todosService.Create(createDto)

	if err != nil {
		http.Error(w, "failed to create todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	encoder.Encode(todo)
}

func (controller *TodosController) Update(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)
	params := mux.Vars(r)

	var updateDto TodoUpdateDto

	decoder.Decode(&updateDto)

	todo, err := controller.todosService.Update(params["id"], updateDto)

	if err != nil {
		log.Fatal(err)
		http.Error(w, "failed to update todo", http.StatusInternalServerError)
		return
	}

	if todo.Id == "" {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	encoder.Encode(todo)
}

func (controller *TodosController) Remove(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)

	deletedTodo, err := controller.todosService.Remove(params["id"])

	if err != nil {
		http.Error(w, "failed to delete todo", http.StatusInternalServerError)
		return
	}

	if deletedTodo.Id == "" {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	encoder.Encode(deletedTodo)
}
