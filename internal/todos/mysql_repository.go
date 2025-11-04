package todos

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type TodosMysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository() (TodosRepository, error) {
	db, err := sql.Open("mysql", "root:password@/todos?parseTime=true")

	if err != nil {
		return nil, fmt.Errorf("failed to establish mysql connection, details: %v", err)
	}

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("failed to ping mysql db, details: %v", err)
	}

	query := `
		CREATE TABLE IF NOT EXISTS todos (
			id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
			name TEXT NOT NULL
		)
	`

	_, err = db.Exec(query)

	if err != nil {
		return nil, fmt.Errorf("failed to create todos table, details: %v", err)
	}

	return &TodosMysqlRepository{
		db: db,
	}, nil
}

func (r *TodosMysqlRepository) Find() ([]Todo, error) {
	rows, err := r.db.Query("SELECT id, name FROM todos")

	if err != nil {
		return []Todo{}, fmt.Errorf("failed to execute the query, details: %v", err)
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Name)

		if err != nil {
			return []Todo{}, fmt.Errorf("failed to convert db values, details: %v", err)
		}

		todos = append(todos, todo)
	}

	err = rows.Err()

	if err != nil {
		return []Todo{}, fmt.Errorf("failed to fetch all todos, details: %v", err)
	}

	return todos, nil
}

func (r *TodosMysqlRepository) FindOne(id string) (Todo, error) {
	query := `SELECT id, name FROM todos WHERE id = ?`

	var todo Todo

	err := r.db.QueryRow(query, id).Scan(&todo.Id, &todo.Name)

	if err == sql.ErrNoRows {
		return Todo{}, nil
	} else if err != nil {
		return Todo{}, fmt.Errorf("failed to fetch todo, details: %v", err)
	}

	return todo, nil
}

func (r *TodosMysqlRepository) Create(createDto TodoCreateDto) (Todo, error) {
	id := uuid.NewString()

	_, err := r.db.Exec(`INSERT INTO todos (id, name) VALUES (?, ?)`, id, createDto.Name)

	if err != nil {
		return Todo{}, fmt.Errorf("failed to insert a todo, details: %v", err)
	}

	var todo Todo = Todo{
		Id: id,
	}

	if err := r.db.QueryRow(`SELECT name FROM todos WHERE id = ?`, id).Scan(&todo.Name); err != nil {
		return Todo{}, fmt.Errorf("failed to retrieve inserted todo, details: %v", err)
	}

	return todo, nil
}

func (r *TodosMysqlRepository) Update(id string, updateDto TodoUpdateDto) (Todo, error) {
	result, err := r.db.Exec(`UPDATE todos SET name = ? WHERE id = ?`, updateDto.Name, id)

	if err != nil {
		return Todo{}, fmt.Errorf("failed to update todo, details: %v", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return Todo{}, fmt.Errorf("failed to fetch count of affected rows, details: %v", err)
	}

	if rowsAffected <= 0 {
		return Todo{}, nil
	}

	var todo Todo

	if err := r.db.QueryRow("SELECT id, name FROM todos WHERE id = ?", id).Scan(&todo.Id, &todo.Name); err != nil {
		return Todo{}, fmt.Errorf("failed to fetch updated todo, details: %v", err)
	}

	return todo, nil
}

func (r *TodosMysqlRepository) Remove(id string) (Todo, error) {
	var todo Todo

	err := r.db.QueryRow("SELECT id, name FROM todos WHERE id = ?", id).Scan(&todo.Id, &todo.Name)

	if err == sql.ErrNoRows {
		return Todo{}, nil
	}

	if err != nil {
		return Todo{}, fmt.Errorf("failed to fetch todo, details: %v", err)
	}

	result, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)

	if err != nil {
		return Todo{}, fmt.Errorf("failed to delete todo, details: %v", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return Todo{}, fmt.Errorf("failed to fetch affected rows count")
	}

	if rowsAffected <= 0 {
		return Todo{}, nil
	}

	return todo, nil
}
