package infra

import (
	"database/sql"
	"os"
	"path/filepath"
)

type DB struct {
	*sql.DB
}

func DefaultDBPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(configDir, "todo-cli")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(appDir, "todo.db"), nil
}

func New(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS todos (
				id TEXT PRIMARY KEY,
				title TEXT NOT NULL,
				completed INTEGER NOT NULL DEFAULT 0,
				created_at DATETIME NOT NULL,
				completed_at DATETIME
			)
	`)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.DB.Close()
}
