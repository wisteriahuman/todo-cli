package main

import (
	"log"
	"os"

	"github.com/wisteriahuman/todo-cli/internal/infra"
	"github.com/wisteriahuman/todo-cli/internal/usecase"
)

var uc *usecase.TodoUsecase

func main() {
	dbPath, err := infra.DefaultDBPath()
	if err != nil {
		log.Fatal(err)
	}

	db, err := infra.New(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	uc = usecase.NewTodoUsecase(db)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
