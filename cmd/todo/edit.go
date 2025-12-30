package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit [id] [new title]",
	Aliases: []string{"e"},
	Short:   "タスクを編集",
	Args:    cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var id, title string

		switch len(args) {
		case 0:
			todos, err := uc.ListTodos(nil)
			if err != nil {
				return err
			}
			todo, err := selectTodo(todos)
			if err != nil {
				return err
			}
			id = todo.ID

			fmt.Printf("新しいタイトル [%s]: ", todo.Title)
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			title = strings.TrimSpace(input)
			if title == "" {
				title = todo.Title
			}
		case 1:
			return fmt.Errorf("新しいタイトルを指定してください")
		case 2:
			id = args[0]
			title = args[1]
		}

		if err := uc.UpdateTodo(id, title); err != nil {
			return err
		}
		fmt.Println("タスクを更新しました")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
