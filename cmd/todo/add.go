package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "タスクを追加",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := uc.AddTodo(args[0]); err != nil {
			return err
		}
		fmt.Println("タスクを追加しました")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
