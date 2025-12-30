package usecase

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wisteriahuman/todo-cli/internal/domain/entity"
	"github.com/wisteriahuman/todo-cli/internal/domain/repository"
)

type TodoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(repo repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{repo: repo}
}

func (uc *TodoUsecase) AddTodo(title string) error {
	todo := &entity.Todo{
		ID:        uuid.New().String(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	_, err := uc.repo.Create(todo)
	return err
}

func (uc *TodoUsecase) ListTodos(opts *repository.ListOptions) ([]*entity.Todo, error) {
	return uc.repo.FindAll(opts)
}

func (uc *TodoUsecase) CompleteTodo(id string) error {
	todo, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}
	if todo == nil {
		return fmt.Errorf("todo not found: %s", id)
	}
	now := time.Now()
	todo.Completed = true
	todo.CompletedAt = &now
	_, err = uc.repo.Update(todo)
	return err
}

func (uc *TodoUsecase) UpdateTodo(id, title string) error {
	todo, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}
	if todo == nil {
		return fmt.Errorf("todo not found: %s", id)
	}
	todo.Title = title
	_, err = uc.repo.Update(todo)
	return err
}

func (uc *TodoUsecase) DeleteTodo(id string) error {
	return uc.repo.Delete(id)
}

func (uc *TodoUsecase) CompleteTodos(ids []string) error {
	for _, id := range ids {
		if err := uc.CompleteTodo(id); err != nil {
			return err
		}
	}
	return nil
}

func (uc *TodoUsecase) DeleteTodos(ids []string) error {
	for _, id := range ids {
		if err := uc.DeleteTodo(id); err != nil {
			return err
		}
	}
	return nil
}
