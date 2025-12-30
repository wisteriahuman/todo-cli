package main

import (
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/spf13/cobra"
	"github.com/wisteriahuman/todo-cli/internal/domain/repository"
)

var (
	listFilter string
	listSearch string
	listSort   string
	listDesc   bool
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "タスク一覧を表示",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := repository.DefaultListOptions()

		switch listFilter {
		case "completed":
			opts.Filter = repository.FilterCompleted
		case "incomplete":
			opts.Filter = repository.FilterIncomplete
		}

		opts.Search = listSearch

		switch listSort {
		case "title":
			opts.SortBy = repository.SortByTitle
		case "completed":
			opts.SortBy = repository.SortByCompleted
		}

		opts.Desc = listDesc

		todos, err := uc.ListTodos(opts)
		if err != nil {
			return err
		}
		t := table.New(os.Stdout)
		t.SetRowLines(false)
		t.SetHeaders("#", "ID", "Title", "Completed", "Created At", "Completed At")
		t.SetHeaderStyle(table.StyleBold)
		t.SetLineStyle(table.StyleBlue)
		t.SetDividers(table.UnicodeDividers)
		if len(todos) == 0 {
			t.Render()
			return nil
		}
		for index, todo := range todos {
			completed := "❌️"
			completedAt := ""
			if todo.Completed {
				completed = "✅️"
				completedAt = todo.CompletedAt.Format(time.ANSIC)
			}
			t.AddRow(strconv.Itoa(index), todo.ID[:8], todo.Title, completed, todo.CreatedAt.Format(time.ANSIC), completedAt)
		}
		t.Render()
		return nil
	},
}

func init() {
	listCmd.Flags().StringVarP(&listFilter, "filter", "f", "all", "Filter: all, completed, incomplete")
	listCmd.Flags().StringVarP(&listSearch, "search", "q", "", "Search by title")
	listCmd.Flags().StringVarP(&listSort, "sort", "s", "created", "Sort: created, title, completed")
	listCmd.Flags().BoolVarP(&listDesc, "desc", "d", false, "Sort descending")
	rootCmd.AddCommand(listCmd)
}
