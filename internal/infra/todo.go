package infra

import (
	"database/sql"
	"fmt"
	"strings"

	_ "modernc.org/sqlite"

	"github.com/wisteriahuman/todo-cli/internal/domain/entity"
	"github.com/wisteriahuman/todo-cli/internal/domain/repository"
)

func (db *DB) Create(e *entity.Todo) (*entity.Todo, error) {
	_, err := db.Exec(`
		INSERT INTO todos
			(id, title, completed, created_at, completed_at)
		VALUES
			(?, ?, ?, ?, ?)
	`, e.ID, e.Title, e.Completed, e.CreatedAt, e.CompletedAt,
	)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (db *DB) FindAll(opts *repository.ListOptions) ([]*entity.Todo, error) {
	if opts == nil {
		opts = repository.DefaultListOptions()
	}

	query := "SELECT id, title, completed, created_at, completed_at FROM todos"
	var conditions []string
	var args []interface{}

	switch opts.Filter {
	case repository.FilterCompleted:
		conditions = append(conditions, "completed = 1")
	case repository.FilterIncomplete:
		conditions = append(conditions, "completed = 0")
	}

	if opts.Search != "" {
		conditions = append(conditions, "title LIKE ?")
		args = append(args, "%"+opts.Search+"%")
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	var orderBy string
	switch opts.SortBy {
	case repository.SortByTitle:
		orderBy = "title"
	case repository.SortByCompleted:
		orderBy = "completed"
	default:
		orderBy = "created_at"
	}

	if opts.Desc {
		query += fmt.Sprintf(" ORDER BY %s DESC", orderBy)
	} else {
		query += fmt.Sprintf(" ORDER BY %s ASC", orderBy)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*entity.Todo{}
	for rows.Next() {
		todo := &entity.Todo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.CompletedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (db *DB) FindByID(id string) (*entity.Todo, error) {
	row := db.QueryRow(`
		SELECT
			id, title, completed, created_at, completed_at
		FROM todos
		WHERE id = ?
	`, id)

	todo := &entity.Todo{}
	err := row.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.CompletedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (db *DB) Update(e *entity.Todo) (*entity.Todo, error) {
	_, err := db.Exec(`
		UPDATE todos
		SET
			title = ?, completed = ?, completed_at = ?
		WHERE id = ?
	`, e.Title, e.Completed, e.CompletedAt, e.ID,
	)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (db *DB) Delete(id string) error {
	_, err := db.Exec(`
		DELETE
		FROM todos
		WHERE id = ?
		`, id)
	return err
}
