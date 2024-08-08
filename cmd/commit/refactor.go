/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package commit

import (
	"committee/internal"
	"fmt"

	"github.com/spf13/cobra"
)

// refactorCmd represents the refactor command
var refactorCmd = &cobra.Command{
	Use:   "refactor",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		com := internal.NewCommitMessage("refactor", message)
		internal.Commit(com)
	},
}

func init() {
	refactorCmd.Flags().StringVarP(&message, "message", "m", "", "")

	if err := refactorCmd.MarkFlagRequired("message"); err != nil {
		fmt.Println(err.Error())
	}

	CommitCmd.AddCommand(refactorCmd)
}
