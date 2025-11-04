package todos

type TodosRepository interface {
	Find() ([]Todo, error)
	FindOne(id string) (Todo, error)
	Create(createDto TodoCreateDto) (Todo, error)
	Update(id string, updateDto TodoUpdateDto) (Todo, error)
	Remove(id string) (Todo, error)
}
