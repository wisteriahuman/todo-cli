package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wisteriahuman/todo-cli/internal/domain/repository"
)

var (
	completeFilter string
	completeSearch string
	completeMulti  bool
)

var completeCmd = &cobra.Command{
	Use:     "complete [id]",
	Aliases: []string{"done", "c"},
	Short:   "タスクを完了",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Filter mode: complete all matching todos
		if completeFilter != "" || completeSearch != "" {
			opts := repository.DefaultListOptions()
			opts.Filter = repository.FilterIncomplete // Only incomplete
			if completeSearch != "" {
				opts.Search = completeSearch
			}

			todos, err := uc.ListTodos(opts)
			if err != nil {
				return err
			}

			if len(todos) == 0 {
				fmt.Println("対象のタスクがありません")
				return nil
			}

			ids := make([]string, len(todos))
			for i, t := range todos {
				ids[i] = t.ID
			}

			if err := uc.CompleteTodos(ids); err != nil {
				return err
			}
			fmt.Printf("%d件のタスクを完了しました\n", len(ids))
			return nil
		}

		if len(args) == 1 {
			if err := uc.CompleteTodo(args[0]); err != nil {
				return err
			}
			fmt.Println("タスクを完了しました")
			return nil
		}

		opts := repository.DefaultListOptions()
		opts.Filter = repository.FilterIncomplete
		todos, err := uc.ListTodos(opts)
		if err != nil {
			return err
		}

		if completeMulti {
			selected, err := selectTodos(todos)
			if err != nil {
				return err
			}
			ids := make([]string, len(selected))
			for i, t := range selected {
				ids[i] = t.ID
			}
			if err := uc.CompleteTodos(ids); err != nil {
				return err
			}
			fmt.Printf("%d件のタスクを完了しました\n", len(ids))
		} else {
			todo, err := selectTodo(todos)
			if err != nil {
				return err
			}
			if err := uc.CompleteTodo(todo.ID); err != nil {
				return err
			}
			fmt.Println("タスクを完了しました")
		}
		return nil
	},
}

func init() {
	completeCmd.Flags().StringVarP(&completeFilter, "filter", "f", "", "Filter: incomplete (apply to all matching)")
	completeCmd.Flags().StringVarP(&completeSearch, "search", "q", "", "Search by title (apply to all matching)")
	completeCmd.Flags().BoolVarP(&completeMulti, "multi", "m", false, "Multi-select mode")
	rootCmd.AddCommand(completeCmd)
}
