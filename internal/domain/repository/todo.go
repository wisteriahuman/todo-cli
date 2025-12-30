package repository

import "github.com/wisteriahuman/todo-cli/internal/domain/entity"

type FilterStatus string

const (
	FilterAll        FilterStatus = "all"
	FilterCompleted  FilterStatus = "completed"
	FilterIncomplete FilterStatus = "incomplete"
)

type SortField string

const (
	SortByCreated   SortField = "created"
	SortByTitle     SortField = "title"
	SortByCompleted SortField = "completed"
)

type ListOptions struct {
	Filter  FilterStatus
	Search  string
	SortBy  SortField
	Desc    bool
}

func DefaultListOptions() *ListOptions {
	return &ListOptions{
		Filter: FilterAll,
		SortBy: SortByCreated,
		Desc:   false,
	}
}

type TodoRepository interface {
	Create(e *entity.Todo) (*entity.Todo, error)
	FindAll(opts *ListOptions) ([]*entity.Todo, error)
	FindByID(id string) (*entity.Todo, error)
	Update(e *entity.Todo) (*entity.Todo, error)
	Delete(id string) error
}
