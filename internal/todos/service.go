package todos

import "fmt"

type TodosService struct {
	todosRepository TodosRepository
}

func NewService() (TodosService, error) {
	mysqlRepository, err := NewMysqlRepository()

	if err != nil {
		return TodosService{}, fmt.Errorf("failed to create instance of mysql repository, details: %v", err)
	}

	return TodosService{
		todosRepository: mysqlRepository,
	}, nil
}

func (service *TodosService) Find() ([]Todo, error) {
	return service.todosRepository.Find()
}

func (service *TodosService) FindOne(id string) (Todo, error) {
	return service.todosRepository.FindOne(id)
}

func (service *TodosService) Create(createDto TodoCreateDto) (Todo, error) {
	return service.todosRepository.Create(createDto)
}

func (service *TodosService) Update(id string, updateDto TodoUpdateDto) (Todo, error) {
	return service.todosRepository.Update(id, updateDto)
}

func (service *TodosService) Remove(id string) (Todo, error) {
	return service.todosRepository.Remove(id)
}
