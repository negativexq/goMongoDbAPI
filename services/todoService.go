package services

import (
	"goMongoDbAPI/dto"
	"goMongoDbAPI/models"
	"goMongoDbAPI/repository"
)

type DefaultTodoService struct {
	Repo repository.TodoRepository
}

type TodoService interface {
	TodoInsert(todo models.Todo) (*dto.TodoDTO, error)
}

func (t DefaultTodoService) TodoInsert(todo models.Todo) (*dto.TodoDTO, error) {
	var res dto.TodoDTO
	if len(todo.Title) <= 2 {
		res.Status = false
		return &res, nil
	}
	result, err := t.Repo.Insert(todo)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}
	res = dto.TodoDTO{Status: result}
	return &res, nil
}

func NewTodoService(Repo repository.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}
