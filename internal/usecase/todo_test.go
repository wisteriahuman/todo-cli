package usecase

import (
	"testing"

	"github.com/wisteriahuman/todo-cli/internal/domain/entity"
	"github.com/wisteriahuman/todo-cli/internal/domain/repository"
)

type mockTodoRepository struct {
	todos map[string]*entity.Todo
}

func newMockRepository() *mockTodoRepository {
	return &mockTodoRepository{
		todos: make(map[string]*entity.Todo),
	}
}

func (m *mockTodoRepository) Create(e *entity.Todo) (*entity.Todo, error) {
	m.todos[e.ID] = e
	return e, nil
}

func (m *mockTodoRepository) FindAll(opts *repository.ListOptions) ([]*entity.Todo, error) {
	result := make([]*entity.Todo, 0, len(m.todos))
	for _, todo := range m.todos {
		result = append(result, todo)
	}
	return result, nil
}

func (m *mockTodoRepository) FindByID(id string) (*entity.Todo, error) {
	todo, exists := m.todos[id]
	if !exists {
		return nil, nil
	}
	return todo, nil
}

func (m *mockTodoRepository) Update(e *entity.Todo) (*entity.Todo, error) {
	m.todos[e.ID] = e
	return e, nil
}

func (m *mockTodoRepository) Delete(id string) error {
	delete(m.todos, id)
	return nil
}

// TestAddTodo は有効なタイトルでTodoが正常に作成されることをテストする
func TestAddTodo(t *testing.T) {
	repo := newMockRepository()
	uc := NewTodoUsecase(repo)

	err := uc.AddTodo("Test task")

	if err != nil {
		t.Errorf("AddTodo() error = %v, want nil", err)
	}
	if len(repo.todos) != 1 {
		t.Errorf("expected 1 todo, got %d", len(repo.todos))
	}
}

// TestListTodos は複数のTodoが正しく取得できることをテストする
func TestListTodos(t *testing.T) {
	repo := newMockRepository()
	uc := NewTodoUsecase(repo)

	uc.AddTodo("Task 1")
	uc.AddTodo("Task 2")

	todos, err := uc.ListTodos(repository.DefaultListOptions())

	if err != nil {
		t.Errorf("ListTodos() error = %v, want nil", err)
	}
	if len(todos) != 2 {
		t.Errorf("ListTodos() returned %d todos, want 2", len(todos))
	}
}

// TestListTodos_Empty は空のリストが正しく取得できることをテストする
func TestListTodos_Empty(t *testing.T) {
	repo := newMockRepository()
	uc := NewTodoUsecase(repo)

	todos, err := uc.ListTodos(repository.DefaultListOptions())

	if err != nil {
		t.Errorf("ListTodos() error = %v, want nil", err)
	}
	if len(todos) != 0 {
		t.Errorf("ListTodos() returned %d todos, want 0", len(todos))
	}
}

// TestCompleteTodo はTodoが正しく完了状態になることをテストする
func TestCompleteTodo(t *testing.T) {
	repo := newMockRepository()
	uc := NewTodoUsecase(repo)

	uc.AddTodo("Test task")

	var todoID string
	for id := range repo.todos {
		todoID = id
		break
	}

	err := uc.CompleteTodo(todoID)

	if err != nil {
		t.Errorf("CompleteTodo() error = %v, want nil", err)
	}
	if !repo.todos[todoID].Completed {
		t.Error("CompleteTodo() did not set Completed to true")
	}
	if repo.todos[todoID].CompletedAt == nil {
		t.Error("CompleteTodo() did not set CompletedAt")
	}
}

// TestCompleteTodo_NotFound は存在しないIDでエラーが返ることをテストする
func TestCompleteTodo_NotFound(t *testing.T) {
	repo := newMockRepository()
	uc := NewTodoUsecase(repo)

	err := uc.CompleteTodo("non-existent-id")

	if err == nil {
		t.Error("CompleteTodo() expected error for non-existent todo, got nil")
	}
}

// TestUpdateTodo はTodoのタイトルが正しく更新されることをテストする
func TestUpdateTodo(t *testing.T) {
	repo := newMockRepository()
	uc := NewTodoUsecase(repo)

	uc.AddTodo("Original title")

	var todoID string
	for id := range repo.todos {
		todoID = id
		break
	}

	err := uc.UpdateTodo(todoID, "Updated title")

	if err != nil {
		t.Errorf("UpdateTodo() error = %v, want nil", err)
	}
	if repo.todos[todoID].Title != "Updated title" {
		t.Errorf("UpdateTodo() title = %v, want 'Updated title'", repo.todos[todoID].Title)
	}
}

// TestUpdateTodo_NotFound は存在しないIDでエラーが返ることをテストする
func TestUpdateTodo_NotFound(t *testing.T) {
	repo := newMockRepository()
	uc := NewTodoUsecase(repo)

	err := uc.UpdateTodo("non-existent-id", "New title")

	if err == nil {
		t.Error("UpdateTodo() expected error for non-existent todo, got nil")
	}
}

// TestDeleteTodo はTodoが正しく削除されることをテストする
func TestDeleteTodo(t *testing.T) {
	repo := newMockRepository()
	uc := NewTodoUsecase(repo)

	uc.AddTodo("Test task")

	var todoID string
	for id := range repo.todos {
		todoID = id
		break
	}

	err := uc.DeleteTodo(todoID)

	if err != nil {
		t.Errorf("DeleteTodo() error = %v, want nil", err)
	}
	if len(repo.todos) != 0 {
		t.Errorf("DeleteTodo() did not remove todo, %d remain", len(repo.todos))
	}
}
