/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commit

import (
	"committee/internal"
	"fmt"

	"github.com/spf13/cobra"
)

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		com := internal.NewCommitMessage("fix", message)
		internal.Commit(com)
	},
}

func init() {
	fixCmd.Flags().StringVarP(&message, "message", "m", "", "")

	if err := fixCmd.MarkFlagRequired("message"); err != nil {
		fmt.Println(err.Error())
	}
	CommitCmd.AddCommand(fixCmd)
}
