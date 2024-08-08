/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commit

import (
	"committee/internal"
	"fmt"

	"github.com/spf13/cobra"
)

// choreCmd represents the chore command
var choreCmd = &cobra.Command{
	Use:   "chore",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		com := internal.NewCommitMessage("chore", message)
		internal.Commit(com)
	},
}

func init() {
	choreCmd.Flags().StringVarP(&message, "message", "m", "", "")

	if err := choreCmd.MarkFlagRequired("message"); err != nil {
		fmt.Println(err.Error())
	}

	CommitCmd.AddCommand(choreCmd)
}
