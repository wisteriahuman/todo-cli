package main

import (
	"log"
	"os"

	"github.com/wisteriahuman/todo-cli/internal/infra"
	"github.com/wisteriahuman/todo-cli/internal/usecase"
)

var uc *usecase.TodoUsecase

func main() {
	db, err := infra.New("todo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	uc = usecase.NewTodoUsecase(db)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
