package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wisteriahuman/todo-cli/internal/domain/repository"
)

var (
	deleteFilter string
	deleteSearch string
	deleteMulti  bool
)

var deleteCmd = &cobra.Command{
	Use:     "delete [id]",
	Aliases: []string{"rm", "del"},
	Short:   "タスクを削除",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Filter mode: delete all matching todos
		if deleteFilter != "" || deleteSearch != "" {
			opts := repository.DefaultListOptions()
			switch deleteFilter {
			case "completed":
				opts.Filter = repository.FilterCompleted
			case "incomplete":
				opts.Filter = repository.FilterIncomplete
			}
			if deleteSearch != "" {
				opts.Search = deleteSearch
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

			if err := uc.DeleteTodos(ids); err != nil {
				return err
			}
			fmt.Printf("%d件のタスクを削除しました\n", len(ids))
			return nil
		}

		if len(args) == 1 {
			if err := uc.DeleteTodo(args[0]); err != nil {
				return err
			}
			fmt.Println("タスクを削除しました")
			return nil
		}

		todos, err := uc.ListTodos(nil)
		if err != nil {
			return err
		}

		if deleteMulti {
			selected, err := selectTodos(todos)
			if err != nil {
				return err
			}
			ids := make([]string, len(selected))
			for i, t := range selected {
				ids[i] = t.ID
			}
			if err := uc.DeleteTodos(ids); err != nil {
				return err
			}
			fmt.Printf("%d件のタスクを削除しました\n", len(ids))
		} else {
			todo, err := selectTodo(todos)
			if err != nil {
				return err
			}
			if err := uc.DeleteTodo(todo.ID); err != nil {
				return err
			}
			fmt.Println("タスクを削除しました")
		}
		return nil
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&deleteFilter, "filter", "f", "", "Filter: completed, incomplete (apply to all matching)")
	deleteCmd.Flags().StringVarP(&deleteSearch, "search", "q", "", "Search by title (apply to all matching)")
	deleteCmd.Flags().BoolVarP(&deleteMulti, "multi", "m", false, "Multi-select mode")
	rootCmd.AddCommand(deleteCmd)
}
