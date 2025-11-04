package todos

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type TodosCSVRepository struct{}

func NewCSVRepository() TodosRepository {
	return &TodosCSVRepository{}
}

func (r *TodosCSVRepository) Find() ([]Todo, error) {
	file, err := os.OpenFile("todos.csv", os.O_CREATE|os.O_RDONLY, 0644)

	if err != nil {
		return []Todo{}, fmt.Errorf(
			"failed to open todos.csv, details: %v",
			err,
		)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return []Todo{}, fmt.Errorf(
			"failed to read records from todos.csv, details: %v",
			err,
		)
	}

	var todos []Todo = []Todo{}

	for _, record := range records {
		todoId := record[0]
		todoName := record[1]

		var todo Todo = Todo{
			Id:   todoId,
			Name: todoName,
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *TodosCSVRepository) FindOne(id string) (Todo, error) {
	return Todo{}, nil
}

func (r *TodosCSVRepository) Create(createDto TodoCreateDto) (Todo, error) {
	file, err := os.OpenFile("todos.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return Todo{}, fmt.Errorf(
			"failed to open todos.csv, details: %v",
			err,
		)
	}

	defer file.Close()

	todo := Todo{
		Id:   uuid.New().String(),
		Name: createDto.Name,
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{todo.Id, todo.Name})
	writer.Flush()

	return todo, nil
}

func (r *TodosCSVRepository) Update(id string, updateDto TodoUpdateDto) (Todo, error) {
	return Todo{}, nil
}

func (r *TodosCSVRepository) Remove(id string) (Todo, error) {
	file, err := os.OpenFile("todos.csv", os.O_RDONLY, 0644)

	if err != nil {
		return Todo{}, fmt.Errorf("failed to open todos.csv, details: %v", err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return Todo{}, fmt.Errorf(
			"failed to read records from todos.csv, details: %v",
			err,
		)
	}

	var recordsToWrite [][]string = [][]string{}
	var todoToDelete Todo

	for _, record := range records {
		if record[0] != id {
			recordsToWrite = append(recordsToWrite, record)
		} else {
			todoToDelete = Todo{
				Id:   record[0],
				Name: record[1],
			}
		}
	}

	file.Close()

	fileToWrite, err := os.Create("todos.csv")

	if err != nil {
		return Todo{}, fmt.Errorf(
			"failed to open todos.csv for writing, details: %v",
			err,
		)
	}

	writer := csv.NewWriter(fileToWrite)
	writer.WriteAll(recordsToWrite)
	writer.Flush()

	return todoToDelete, nil
}
